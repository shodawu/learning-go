package entities

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserStatus uint

const UserRegistered UserStatus = 101
const UserConfirmed UserStatus = 102
const UserBlocked UserStatus = 103

type User struct {
	gorm.Model
	Name   string
	Email  string
	Status UserStatus
}

func InitDB() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	var users []User

	db.Model(&users).Find(&users)

	if len(users) == 0 {
		users := []User{
			{Name: "UserOne", Email: "user.one@example.com", Status: UserConfirmed},
			{Name: "UserTwo", Email: "user.two@example.com", Status: UserConfirmed},
			{Name: "UserThree", Email: "user.three@example.com", Status: UserConfirmed},
			{Name: "UserFour", Email: "user.four@example.com", Status: UserBlocked},
			{Name: "UserFive", Email: "user.five@example.com", Status: UserRegistered},
		}

		for _, usr := range users {
			er := db.Create(&usr).Error
			if er != nil {
				fmt.Println("creating users error ", er)
			}
		}
	}

	return db
}
