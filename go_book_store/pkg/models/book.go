package models

import (
	"example/go_bookstore/pkg/config"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name" isRequired:"required"`
	Author      string `json:"author" isRequired:"required"`
	Publication string `json:"publication" isRequired:"required"`
}

func init() {
	config.Connect()
	DB = config.GetDB()
	DB.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, error) {

	DB.NewRecord(b)

	if err := DB.Create(&b).Error; err != nil {
		return &Book{}, err
	}
	return b, nil
}

func GetAllBooks() (*[]Book, error) {
	var books []Book
	if err := DB.Find(&books).Error; err != nil {
		return &[]Book{}, err
	}
	return &books, nil
}
func GetBookById(id int64) (*Book, error) {
	var book Book
	if err := DB.Where("id = ?", id).First(&book).Error; err != nil {
		return &book, err
	}
	return &book, nil
}
func DeleteBook(id int64) (*Book, error) {
	var book Book
	if err := DB.Where("id = ?", id).Delete(&book).Error; err != nil {
		return &book, err
	}
	return &book, nil
}
