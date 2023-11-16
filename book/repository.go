package book

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	FindAll() (*[]Entity, error)
	FindById(bookId string) (*Entity, error)
	Create(book Entity) (*Entity, error)
	Update(book Entity) (*Entity, error)
	Delete(bookId string) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() (*[]Entity, error) {
	var bookList []Entity

	err := r.db.Find(&bookList).Error
	if err != nil {
		return nil, err
	}

	return &bookList, err
}

func (r *repository) FindById(bookId string) (*Entity, error) {
	var book Entity

	err := r.db.First(&book, bookId).Error
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *repository) Create(book Entity) (*Entity, error) {
	err := r.db.Create(&book).Error
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *repository) Update(book Entity) (*Entity, error) {
	err := r.db.Save(&book).Error
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *repository) Delete(bookId string) error {
	err := r.db.Delete(&Entity{}, bookId).Error
	if err != nil {
		return err
	}

	return nil
}
