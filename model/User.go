package model

import (
	"blog/utils/errmsg"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type: varchar(20); not null " json:"username"`
	Password string `gorm:"type: varchar(20); not null" json:"password"`
	Role     int    `gorm:"type: int" json:"role"`
	// Avator string
}

func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)

	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}

	return errmsg.SUCCESS
}

// 新增用户 返回一个code
func CreateUser(data *User) int {
	err := db.Create(&data).Error

	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}
