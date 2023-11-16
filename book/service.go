package book

import (
	"time"

	"github.com/google/uuid"
)

type service struct {
	repository Repository
}

type Service interface {
	GetAll() (*[]Entity, error)
	Get(bookId string) (*Entity, error)
	Create(bookInput BookRequest) (*Entity, error)
	Update(bookId string, bookInput BookRequest) (*Entity, error)
	Delete(bookId string) error
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetAll() (*[]Entity, error) {
	books, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *service) Get(bookId string) (*Entity, error) {
	book, err := s.repository.FindById(bookId)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *service) Create(bookInput BookRequest) (*Entity, error) {
	newBook := Entity{
		ID:          uuid.NewString(),
		Title:       bookInput.Title,
		Description: bookInput.Description,
		Genre:       bookInput.Genre,
		Author:      bookInput.Author,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	book, err := s.repository.Create(newBook)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *service) Update(bookId string, bookInput BookRequest) (*Entity, error) {
	book, err := s.repository.FindById(bookId)
	if err != nil {
		return nil, err
	}

	book.Title = bookInput.Title
	book.Description = bookInput.Description
	book.Genre = bookInput.Genre
	book.Author = bookInput.Author
	book.UpdatedAt = time.Now()

	updatedBook, err := s.repository.Update(*book)
	if err != nil {
		return nil, err
	}

	return updatedBook, nil
}

func (s *service) Delete(bookId string) error {
	_, err := s.repository.FindById(bookId)
	if err != nil {
		return err
	}

	err = s.repository.Delete(bookId)
	if err != nil {
		return err
	}

	return nil
}
