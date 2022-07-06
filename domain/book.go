package domain

import (
	"context"
)

type Book struct {
	ID          int64  `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"email"`
}

type BookRequest struct {
	Code        string `json:"code" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type BookResponse struct {
	ID          int64  `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"email"`
}

type BookService interface {
	Get(ctx context.Context, params *Request) ([]Book, int64, error)
}

type BookRepository interface {
	Get(ctx context.Context, params *Request) (new []Book, total int64, err error)
}
