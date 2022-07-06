package server

import (
	"time"

	"github.com/upgradeskill/bookstore-oktaatr/domain"
	_bookService "github.com/upgradeskill/bookstore-oktaatr/modules/book/service"
	"github.com/upgradeskill/bookstore-oktaatr/utils"
)

type Service struct {
	BookService domain.BookService
	AuthService domain.AuthService
}

func NewService(conn *utils.Conn, r *Repository, timeoutContext time.Duration) *Service {
	return &Service{
		// AuthService:    _authUcase.NewAuthService(cfg, r.UserRepo, r.UnitRepo, r.RoleRepo, r.RolePermRepo, timeoutContext),
		BookService: _bookService.NewBookService(r.BookRepo, timeoutContext),
	}
}
