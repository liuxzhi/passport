package models

type MerchandiseItem struct {
	Id                int    `json:"id" xorm:"not null pk autoincr comment('商品单品id') INT(11)"`
	MerchandiseId     int    `json:"merchandise_id" xorm:"not null default 0 comment('商品id') index(idx_merchandise_attributeValueIds) INT(11)"`
	MerchandiseNo     string `json:"merchandise_no" xorm:"not null default '' VARCHAR(255)"`
	AttributeIds      string `json:"attribute_ids" xorm:"not null default '' VARCHAR(255)"`
	AttributeValueIds string `json:"attribute_value_ids" xorm:"not null default '' index(idx_merchandise_attributeValueIds) VARCHAR(255)"`
	Storage           int    `json:"storage" xorm:"not null default 0 INT(10)"`
	Name              string `json:"name" xorm:"not null default '' comment('商品单品名称') VARCHAR(255)"`
	Image             string `json:"image" xorm:"not null default '' comment('单品图片') VARCHAR(255)"`
	SkuId             string `json:"sku_id" xorm:"not null default '' comment('第三方sku_id') VARCHAR(255)"`
	DeletedAt         int    `json:"deleted_at" xorm:"comment('删除时间戳') INT(11)"`
	UpdatedAt         int    `json:"updated_at" xorm:"not null default 0 comment('更新时间戳') INT(11)"`
	CreatedAt         int    `json:"created_at" xorm:"not null default 0 comment('创建时间戳') INT(11)"`
}
