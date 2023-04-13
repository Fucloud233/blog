package model

import (
	"blog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	// 需要设置默认的gorm参数
	// 使用gorm中的外键 foreignKey (关联关系)
	Category Category `gorm:"foreignKey:Cid"`
	Title    string   `gorm:"type: varchar(20); not null " json:"title"`
	Cid      int      `gorm:"type: int ; not null " json:"cid"`
	Desc     string   `gorm:"type: varchar(20); not null " json:"desc"`
	// 文章的内容
	Content string `gorm:"type: longtext; not null " json:"content"`
	Img     string `gorm:"type: varchar(20); " json:"img"`
}

// GetArticleFromCategory 查询分类下的所有文章
func GetArticleFromCategory(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var articles []Article
	var total int64

	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid=?", id).Find(&articles).Error
	if err != nil {
		return nil, errmsg.ERROR_CATEGORY_NOT_EXSIT, 0
	}
	db.Model(&articles).Where("cid=?", id).Count(&total)
	return articles, errmsg.SUCCESS, total
}

// GetArticle 查询单个文章
func GetArticle(id int) (Article, int) {
	var article Article
	err := db.Preload("Category").First(&article, id).Error
	if err != nil {
		return article, errmsg.ERROR_Article_NOT_EXSIT
	}
	return article, errmsg.SUCCESS
}

// GetArticles 查询文章列表
func GetArticles(pageSize int, pageNum int) ([]Article, int) {
	var articles []Article
	// Preload 使用预加载
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}

	return articles, errmsg.SUCCESS
}

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
