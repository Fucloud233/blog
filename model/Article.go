package model

import (
	"blog/utils/errmsg"
	"gorm.io/gorm"
)

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

// todo 查询分类下的所有文章

// todo 查询单个文章

// todo 查询文章列表

func CreateArticle(data *Article) int {
	err := db.Create(data)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteArticle(id int) int {
	var article Article
	err := db.Delete(&article, id).Error
	if err != nil {
		return errmsg.ERROR_Article_NOT_EXSIT
	}
	return errmsg.SUCCESS
}

func EditArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{})
	//maps["category"] = data.Category
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err := db.Model(&article).Where("id = ?", id).Updates(&maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
