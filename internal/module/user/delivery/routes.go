package delivery

import (
	"murakali/internal/middleware"
	"murakali/internal/module/user"

	"github.com/gin-gonic/gin"
)

func MapUserRoutes(userGroup *gin.RouterGroup, h user.Handlers, mw *middleware.MWManager) {
	userGroup.POST("/transaction/slp-payment/:id", h.SLPPaymentCallback)
	userGroup.POST("/transaction/wallet-payment/:id", h.WalletPaymentCallback)
	userGroup.Use(mw.AuthJWTMiddleware())
	userGroup.GET("/address", h.GetAddress)
	userGroup.POST("/address", h.CreateAddress)
	userGroup.GET("/address/:id", h.GetAddressByID)
	userGroup.PUT("/address/:id", h.UpdateAddressByID)
	userGroup.DELETE("/address/:id", h.DeleteAddressByID)
	userGroup.PUT("/profile", h.EditUser)
	userGroup.POST("/email", h.EditEmail)
	userGroup.GET("/email", h.EditEmailUser)
	userGroup.GET("/sealab-pay", h.GetSealabsPay)
	userGroup.POST("/sealab-pay", h.AddSealabsPay)
	userGroup.PATCH("/sealab-pay/:cardNumber", h.PatchSealabsPay)
	userGroup.DELETE("/sealab-pay/:cardNumber", h.DeleteSealabsPay)
	userGroup.POST("/register-merchant", h.RegisterMerchant)
	userGroup.GET("/profile", h.GetUserProfile)
	userGroup.POST("/profile/picture", h.UploadProfilePicture)
	userGroup.POST("/password", h.VerifyPasswordChange)
	userGroup.POST("/verify", h.VerifyOTP)
	userGroup.PATCH("/password", h.ChangePassword)
	userGroup.GET("/transaction/detail/:transaction_id", h.GetTransactionDetailByID)
	userGroup.GET("/transaction", h.GetTransactions)
	userGroup.GET("/transaction/:id", h.GetTransaction)
	userGroup.POST("/transaction", h.CreateTransaction)
	userGroup.POST("/transaction/slp-payment", h.CreateSLPPayment)
	userGroup.POST("/transaction/wallet-payment", h.CreateWalletPayment)
	userGroup.GET("/order", h.GetOrder)
	userGroup.GET("/order/:order_id", h.GetOrderByOrderID)
	userGroup.POST("/wallet", h.ActivateWallet)
	userGroup.GET("/wallet", h.GetWallet)
	userGroup.GET("/wallet/history", h.GetWalletHistory)
	userGroup.GET("/wallet/history/:wallet_history_id", h.GetWalletHistoryByID)
	userGroup.PATCH("/wallet", h.TopUpWallet)
	userGroup.POST("/wallet/step-up/pin", h.WalletStepUp)
	userGroup.POST("/wallet/step-up/password", h.ChangeWalletPinStepUp)
	userGroup.PATCH("/wallet/pin", h.ChangeWalletPin)
}
