package models

type Merchandise struct {
	Id           int    `json:"id" xorm:"not null pk autoincr comment('商品id') INT(11)"`
	Name         string `json:"name" xorm:"not null default '' comment('商品名称') VARCHAR(255)"`
	Introduction string `json:"introduction" xorm:"not null default '' comment('商品价格') VARCHAR(255)"`
	State        int    `json:"state" xorm:"not null default 0 comment('上下架状态') TINYINT(4)"`
	Sort         int    `json:"sort" xorm:"not null default 0 comment('排序值') TINYINT(4)"`
	DeletedAt    int    `json:"deleted_at" xorm:"default 0 INT(11)"`
	UpdatedAt    int    `json:"updated_at" xorm:"not null default 0 INT(11)"`
	CreatedAt    int    `json:"created_at" xorm:"not null default 0 INT(11)"`
}
