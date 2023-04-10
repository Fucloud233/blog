package model

import (
	"blog/utils/errmsg"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type: varchar(20); not null " json:"username"`
	// 为了保证能够存加密值
	Password string `gorm:"type: varchar(60); not null" json:"password"`
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
	// 添加密码加
	data.Password = scrpyPw(data.Password)

	err := db.Create(&data).Error

	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}

func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}

	return users
}

// bcrpyPw 密码加密
// https://pkg.go.dev/golang.org/x/crypto/bcrypt
func bcrpyPw(password string) string {
	const cost = 10
	hashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(string(hashPw)))
	return string(hashPw)
}
