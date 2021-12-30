package models

type MerchandiseBusiness struct {
	Id            int64 `json:"id" xorm:"pk autoincr BIGINT(20)"`
	BusinessId    int   `json:"business_id" xorm:"not null comment('业务线id') INT(11)"`
	MerchandiseId int   `json:"merchandise_id" xorm:"not null INT(11)"`
	State         int   `json:"state" xorm:"not null TINYINT(4)"`
	DeletedAt     int   `json:"deleted_at" xorm:"TINYINT(4)"`
	UpdatedAt     int   `json:"updated_at" xorm:"not null default 0 INT(11)"`
	CreatedAt     int   `json:"created_at" xorm:"not null INT(11)"`
}
