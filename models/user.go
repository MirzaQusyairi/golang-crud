package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
	ClientID int    `json:"client_id"`
}
