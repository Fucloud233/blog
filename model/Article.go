package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	// 需要设置默认的gorm参数
	Category Category `gorm:"-"`
	Title    string   `gorm:"type: varchar(20); not null " json:"title"`
	Cid      int      `gorm:"type: int ; not null " json:"cid"`
	Desc     string   `gorm:"type: varchar(20); not null " json:"desc"`
	// 文章的内容
	Content string `gorm:"type: longtext; not null " json:"content"`
	Img     string `gorm:"type: varchar(20); " json:"img"`
}
