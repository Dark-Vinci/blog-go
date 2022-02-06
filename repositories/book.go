package repositories

import (
	entity "new-proj/entities"

	"gorm.io/gorm"
)

type BookRepository interface {
	InsertBook(b entity.Book) entity.Book
	UpdateBook(b entity.Book) entity.Book
	DeleteBook(b entity.Book)
	AllBooks() [] entity.Book
	FindById(bookID uint64) entity.Book
}

type bookRepository struct {
	bcon *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		bcon: db,
	}
}

func (bk *bookRepository) InsertBook(b entity.Book) entity.Book {
	bk.bcon.Save(&b)
	bk.bcon.Preload("User").Find(&b)

	return b
}

func (bk *bookRepository) UpdateBook(b entity.Book) entity.Book {
	bk.bcon.Save(&b)
	bk.bcon.Preload("User").Find(&b)

	return b
}

func (bk *bookRepository) DeleteBook(b entity.Book) {
	bk.bcon.Delete(&b)
}

func (bk *bookRepository) AllBooks() [] entity.Book {
	var books [] entity.Book

	bk.bcon.Preload("User").Find(&books)

	return books
}

func (bk *bookRepository) FindById(bookID uint64) entity.Book {
	var book entity.Book

	bk.bcon.Preload("User").Find(&book, bookID)

	return book
}