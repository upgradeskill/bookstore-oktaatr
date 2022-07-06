package repository

import (
	"context"

	"github.com/upgradeskill/bookstore-oktaatr/domain"
	"github.com/upgradeskill/bookstore-oktaatr/helpers"
	"github.com/upgradeskill/bookstore-oktaatr/modules/book"
	"gorm.io/gorm"
)

type bookRepository struct {
	Conn *gorm.DB
}

func NewBookRepository(Conn *gorm.DB) domain.BookRepository {
	return &bookRepository{Conn}
}

func (m *bookRepository) Get(ctx context.Context, params *domain.Request) (res []domain.Book, total int64, err error) {
	var totalRows int64

	resBook := []domain.Book{}
	bk := []book.Book{}

	query := m.Conn.Model(&bk).Debug()

	query.Scopes(helpers.PaginateQuery(params)).Find(&bk).Count(&totalRows)
	if query.Error != nil {
		return []domain.Book{}, 0, query.Error
	}

	for _, data := range bk {
		resBook = append(resBook, domain.Book{
			ID:          int64(data.ID),
			Code:        data.Code,
			Name:        data.Name,
			Description: data.Description,
		})
	}

	return resBook, totalRows, nil
}
