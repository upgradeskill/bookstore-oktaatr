package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/bookstore-oktaatr/domain"
	"github.com/upgradeskill/bookstore-oktaatr/helpers"
	"github.com/upgradeskill/bookstore-oktaatr/modules/book"
)

type BookHandler struct {
	BookSvc domain.BookService
}

func NewBookHandler(e *echo.Group, bs domain.BookService) {
	handler := &BookHandler{
		BookSvc: bs,
	}
	e.GET("/books", handler.GetAll)
}

func (h *BookHandler) GetAll(c echo.Context) error {

	cx := c.Request().Context()

	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"code": c.QueryParam("code"),
		"name": c.QueryParam("name"),
	}

	listBook, total, err := h.BookSvc.Get(cx, &params)

	if err != nil {
		return err
	}

	res := helpers.Paginate(c, book.NewListBookResponse(listBook), total, params)

	return c.JSON(http.StatusOK, res)
}
