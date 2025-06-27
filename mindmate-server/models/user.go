/*
 * Define database operations of users
 */

package models

import (
	"github.com/noahzyl/mindmate/dao"
	"time"
)

// User struct is consistent with the user table in database, and used to save a user's data
type User struct {
	Id         int // Primary key, assigned by MySQL automatically
	Username   string
	Password   string
	AddTime    time.Time
	UpdateTime time.Time
}

// Set the table name which gorm will use
func (User) TableName() string {
	// gorm will use "user" as the specific table name when functions in models/user.go are executed
	return "user"
}

func GetUserInfoByUsername(username string) (User, error) {
	var user User
	err := dao.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

func AddUser(username string, password string) (int, error) {
	user := User{
		Username:   username,
		Password:   password,
		AddTime:    time.Now(),
		UpdateTime: time.Now(),
	}
	err := dao.DB.Create(&user).Error
	return user.Id, err
}
