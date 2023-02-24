package models

import (
	"gin-api/global"
)

type User struct {
	Model
	Id                int    `json:"id"`
	Mobile            string `json:"mobile"`
	Nation            int    `json:"nation"`
	Avatar            string `json:"avatar"`
	Passwd            string `json:"passwd"`
	PayPasswd         string `json:"pay_passwd"`
	GoogleSecret      string `json:"google_secret"`
	IsSetGoogleSecret int    `json:"is_set_google_secret"`
	Status            int    `json:"status"`
	VerifyStatus      int    `json:"verify_status"`
	CreatedAt         int    `json:"created_at"`
	UpdatedAt         int    `json:"updated_at"`
}

func (u *User) TableName() string {
	return "user"
}

func GetUser(id int) (*User, error) {
	var user User
	global.App.DB.First(&user, id)
	return &user, global.App.DB.Error
}
