package models

import (
	"time"
)

type User struct {
	Email       string    `json:"email" gorm:"primary_key;unique;not null"`
	Password    string    `json:"password"`
	Phone       string    `json:"phone"`
	Useraddress string    `json:"useraddress"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	VerifyCode  string    `json:"verifyCode"`
	VerifyExp   time.Time `json:"verifyExp"`
}
