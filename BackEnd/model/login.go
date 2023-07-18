package model

type Login struct {
	User_id       string `json:"user_id" gorm:"column:user_id"`
	User_password string `json:"user_password" gorm:"column:user_id"`
}
