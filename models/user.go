package models

type User struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone" gorm:"unique"`
	Password []byte `json:"-"`
}