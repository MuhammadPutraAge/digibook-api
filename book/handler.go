package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/muhammadputraage/digibook-api/utils"
)

type handler struct {
	service Service
}

func InitRouter(r *gin.RouterGroup, service Service) {
	bookHandler := &handler{service}

	r.GET("/books", bookHandler.GetBooks)
	r.GET("/books/:id", bookHandler.GetBook)
	r.POST("/books", bookHandler.CreateBook)
	r.PUT("/books/:id", bookHandler.UpdateBook)
	r.DELETE("/books/:id", bookHandler.DeleteBook)
}

func (h *handler) GetBooks(c *gin.Context) {
	books, err := h.service.GetAll()
	if err != nil {
		response := utils.HandleErrorResponse("failed to get all books", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.HandleSuccessResponse("success get all books", gin.H{"books": books})
	c.JSON(http.StatusOK, response)
}

func (h *handler) GetBook(c *gin.Context) {
	book, err := h.service.Get(c.Param("id"))
	if err != nil {
		response := utils.HandleErrorResponse("failed to get book", err.Error())
		c.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	response := utils.HandleSuccessResponse("success get book", gin.H{"book": book})
	c.JSON(http.StatusOK, response)
}

func (h *handler) CreateBook(c *gin.Context) {
	var bookInput BookRequest

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		response := utils.HandleErrorResponse("failed to create book", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(bookInput)
	if err != nil {
		errMessage := err.(validator.ValidationErrors)[0].Error()

		response := utils.HandleErrorResponse("failed to create book", errMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	newBook, err := h.service.Create(bookInput)
	if err != nil {
		response := utils.HandleErrorResponse("failed to create book", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.HandleSuccessResponse("success create book", gin.H{"book": newBook})
	c.JSON(http.StatusOK, response)
}

func (h *handler) UpdateBook(c *gin.Context) {
	var bookInput BookRequest

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		response := utils.HandleErrorResponse("failed to update book", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(bookInput)
	if err != nil {
		errMessage := err.(validator.ValidationErrors)[0].Error()

		response := utils.HandleErrorResponse("failed to update book", errMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	updatedBook, err := h.service.Update(c.Param("id"), bookInput)
	if err != nil {
		response := utils.HandleErrorResponse("failed to update book", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.HandleSuccessResponse("success update book", gin.H{"book": updatedBook})
	c.JSON(http.StatusOK, response)
}

func (h *handler) DeleteBook(c *gin.Context) {
	err := h.service.Delete(c.Param("id"))
	if err != nil {
		response := utils.HandleErrorResponse("failed to delete book", err.Error())
		c.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	response := utils.HandleSuccessResponse("success delete book", gin.H{})
	c.JSON(http.StatusOK, response)
}
