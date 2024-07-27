package domain

import "github.com/google/uuid"

type User struct {
	Id          uuid.UUID `gorm:"type:primaryKey;uniqueIndex;not null" json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Addresses   []Address `gorm:"foreignKey:UserId;references:Id" json:"addresses"`
}
