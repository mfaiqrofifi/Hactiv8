package model

import "time"

type BookDomain struct {
	ID        int       `json:"id" gorm:"primaryKey;type:serial"`
	Title     string    `json:"name_book" gorm:"type:varchar(50);unique"`
	Author    string    `json:"author" gorm:"type:varchar(50)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Division  string `json:"division" db:"division"`
}
