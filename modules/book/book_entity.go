package book

import "gorm.io/gorm"

func (Book) TableName() string {
	return "m_book"
}

type Book struct {
	gorm.Model
	Code        string `gorm:"column:code"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
}
