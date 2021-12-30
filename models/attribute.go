package models

type Attribute struct {
	Id        int    `json:"id" xorm:"not null pk autoincr comment('属性id') INT(11)"`
	Name      string `json:"name" xorm:"not null default '' comment('属性名称') VARCHAR(255)"`
	State     int    `json:"state" xorm:"not null comment('状态') TINYINT(4)"`
	DeletedAt int    `json:"deleted_at" xorm:"INT(11)"`
	UpdatedAt int    `json:"updated_at" xorm:"not null INT(11)"`
	CreatedAt int    `json:"created_at" xorm:"not null INT(11)"`
}
