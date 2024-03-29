package constant

const (
	AdminMarketplaceID = "4df967a8-5b05-4d2a-bb72-da3921dce8fb"

	RegisterTokenCookie        = "register_token"
	RefreshTokenCookie         = "refresh_token"
	WalletTokenCookie          = "wallet_token"
	ChangeWalletPinTokenCookie = "change_wallet_pin_token"
	ResetPasswordTokenCookie   = "reset_password_token"
	ChangePasswordTokenCookie  = "change_password_token"

	ProvinceKey    = "location:province"
	CityKey        = "location:city"
	SubDistrictKey = "location:subdistrict"
	UrbanKey       = "location:urban"
	OtpKey         = "user:otp"
	OtpDuration    = "30m"
	AddressDefault = "true"

	RoleUser   = 1
	RoleSeller = 2
	RoleAdmin  = 3

	ImgMaxSize = 500000

	SLPStatusPaid      = "TXN_PAID"
	SlPMessagePaid     = "Payment successful"
	SLPStatusCanceled  = "TXN_FAILED"
	SLPMessageCanceled = "Transaction Canceled by user"

	TRUE  = "true"
	FALSE = "false"
	ASC   = "asc"
	DESC  = "desc"

	LoginOauth     = "/login"
	RegisterOauth  = "/register"
	ONGKIR_API_URL = "https://api.rajaongkir.com/starter"
	KODE_POS_URL   = "https://kode-pos-murakali.vercel.app/"
	ONGKIR_API_KEY = "9bc050f97d720a0477556cfa4d8f6de5"

	OrderStatusWaitingToPay     = 1
	OrderStatusWaitingForSeller = 2
	OrderStatusProcessed        = 3
	OrderStatusOnDelivery       = 4
	OrderStatusDelivered        = 5
	OrderStatusReceived         = 6
	OrderStatusCompleted        = 7
	OrderStatusCanceled         = 8
	OrderStatusRefunded         = 9
)
