package product

import (
	"context"
	"murakali/internal/model"
	"murakali/internal/module/product/delivery/body"
	"murakali/pkg/pagination"
	"murakali/pkg/postgre"

	"github.com/google/uuid"
)

type Repository interface {
	GetCategories(ctx context.Context) ([]*model.Category, error)
	GetBanners(ctx context.Context) ([]*model.Banner, error)
	GetCategoriesByName(ctx context.Context, name string) ([]*model.Category, error)
	GetCategoriesByParentID(ctx context.Context, parentID uuid.UUID) ([]*model.Category, error)
	GetRecommendedProducts(ctx context.Context, pgn *pagination.Pagination) ([]*body.Products, []*model.Promotion, []*model.Voucher, error)
	GetTotalProduct(ctx context.Context) (int64, error)
	GetProductInfo(ctx context.Context, productID string) (*body.ProductInfo, error)
	GetProductDetail(ctx context.Context, productID string, promo *body.PromotionInfo) ([]*body.ProductDetail, error)
	GetAllImageByProductDetailID(ctx context.Context, productDetailID string) ([]*string, error)
	GetPromotionInfo(ctx context.Context, productID string) (*body.PromotionInfo, error)
	GetProducts(ctx context.Context, pgn *pagination.Pagination, query *body.GetProductQueryRequest) ([]*body.Products,
		[]*model.Promotion, []*model.Voucher, error)
	GetAllTotalProduct(ctx context.Context, query *body.GetProductQueryRequest) (int64, error)
	GetFavoriteProducts(ctx context.Context, pgn *pagination.Pagination, query *body.GetProductQueryRequest, userID string) ([]*body.Products,
		[]*model.Promotion, []*model.Voucher, error)
	GetAllFavoriteTotalProduct(ctx context.Context, query *body.GetProductQueryRequest, userID string) (int64, error)
	CountUserFavoriteProduct(ctx context.Context, userID, productID string) (int64, error)
	CountSpecificFavoriteProduct(ctx context.Context, productID string) (int64, error)
	CreateFavoriteProduct(ctx context.Context, tx postgre.Transaction, userID, productID string) error
	DeleteFavoriteProduct(ctx context.Context, tx postgre.Transaction, userID, productID string) error
	FindFavoriteProduct(ctx context.Context, userID, productID string) (bool, error)
	GetProductReviews(ctx context.Context, pgn *pagination.Pagination, productID string,
		query *body.GetReviewQueryRequest) ([]*body.ReviewProduct, error)
	GetTotalAllReviewProduct(ctx context.Context, productID string, query *body.GetReviewQueryRequest) (int64, error)
	GetTotalReviewRatingByProductID(ctx context.Context, productID string) ([]*body.RatingProduct, error)
	FindReview(ctx context.Context, reviewID string) (*body.ReviewProduct, error)
	CreateProductReview(ctx context.Context, tx postgre.Transaction, userID string, reqBody body.ReviewProductRequest) error
	DeleteReview(ctx context.Context, tx postgre.Transaction, reviewID string) error
	GetShopIDByUserID(ctx context.Context, userID string) (string, error)

	CreateProduct(ctx context.Context, tx postgre.Transaction, requestBody body.CreateProductInfoForQuery) (string, error)
	CreateProductDetail(ctx context.Context, tx postgre.Transaction, requestBody body.CreateProductDetailRequest, ProductID string) (string, error)
	CreatePhoto(ctx context.Context, tx postgre.Transaction, productDetailID, url string) error
	CreateVariant(ctx context.Context, tx postgre.Transaction, productDetailID string, variantDetailID string) error
	CreateVariantDetail(ctx context.Context, tx postgre.Transaction, requestBody body.VariantDetailRequest) (string, error)
	GetListedStatus(ctx context.Context, productID string) (bool, error)
	UpdateListedStatus(ctx context.Context, tx postgre.Transaction, listedStatus bool, productID string) error
	UpdateProduct(ctx context.Context, tx postgre.Transaction, requestBody body.UpdateProductInfoForQuery, productID string) error
	UpdateProductDetail(ctx context.Context, tx postgre.Transaction, requestBody body.UpdateProductDetailRequest, productID string) error
	DeletePhoto(ctx context.Context, tx postgre.Transaction, productDetailID string) error
	DeleteVariant(ctx context.Context, tx postgre.Transaction, productID string) error
	GetMaxMinPriceByID(ctx context.Context, productID string) (*body.RangePrice, error)
	UpdateVariant(ctx context.Context, tx postgre.Transaction, variantID, variantDetailID string) error
	DeleteProductDetail(ctx context.Context, tx postgre.Transaction, productDetailID string) error
	GetFavoriteProduct(ctx context.Context) ([]*model.ProductFavorite, error)
	GetRatingProduct(ctx context.Context) ([]*model.ProductRating, error)
	UpdateProductFavorite(ctx context.Context, productID string, favCount int) error
	UpdateProductRating(ctx context.Context, productID string, ratingAvg float64) error
	UpdateShopProductRating(ctx context.Context, shop *model.ShopProductRating) error
	GetShopProductRating(ctx context.Context, shopID string) (*model.ShopProductRating, error)
}
