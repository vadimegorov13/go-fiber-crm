package models

import (
	"gorm.io/gorm"
)

/*
gorm.model will add properties such as ID, CreatedAt, UpdatedAt and DeletedAt
*/

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}
