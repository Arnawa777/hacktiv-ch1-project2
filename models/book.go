package models

import "time"

type Book struct {
	ID          uint   `gorm:"primary_key"`
	Title       string `gorm:"not null; type=varchar(100)"`
	Author      string `gorm:"not null; type=varchar(100)"`
	Description string `gorm:"type=text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
