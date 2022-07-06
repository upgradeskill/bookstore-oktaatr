package book

import "github.com/upgradeskill/bookstore-oktaatr/domain"

func NewListBookResponse(datas []domain.Book) []domain.Book {
	resp := []domain.Book{}

	for _, data := range datas {
		resp = append(resp, NewBookResponse(data))
	}

	return resp
}

func NewBookResponse(user domain.Book) domain.Book {
	return domain.Book{
		ID:          user.ID,
		Code:        user.Code,
		Name:        user.Name,
		Description: user.Description,
	}
}
