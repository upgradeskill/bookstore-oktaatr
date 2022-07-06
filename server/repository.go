package server

import (
	"github.com/upgradeskill/bookstore-oktaatr/domain"
	_bookRepo "github.com/upgradeskill/bookstore-oktaatr/modules/book/repository"
	"github.com/upgradeskill/bookstore-oktaatr/utils"
)

type Repository struct {
	BookRepo domain.BookRepository
}

// NewRepository will create an object that represent all repos interface
func NewRepository(conn *utils.Conn) *Repository {
	return &Repository{
		BookRepo: _bookRepo.NewBookRepository(conn.GORM),
	}
}
