package models

type MerchandiseItemAttribute struct {
	Id            int `json:"id" xorm:"not null pk autoincr comment('商品单品id') INT(11)"`
	MerchandiseId int `json:"merchandise_id" xorm:"not null default 0 comment('商品id') index INT(11)"`
	ItemId        int `json:"item_id" xorm:"not null default 0 comment('商品单品id') INT(11)"`
	AttributeId   int `json:"attribute_id" xorm:"not null default 0 comment('属性id') INT(11)"`
	State         int `json:"state" xorm:"not null default 0 comment('状态') TINYINT(4)"`
	DeletedAt     int `json:"deleted_at" xorm:"comment('是否删除') INT(11)"`
	UpdatedAt     int `json:"updated_at" xorm:"not null default 0 comment('修改的时间戳') INT(11)"`
	CreatedAt     int `json:"created_at" xorm:"not null default 0 comment('创建的时间戳') INT(11)"`
}
