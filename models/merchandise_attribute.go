package models

type MerchandiseAttribute struct {
	Id            int `json:"id" xorm:"not null pk autoincr INT(11)"`
	MerchandiseId int `json:"merchandise_id" xorm:"not null default 0 comment('商品id') index(idx_merchandiseId_attribute_id) INT(11)"`
	AttributeId   int `json:"attribute_id" xorm:"not null default 0 index(idx_merchandiseId_attribute_id) INT(11)"`
	IsPrime       int `json:"is_prime" xorm:"not null default 0 comment('是否主属性') TINYINT(4)"`
	State         int `json:"state" xorm:"not null default 0 TINYINT(4)"`
	DeletedAt     int `json:"deleted_at" xorm:"INT(11)"`
	UpdatedAt     int `json:"updated_at" xorm:"not null default 0 INT(11)"`
	CreatedAt     int `json:"created_at" xorm:"not null default 0 INT(11)"`
}
