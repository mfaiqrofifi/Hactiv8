package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type BookDomain struct {
	ID        int       `json:"id" gorm:"primaryKey;type:serial"`
	Title     string    `json:"name_book" gorm:"type:varchar(50);unique"`
	Author    string    `json:"author" gorm:"type:varchar(50)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Division  string `json:"division" db:"division"`
}

func (e BookDomain) Validation() error { // custom validation
	return validation.ValidateStruct(&e,
		validation.Field(&e.Title, validation.Required),
		validation.Field(&e.Author, validation.Required),
	)
}
