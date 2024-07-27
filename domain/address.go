package domain

import "github.com/google/uuid"

type Address struct {
	Id      int       `gorm:"type:int;primaryKey;uniqueIndex;not null" json:"-"`
	UserId  uuid.UUID `json:"-"`
	Street  string    `json:"street"`
	City    string    `json:"city"`
	State   string    `json:"state"`
	ZipCode string    `json:"zip_code"`
	Country string    `json:"country"`
}
