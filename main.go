package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadputraage/digibook-api/book"
	"github.com/muhammadputraage/digibook-api/config"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
}

func main() {
	bookRepository := book.NewRepository(config.DB)
	bookService := book.NewService(bookRepository)

	r := gin.Default()

	apiV1 := r.Group("/api/v1")
	book.InitRouter(apiV1, bookService)

	r.Run()
}
