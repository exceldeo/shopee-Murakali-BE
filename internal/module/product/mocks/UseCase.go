// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	body "murakali/internal/module/product/delivery/body"

	mock "github.com/stretchr/testify/mock"

	model "murakali/internal/model"

	pagination "murakali/pkg/pagination"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// CheckProductIsFavorite provides a mock function with given fields: ctx, userID, productID
func (_m *UseCase) CheckProductIsFavorite(ctx context.Context, userID string, productID string) bool {
	ret := _m.Called(ctx, userID, productID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, userID, productID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CountSpecificFavoriteProduct provides a mock function with given fields: ctx, productID
func (_m *UseCase) CountSpecificFavoriteProduct(ctx context.Context, productID string) (int64, error) {
	ret := _m.Called(ctx, productID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, productID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFavoriteProduct provides a mock function with given fields: ctx, productID, userID
func (_m *UseCase) CreateFavoriteProduct(ctx context.Context, productID string, userID string) error {
	ret := _m.Called(ctx, productID, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, productID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateProduct provides a mock function with given fields: ctx, requestBody, userID
func (_m *UseCase) CreateProduct(ctx context.Context, requestBody body.CreateProductRequest, userID string) error {
	ret := _m.Called(ctx, requestBody, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, body.CreateProductRequest, string) error); ok {
		r0 = rf(ctx, requestBody, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateProductReview provides a mock function with given fields: ctx, reqBody, userID
func (_m *UseCase) CreateProductReview(ctx context.Context, reqBody body.ReviewProductRequest, userID string) error {
	ret := _m.Called(ctx, reqBody, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, body.ReviewProductRequest, string) error); ok {
		r0 = rf(ctx, reqBody, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteFavoriteProduct provides a mock function with given fields: ctx, productID, userID
func (_m *UseCase) DeleteFavoriteProduct(ctx context.Context, productID string, userID string) error {
	ret := _m.Called(ctx, productID, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, productID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteProductReview provides a mock function with given fields: ctx, reviewID, userID
func (_m *UseCase) DeleteProductReview(ctx context.Context, reviewID string, userID string) error {
	ret := _m.Called(ctx, reviewID, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, reviewID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllProductImage provides a mock function with given fields: ctx, productID
func (_m *UseCase) GetAllProductImage(ctx context.Context, productID string) ([]*body.GetImageResponse, error) {
	ret := _m.Called(ctx, productID)

	var r0 []*body.GetImageResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) []*body.GetImageResponse); ok {
		r0 = rf(ctx, productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*body.GetImageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBanners provides a mock function with given fields: ctx
func (_m *UseCase) GetBanners(ctx context.Context) ([]*model.Banner, error) {
	ret := _m.Called(ctx)

	var r0 []*model.Banner
	if rf, ok := ret.Get(0).(func(context.Context) []*model.Banner); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Banner)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCategories provides a mock function with given fields: ctx
func (_m *UseCase) GetCategories(ctx context.Context) ([]*body.CategoryResponse, error) {
	ret := _m.Called(ctx)

	var r0 []*body.CategoryResponse
	if rf, ok := ret.Get(0).(func(context.Context) []*body.CategoryResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*body.CategoryResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCategoriesByName provides a mock function with given fields: ctx, name
func (_m *UseCase) GetCategoriesByName(ctx context.Context, name string) ([]*body.CategoryResponse, error) {
	ret := _m.Called(ctx, name)

	var r0 []*body.CategoryResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) []*body.CategoryResponse); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*body.CategoryResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFavoriteProducts provides a mock function with given fields: ctx, pgn, query, userID
func (_m *UseCase) GetFavoriteProducts(ctx context.Context, pgn *pagination.Pagination, query *body.GetProductQueryRequest, userID string) (*pagination.Pagination, error) {
	ret := _m.Called(ctx, pgn, query, userID)

	var r0 *pagination.Pagination
	if rf, ok := ret.Get(0).(func(context.Context, *pagination.Pagination, *body.GetProductQueryRequest, string) *pagination.Pagination); ok {
		r0 = rf(ctx, pgn, query, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pagination.Pagination)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pagination.Pagination, *body.GetProductQueryRequest, string) error); ok {
		r1 = rf(ctx, pgn, query, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductDetail provides a mock function with given fields: ctx, productID
func (_m *UseCase) GetProductDetail(ctx context.Context, productID string) (*body.ProductDetailResponse, error) {
	ret := _m.Called(ctx, productID)

	var r0 *body.ProductDetailResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) *body.ProductDetailResponse); ok {
		r0 = rf(ctx, productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*body.ProductDetailResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductReviews provides a mock function with given fields: ctx, pgn, productID, query
func (_m *UseCase) GetProductReviews(ctx context.Context, pgn *pagination.Pagination, productID string, query *body.GetReviewQueryRequest) (*pagination.Pagination, error) {
	ret := _m.Called(ctx, pgn, productID, query)

	var r0 *pagination.Pagination
	if rf, ok := ret.Get(0).(func(context.Context, *pagination.Pagination, string, *body.GetReviewQueryRequest) *pagination.Pagination); ok {
		r0 = rf(ctx, pgn, productID, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pagination.Pagination)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pagination.Pagination, string, *body.GetReviewQueryRequest) error); ok {
		r1 = rf(ctx, pgn, productID, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProducts provides a mock function with given fields: ctx, pgn, query
func (_m *UseCase) GetProducts(ctx context.Context, pgn *pagination.Pagination, query *body.GetProductQueryRequest) (*pagination.Pagination, error) {
	ret := _m.Called(ctx, pgn, query)

	var r0 *pagination.Pagination
	if rf, ok := ret.Get(0).(func(context.Context, *pagination.Pagination, *body.GetProductQueryRequest) *pagination.Pagination); ok {
		r0 = rf(ctx, pgn, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pagination.Pagination)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pagination.Pagination, *body.GetProductQueryRequest) error); ok {
		r1 = rf(ctx, pgn, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecommendedProducts provides a mock function with given fields: ctx, pgn
func (_m *UseCase) GetRecommendedProducts(ctx context.Context, pgn *pagination.Pagination) (*pagination.Pagination, error) {
	ret := _m.Called(ctx, pgn)

	var r0 *pagination.Pagination
	if rf, ok := ret.Get(0).(func(context.Context, *pagination.Pagination) *pagination.Pagination); ok {
		r0 = rf(ctx, pgn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pagination.Pagination)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pagination.Pagination) error); ok {
		r1 = rf(ctx, pgn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalReviewRatingByProductID provides a mock function with given fields: ctx, productID
func (_m *UseCase) GetTotalReviewRatingByProductID(ctx context.Context, productID string) (*body.AllRatingProduct, error) {
	ret := _m.Called(ctx, productID)

	var r0 *body.AllRatingProduct
	if rf, ok := ret.Get(0).(func(context.Context, string) *body.AllRatingProduct); ok {
		r0 = rf(ctx, productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*body.AllRatingProduct)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateListedStatus provides a mock function with given fields: ctx, productID
func (_m *UseCase) UpdateListedStatus(ctx context.Context, productID string) error {
	ret := _m.Called(ctx, productID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, productID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProduct provides a mock function with given fields: ctx, requestBody, userID, productID
func (_m *UseCase) UpdateProduct(ctx context.Context, requestBody body.UpdateProductRequest, userID string, productID string) error {
	ret := _m.Called(ctx, requestBody, userID, productID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, body.UpdateProductRequest, string, string) error); ok {
		r0 = rf(ctx, requestBody, userID, productID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProductListedStatusBulk provides a mock function with given fields: ctx, _a1
func (_m *UseCase) UpdateProductListedStatusBulk(ctx context.Context, _a1 body.UpdateProductListedStatusBulkRequest) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, body.UpdateProductListedStatusBulkRequest) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProductMetadata provides a mock function with given fields: ctx
func (_m *UseCase) UpdateProductMetadata(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCase(t mockConstructorTestingTNewUseCase) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
