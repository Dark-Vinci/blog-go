package service

import (
	"log"
	"fmt"
	
	"new-proj/dto"
	entity "new-proj/entities"
	"new-proj/repositories"

	"github.com/mashingan/smapping"
)

type BookService interface {
	InsertBook(b dto.BookCreateDTO) entity.Book
	UpdateBook(b dto.BookUpdateDTO) entity.Book
	DeleteBook(b entity.Book)
	AllBooks() []entity.Book
	FindById(bookID uint64) entity.Book
	IsAllowedToEdit(userID string, bookID uint64) bool
}


type bookService struct {
	bookRepo repositories.BookRepository
}

func NewBookService(bookRepo repositories.BookRepository) BookService {
	return &bookService{
		bookRepo: bookRepo,
	}
}

func (s *bookService) InsertBook(b dto.BookCreateDTO) entity.Book {
	book := entity.Book{}

	err := smapping.FillStruct(&book, smapping.MapFields(&b))

	if err != nil {
		log.Fatalf("failed to map %v", err)
	}

	res := s.bookRepo.InsertBook(book)

	return res
}

func (s *bookService) UpdateBook(b dto.BookUpdateDTO) entity.Book {
	book := entity.Book{}

	err := smapping.FillStruct(&book, smapping.MapFields(&b))

	if err != nil {
		log.Fatalf("failed to map %v", err)
	}

	res := s.bookRepo.UpdateBook(book)

	return res
}

func (s *bookService) DeleteBook(b entity.Book) {
	s.bookRepo.DeleteBook(b)
}

func (s *bookService) AllBooks() []entity.Book {
	return s.bookRepo.AllBooks()
}

func (s *bookService) FindById(bookID uint64) entity.Book {
	return s.bookRepo.FindById(bookID)
}

func (s *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
	b := s.bookRepo.FindById(bookID)

	id := fmt.Sprintf("%v", b.User.ID)

	return id == userID
}