package repository

const (
	GetTotalAddressQuery          = `SELECT count(id) FROM "address" WHERE "user_id" = $1 AND "name" ILIKE $2 AND "deleted_at" IS NULL`
	GetDefaultAddressQuery        = `SELECT "id", "user_id", "is_default" FROM "address" WHERE "user_id" = $1 AND "is_default" = $2 AND "deleted_at" IS NULL`
	GetDefaultShopAddressQuery    = `SELECT "id", "user_id", "is_shop_default" FROM "address" WHERE "user_id" = $1 AND "is_shop_default" = $2 AND "deleted_at" IS NULL`
	UpdateDefaultAddressQuery     = `UPDATE "address" SET "is_default" = $1 WHERE "id" = $2`
	UpdateDefaultShopAddressQuery = `UPDATE "address" SET "is_shop_default" = $1 WHERE "id" = $2`
	DeleteAddressByIDQuery        = `DELETE FROM "address" WHERE "id" = $1`

	CreateAddressQuery = `INSERT INTO "address" 
    	(user_id, name, province_id, city_id, province, city, district, sub_district, address_detail, zip_code, is_default, is_shop_default)
    	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	GetAddressesQuery = `SELECT 
    	"id", "user_id", "name", "province_id", "city_id", "province", "city", "district", "sub_district",  
    	"address_detail", "zip_code", "is_default", "is_shop_default", "created_at", "updated_at" 
	FROM "address" WHERE "user_id" = $1 AND "name" ILIKE $2 AND "deleted_at" IS NULL ORDER BY $3 LIMIT $4 OFFSET $5`

	GetAddressByIDQuery = `SELECT
		"id", "user_id", "name", "province_id", "city_id", "province", "city", "district", "sub_district",  
    	"address_detail", "zip_code", "is_default", "is_shop_default", "created_at", "updated_at"
	FROM "address" WHERE "id" = $1 AND "user_id" = $2 AND "deleted_at" IS NULL`

	UpdateAddressByIDQuery = `UPDATE "address" SET
		"name" = $1, "province_id" = $2, "city_id" = $3, "province" = $4, "city" = $5, "district" = $6,
		"sub_district" = $7, "address_detail" = $8, "zip_code" = $9, "is_default" = $10, "is_shop_default" = $11, "updated_at" = $12
	WHERE "id" = $13`

	GetSealabsPayByIdQuery         = `SELECT * from sealabs_pay where user_id = $1 and deleted_at is null`
	CreateSealabsPayQuery          = `INSERT INTO "sealabs_pay" (card_number, user_id, name, is_default,active_date) VALUES ($1, $2, $3, $4, $5)`
	CheckDefaultSealabsPayQuery    = `SELECT card_number from "sealabs_pay" where user_id = $1 and is_default is true and deleted_at is null`
	SetDefaultSealabsPayTransQuery = `UPDATE "sealabs_pay" set is_default = FALSE,updated_at = now() where card_number = $1`
	PatchSealabsPayQuery           = `UPDATE "sealabs_pay" set is_default = TRUE,updated_at = now() where card_number = $1`
	SetDefaultSealabsPayQuery      = `UPDATE "sealabs_pay" set is_default = FALSE where card_number <> $1 and user_id = $2`
	DeleteSealabsPayQuery          = `UPDATE "sealabs_pay" set deleted_at = now() where card_number = $1 and is_default = FALSE`
	GetUserByIDQuery                            = `SELECT "id", "role_id", "email", "username", "phone_no", "fullname", "gender", "birth_date", "is_verify" FROM "user" WHERE "id" = $1 AND "deleted_at" IS NULL`
	CheckEmailHistoryQuery         = `SELECT "id", "email" FROM "email_history" WHERE "email" ILIKE $1`
	GetUserByUsernameQuery         = `SELECT "id", "email", "username", "is_verify" FROM "user" WHERE "username" ILIKE $1`
	GetUserByPhoneNoQuery          = `SELECT "id", "email", "phone_no", "is_verify" FROM "user" WHERE "phone_no" ILIKE $1`
	UpdateUserFieldQuery           = `UPDATE "user" SET "username" = $1, "fullname" = $2, "phone_no" = $3, "birth_date" = $4, "gender" = $5, "updated_at" = $6 WHERE "email" = $7`
	UpdateUserEmailQuery           = `UPDATE "user" SET "email" = $1, "updated_at" = $2 WHERE "id" = $3`
	CreateEmailHistoryQuery        = `INSERT INTO "email_history" (email) VALUES ($1)`
)