package models

type MerchandiseAttributeValue struct {
	Id               int `json:"id" xorm:"not null pk autoincr INT(11)"`
	MerchandiseId    int `json:"merchandise_id" xorm:"not null default 0 comment('商品id') index INT(11)"`
	AttributeId      int `json:"attribute_id" xorm:"not null default 0 comment('属性id') INT(11)"`
	AttributeValueId int `json:"attribute_value_id" xorm:"not null default 0 comment('属性值id') INT(11)"`
	State            int `json:"state" xorm:"not null default 0 TINYINT(4)"`
	DeletedAt        int `json:"deleted_at" xorm:"INT(11)"`
	UpdatedAt        int `json:"updated_at" xorm:"not null default 0 INT(11)"`
	CreatedAt        int `json:"created_at" xorm:"not null default 0 INT(11)"`
}
