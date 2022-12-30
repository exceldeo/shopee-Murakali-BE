package repository

const (
	GetTotalOrderQuery = `SELECT count(id) FROM "order" WHERE "shop_id" = $1`

	GetOrdersQuery = `SELECT o.id,o.order_status_id,o.total_price,o.delivery_fee,o.resi_no,s.id,s.name,v.code,o.created_at
	from "order" o
	join "shop" s on s.id = o.shop_id
	join "voucher" v on v.id = o.voucher_shop_id WHERE o.shop_id = $1 ORDER BY o.created_at asc LIMIT $2 OFFSET $3
	`

	GetOrderDetailQuery = `SELECT pd.id,pd.product_id,p.title,ph.url,oi.quantity,oi.item_price,oi.total_price
	from  "product_detail" pd 
	join "photo" ph on pd.id = ph.product_detail_id join "order_item" oi on pd.id = oi.product_detail_id 
	join "product" p on p.id = pd.product_id WHERE oi.order_id = $1`

	GetShopIDByUserQuery = `SELECT id from shop where user_id = $1 and deleted_at is null`
)