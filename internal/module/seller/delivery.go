package seller

import "github.com/gin-gonic/gin"

type Handlers interface {
	GetOrder(c *gin.Context)
	ChangeOrderStatus(c *gin.Context)
	GetOrderByOrderID(c *gin.Context)
	GetCourierSeller(c *gin.Context)
	GetSellerBySellerID(c *gin.Context)
	GetSellerByUserID(c *gin.Context)
	CreateCourierSeller(c *gin.Context)
	DeleteCourierSellerByID(c *gin.Context)
	GetCategoryBySellerID(c *gin.Context)
	UpdateResiNumberInOrderSeller(c *gin.Context)
	GetAllVoucherSeller(c *gin.Context)
	CreateVoucherSeller(c *gin.Context)
	UpdateVoucherSeller(c *gin.Context)
	DeleteVoucherSeller(c *gin.Context)
	DetailVoucherSeller(c *gin.Context)
	GetAllPromotionSeller(c *gin.Context)
	CreatePromotionSeller(c *gin.Context)
	UpdatePromotionSeller(c *gin.Context)
	GetDetailPromotionSellerByID(c *gin.Context)
	UpdateOnDeliveryOrder(c *gin.Context)
	UpdateExpiredAtOrder(c *gin.Context)
}
