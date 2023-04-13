package model

import (
	"blog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	//gorm.Model
	ID   uint   `gorm:"primary_key; auto_increment" json:"id"`
	Name string `gorm:"type: varchar(20); not null " json:"name"`
}

func CheckCategory(name string) int {
	var category Category
	db.Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return errmsg.ERROR_CATEGORYNAME_USED
	}

	return errmsg.SUCCESS
}

func CheckupCategory(id int, name string) int {
	var category Category

	db.Select("id, name").Where("name = ?", name).First(&category)

	if category.ID == uint(id) {
		return errmsg.SUCCESS
	} else if category.ID > 0 {
		return errmsg.ERROR_CATEGORYNAME_USED
	}

	return errmsg.SUCCESS
}

func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}

func DeleteCategory(id int) int {
	var category Category
	err := db.Delete(&category, id).Error
	if err != nil {
		return errmsg.ERROR_CATEGORY_NOT_EXSIT
	}

	return errmsg.SUCCESS
}

func EditCategory(id int, data *Category) int {
	var category Category
	err := db.Model(&category).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR_CATEGORY_NOT_EXSIT
	}

	return errmsg.SUCCESS
}

func GetCategories(pageSize int, pageNum int) []Category {
	var categories []Category
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categories).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}

	return categories
}
