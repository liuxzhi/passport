package models

type MerchandiseItemBusiness struct {
	Id                int64  `json:"id" xorm:"pk autoincr BIGINT(20)"`
	BusinessId        int    `json:"business_id" xorm:"not null default 0 INT(11)"`
	MerchandiseId     int    `json:"merchandise_id" xorm:"not null default 0 comment('商品id(SPU)') INT(11)"`
	MerchandiseItemId int    `json:"merchandise_item_id" xorm:"not null default 0 comment('商品单品id(SKU)') INT(11)"`
	MerchandiseNo     string `json:"merchandise_no" xorm:"not null default '' comment('商品货号') VARCHAR(100)"`
	State             int    `json:"state" xorm:"not null default 0 comment('上下架状态') TINYINT(4)"`
	PlatformPrice     int    `json:"platform_price" xorm:"not null default 0 comment('平台价') INT(11)"`
	Price             int    `json:"price" xorm:"not null default 0 comment('商品价格') INT(11)"`
	DeletedAt         int    `json:"deleted_at" xorm:"INT(11)"`
	CreatedAt         int    `json:"created_at" xorm:"not null default 0 INT(11)"`
	UpdatedAt         int    `json:"updated_at" xorm:"not null default 0 INT(11)"`
}
