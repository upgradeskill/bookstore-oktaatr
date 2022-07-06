package main

import (
	"time"

	"github.com/upgradeskill/bookstore-oktaatr/config"
	"github.com/upgradeskill/bookstore-oktaatr/server"
)

func main() {
	conn := config.NewDBConn()

	timeoutContext := time.Duration(10) * time.Second

	// init repo category repo
	repository := server.NewRepository(conn)
	service := server.NewService(conn, repository, timeoutContext)
	server.NewHandler(service)
}
