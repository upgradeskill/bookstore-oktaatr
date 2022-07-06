package service

import (
	"context"
	"time"

	"github.com/upgradeskill/bookstore-oktaatr/domain"
)

type bookService struct {
	bookRepo       domain.BookRepository
	contextTimeout time.Duration
}

func NewBookService(br domain.BookRepository, timeout time.Duration) domain.BookService {
	return &bookService{
		bookRepo:       br,
		contextTimeout: timeout,
	}
}

func (bs *bookService) Get(c context.Context, params *domain.Request) (res []domain.Book, total int64, err error) {

	ctx, cancel := context.WithTimeout(c, bs.contextTimeout)
	defer cancel()

	res, total, err = bs.bookRepo.Get(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	return
}
