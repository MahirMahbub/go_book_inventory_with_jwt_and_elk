package models

import (
	"go_practice/book/structs"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string    `json:"title" gorm:"unique"`
	UserID      uint      `json:"userId" gorm:"default: null"`
	Description string    `gorm:"size:6000" json:"description"`
	Authors     []*Author `gorm:"many2many:authors_books;" json:"author"`
}

type Books []Book

func (book *Book) GetUserBookByID(ID uint, userID uint) (err error) {
	return DB.Where("id = ? AND user_id = ?", ID, userID).First(&book).Error
}

func (book *Book) GetBookByID(ID uint) (err error) {
	return DB.Where("id = ?", ID).First(&book).Error
}

func (books *Books) GetUserBooksBySelection(userID uint, selection []string) *gorm.DB {
	return DB.Where("user_id = ?", userID).Select(selection).Find(&books)
}

func (books *Books) GetBooksBySelection(selection []string) *gorm.DB {
	return DB.Select(selection).Find(&books)
}

func (book *Book) GetUserBookWithAuthors(ID uint, userID uint) (err error) {
	return DB.Preload("Authors").Where("id = ? AND user_id = ?", ID, userID).First(&book).Error
}

func (book *Book) GetBookWithAuthors(ID uint) (err error) {
	return DB.Preload("Authors").Where("id = ?", ID).First(&book).Error
}

func (book *Book) CreateBookWithAuthors(authors []Author) (err error) {
	return DB.Create(&book).Association("Authors").Append(authors)
}

func (book *Book) UpdateBook(input structs.UpdateBookInput) (err error) {
	if input.Title != "" {
		book.Title = input.Title
	}
	if input.Description != "" {
		book.Description = input.Description
	}
	if err := DB.Save(&book).Error; err != nil {
		return err
	}
	return nil
}

func (book *Book) DeleteBook() (err error) {
	return DB.Delete(&book).Error
}
