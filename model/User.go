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
	// Avatar string
}

func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)

	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}

	return errmsg.SUCCESS
}

func CheckupUser(id int, name string) (code int) {
	var user User
	db.Select("id, username").Where("username = ?", name).First(&user)

	// 当该用户名是其本身 返回SUCCESS
	if user.ID == uint(id) {
		return errmsg.SUCCESS
	} else if user.ID > 0 {
		// 该用户名已经被使用
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// CreateUser 新增用户 返回一个code
func CreateUser(data *User) int {
	// 添加密码加
	data.Password = bcryptPw(data.Password)

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

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Delete(&user, id).Error
	if err != nil {
		return errmsg.ERROR_USER_NOT_EXSIT
	}

	return errmsg.SUCCESS
}

// EditUser 修改用户
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["role"] = data.Role
	maps["username"] = data.Username
	err = db.Model(&user).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// bcryptPw 密码加密
// https://pkg.go.dev/golang.org/x/crypto/bcrypt
func bcryptPw(password string) string {
	const cost = 10
	hashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(string(hashPw)))
	return string(hashPw)
}
