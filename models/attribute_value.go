package models

type AttributeValue struct {
	Id          int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	AttributeId int    `json:"attribute_id" xorm:"not null default 0 INT(11)"`
	Value       string `json:"value" xorm:"not null default '' VARCHAR(255)"`
	State       int    `json:"state" xorm:"not null default 0 TINYINT(4)"`
	DeletedAt   int    `json:"deleted_at" xorm:"INT(11)"`
	UpdatedAt   int    `json:"updated_at" xorm:"not null default 0 INT(11)"`
	CreatedAt   int    `json:"created_at" xorm:"not null default 0 INT(11)"`
}
