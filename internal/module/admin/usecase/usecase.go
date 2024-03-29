package usecase

import (
	"context"
	"database/sql"
	"math"
	"murakali/config"
	"murakali/internal/constant"
	"murakali/internal/model"
	"murakali/internal/module/admin"
	"murakali/internal/module/admin/delivery/body"
	"murakali/pkg/httperror"
	"murakali/pkg/pagination"
	"murakali/pkg/postgre"
	"murakali/pkg/response"
	"net/http"
	"time"
)

type adminUC struct {
	cfg       *config.Config
	txRepo    *postgre.TxRepo
	adminRepo admin.Repository
}

func NewAdminUseCase(cfg *config.Config, txRepo *postgre.TxRepo, adminRepo admin.Repository) admin.UseCase {
	return &adminUC{cfg: cfg, txRepo: txRepo, adminRepo: adminRepo}
}

func (u *adminUC) GetAllVoucher(ctx context.Context, voucherStatusID, sortFilter string, pgn *pagination.Pagination) (*pagination.Pagination, error) {
	totalRows, err := u.adminRepo.GetTotalVoucher(ctx, voucherStatusID)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(pgn.Limit)))
	pgn.TotalRows = totalRows
	pgn.TotalPages = totalPages

	ShopVouchers, err := u.adminRepo.GetAllVoucher(ctx, voucherStatusID, sortFilter, pgn)
	if err != nil {
		return nil, err
	}

	pgn.Rows = ShopVouchers

	return pgn, nil
}

func (u *adminUC) GetRefunds(ctx context.Context, sortFilter string, pgn *pagination.Pagination) (*pagination.Pagination, error) {
	totalRows, err := u.adminRepo.GetTotalRefunds(ctx)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(pgn.Limit)))
	pgn.TotalRows = totalRows
	pgn.TotalPages = totalPages

	refunds, err := u.adminRepo.GetRefunds(ctx, sortFilter, pgn)
	if err != nil {
		return nil, err
	}

	pgn.Rows = refunds
	return pgn, nil
}

func (u *adminUC) RefundOrder(ctx context.Context, refundID string) error {
	refund, err := u.adminRepo.GetRefundByID(ctx, refundID)
	if err != nil {
		if err == sql.ErrNoRows {
			return httperror.New(http.StatusBadRequest, response.RefundNotFound)
		}

		return err
	}

	if refund.RejectedAt.Valid {
		return httperror.New(http.StatusBadRequest, response.RefundRejected)
	}

	if refund.RefundedAt.Valid {
		return httperror.New(http.StatusBadRequest, response.RefundAlreadyFinished)
	}

	order, err := u.adminRepo.GetOrderByID(ctx, refund.OrderID.String())
	if err != nil {
		return err
	}

	errTx := u.txRepo.WithTransaction(func(tx postgre.Transaction) error {
		refund.RefundedAt.Valid = true
		refund.RefundedAt.Time = time.Now()

		if errRefund := u.adminRepo.UpdateRefund(ctx, tx, refund); errRefund != nil {
			return errRefund
		}

		order.OrderStatusID = constant.OrderStatusRefunded
		if errStatus := u.adminRepo.UpdateOrderStatus(ctx, tx, order); errStatus != nil {
			return errStatus
		}

		orderItems, err := u.adminRepo.GetOrderItemsByOrderID(ctx, tx, order.ID.String())
		if err != nil {
			return err
		}

		for _, item := range orderItems {
			productDetailData, errData := u.adminRepo.GetProductDetailByID(ctx, tx, item.ProductDetailID.String())
			if errData != nil {
				return errData
			}

			productDetailData.Stock += float64(item.Quantity)
			errProduct := u.adminRepo.UpdateProductDetailStock(ctx, tx, productDetailData)
			if errProduct != nil {
				return errProduct
			}
		}

		totalReduce := order.TotalPrice
		if refund.IsSellerRefund != nil {
			if *refund.IsSellerRefund {
				totalReduce += order.DeliveryFee
			}
		}

		walletMarketplace, err := u.adminRepo.GetWalletByUserID(ctx, tx, constant.AdminMarketplaceID)
		if err != nil {
			return err
		}

		walletMarketplace.Balance -= totalReduce
		walletMarketplace.UpdatedAt.Valid = true
		walletMarketplace.UpdatedAt.Time = time.Now()
		if errWallet := u.adminRepo.UpdateWalletBalance(ctx, tx, walletMarketplace); errWallet != nil {
			return errWallet
		}

		walletUser, err := u.adminRepo.GetWalletByUserID(ctx, tx, order.UserID.String())
		if err != nil {
			return err
		}

		walletUser.Balance += totalReduce
		walletUser.UpdatedAt.Valid = true
		walletUser.UpdatedAt.Time = time.Now()
		if err := u.adminRepo.UpdateWalletBalance(ctx, tx, walletUser); err != nil {
			return err
		}

		walletMarketplaceHistory := &model.WalletHistory{
			TransactionID: order.TransactionID,
			WalletID:      walletMarketplace.ID,
			From:          walletMarketplace.ID.String(),
			To:            walletUser.ID.String(),
			Description:   "Refund order " + order.ID.String(),
			Amount:        totalReduce,
			CreatedAt:     time.Now(),
		}
		if err := u.adminRepo.InsertWalletHistory(ctx, tx, walletMarketplaceHistory); err != nil {
			return err
		}

		walletUserHistory := &model.WalletHistory{
			TransactionID: order.TransactionID,
			WalletID:      walletUser.ID,
			From:          walletMarketplace.ID.String(),
			To:            walletUser.ID.String(),
			Description:   "Refund order " + order.ID.String(),
			Amount:        totalReduce,
			CreatedAt:     time.Now(),
		}
		if err := u.adminRepo.InsertWalletHistory(ctx, tx, walletUserHistory); err != nil {
			return err
		}

		return nil
	})
	if errTx != nil {
		return errTx
	}

	return nil
}

func (u *adminUC) CreateVoucher(ctx context.Context, requestBody body.CreateVoucherRequest) error {
	count, _ := u.adminRepo.CountCodeVoucher(ctx, requestBody.Code)
	if count > 0 {
		return httperror.New(http.StatusBadRequest, body.CodeVoucherAlreadyExist)
	}

	voucherShop := &model.Voucher{
		Code:               requestBody.Code,
		Quota:              requestBody.Quota,
		ActivedDate:        requestBody.ActiveDateTime,
		ExpiredDate:        requestBody.ExpiredDateTime,
		DiscountPercentage: &requestBody.DiscountPercentage,
		DiscountFixPrice:   &requestBody.DiscountFixPrice,
		MinProductPrice:    &requestBody.MinProductPrice,
		MaxDiscountPrice:   &requestBody.MaxDiscountPrice,
	}

	err := u.adminRepo.CreateVoucher(ctx, voucherShop)
	if err != nil {
		return err
	}

	return nil
}

func (u *adminUC) UpdateVoucher(ctx context.Context, requestBody body.UpdateVoucherRequest) error {
	voucherShop, errVoucher := u.adminRepo.GetVoucherByID(ctx, requestBody.VoucherID)
	if errVoucher != nil {
		if errVoucher == sql.ErrNoRows {
			return httperror.New(http.StatusBadRequest, body.VoucherSellerNotFoundMessage)
		}

		return errVoucher
	}

	voucherShop.Quota = requestBody.Quota
	voucherShop.ActivedDate = requestBody.ActiveDateTime
	voucherShop.ExpiredDate = requestBody.ExpiredDateTime
	voucherShop.DiscountPercentage = &requestBody.DiscountPercentage
	voucherShop.DiscountFixPrice = &requestBody.DiscountFixPrice
	voucherShop.MinProductPrice = &requestBody.MinProductPrice
	voucherShop.MaxDiscountPrice = &requestBody.MaxDiscountPrice

	err := u.adminRepo.UpdateVoucher(ctx, voucherShop)
	if err != nil {
		return err
	}

	return nil
}

func (u *adminUC) GetDetailVoucher(ctx context.Context, voucherID string) (*model.Voucher, error) {
	voucherShop, errVoucher := u.adminRepo.GetVoucherByID(ctx, voucherID)
	if errVoucher != nil {
		if errVoucher == sql.ErrNoRows {
			return nil, httperror.New(http.StatusBadRequest, body.VoucherSellerNotFoundMessage)
		}

		return nil, errVoucher
	}

	return voucherShop, nil
}

func (u *adminUC) DeleteVoucher(ctx context.Context, voucherID string) error {
	_, errVoucher := u.adminRepo.GetVoucherByID(ctx, voucherID)
	if errVoucher != nil {
		if errVoucher == sql.ErrNoRows {
			return httperror.New(http.StatusBadRequest, body.VoucherSellerNotFoundMessage)
		}

		return errVoucher
	}

	if err := u.adminRepo.DeleteVoucher(ctx, voucherID); err != nil {
		return err
	}

	return nil
}

func (u *adminUC) GetCategories(ctx context.Context) ([]*body.CategoryResponse, error) {
	category, err := u.adminRepo.GetCategories(ctx)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (u *adminUC) AddCategory(ctx context.Context, requestBody body.CategoryRequest) error {
	err := u.adminRepo.AddCategory(ctx, requestBody)
	if err != nil {
		return err
	}
	return nil
}

func (u *adminUC) DeleteCategory(ctx context.Context, categoryID string) error {
	productCount, err := u.adminRepo.CountProductCategory(ctx, categoryID)
	if err != nil {
		return err
	}
	if productCount != 0 {
		return httperror.New(http.StatusBadRequest, body.CategoryIsBeingUsed)
	}

	categoryCount, err := u.adminRepo.CountCategoryParent(ctx, categoryID)
	if err != nil {
		return err
	}
	if categoryCount != 0 {
		return httperror.New(http.StatusBadRequest, body.CategoryIsBeingUsed)
	}

	if err := u.adminRepo.DeleteCategory(ctx, categoryID); err != nil {
		return err
	}
	return nil
}

func (u *adminUC) EditCategory(ctx context.Context, requestBody body.CategoryRequest) error {
	err := u.adminRepo.EditCategory(ctx, requestBody)
	if err != nil {
		return err
	}
	return nil
}

func (u *adminUC) GetBanner(ctx context.Context) ([]*body.BannerResponse, error) {
	banner, err := u.adminRepo.GetBanner(ctx)
	if err != nil {
		return nil, err
	}
	return banner, nil
}

func (u *adminUC) AddBanner(ctx context.Context, requestBody body.BannerRequest) error {
	err := u.adminRepo.AddBanner(ctx, requestBody)
	if err != nil {
		return err
	}
	return nil
}

func (u *adminUC) DeleteBanner(ctx context.Context, bannerID string) error {
	if err := u.adminRepo.DeleteBanner(ctx, bannerID); err != nil {
		return err
	}
	return nil
}

func (u *adminUC) EditBanner(ctx context.Context, requestBody body.BannerIDRequest) error {
	err := u.adminRepo.EditBanner(ctx, requestBody)
	if err != nil {
		return err
	}
	return nil
}
