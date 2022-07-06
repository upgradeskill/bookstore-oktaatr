package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_bookHandler "github.com/upgradeskill/bookstore-oktaatr/modules/book/handler"
)

func newAppHandler(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"app": "Book Apps",
		})
	})
}

// NewHandler will create a new handler for the given usecase
func NewHandler(svc *Service) {

	e := echo.New()
	e.HTTPErrorHandler = ErrorHandler

	route := e.Group("")

	newAppHandler(e)
	_bookHandler.NewBookHandler(route, svc.BookService)

	log.Fatal(e.Start(":3000"))
}

// ErrorHandler ...
func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Logger().Error(report)
	c.JSON(report.Code, report)
}
