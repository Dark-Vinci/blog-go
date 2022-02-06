package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"new-proj/dto"
	entity "new-proj/entities"
	"new-proj/helper"
	"new-proj/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type BookController interface {
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	All(ctx *gin.Context)
	FindById(ctx *gin.Context)
}

type bookController struct {
	bookService service.BookService
	jwtSerice service.JWTService
}

func NewBookController(bookService service.BookService, jwtSerice service.JWTService) BookController {
	return &bookController{
		jwtSerice: jwtSerice,
		bookService: bookService,
	}
}

func (c bookController) Insert(ctx *gin.Context) {
	var bookCreateDTO dto.BookCreateDTO

	errDTO := ctx.ShouldBind(&bookCreateDTO)

	if errDTO != nil {
		response := helper.BuildErrorResponse("failed to precess request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	userID := c.getUserIDByToken(authHeader)

	convertedUserId, err := strconv.ParseUint(userID, 10, 64)

	if err == nil {
		bookCreateDTO.UserID = uint16(convertedUserId)
	}

	// bookCreateDTO.UserID = uint16(convertedUserId)
	result := c.bookService.InsertBook(bookCreateDTO)

	response := helper.BuildResponse(true, "OK!", result)

	ctx.JSON(http.StatusCreated, response)
}

func (c bookController) Update(ctx *gin.Context) {
	var bookUpdateDTO dto.BookUpdateDTO
	errDTO := ctx.ShouldBind(&bookUpdateDTO)

	if errDTO != nil {
		response := helper.BuildErrorResponse("failed to precess request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")

	token, errToken := c.jwtSerice.ValidateToken(authHeader)

	if errToken != nil {
		response := helper.BuildErrorResponse("failed to precess request", errToken.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	if !c.bookService.IsAllowedToEdit(userID, uint64(bookUpdateDTO.ID)) {
		response := helper.BuildErrorResponse("you dont have the permission", "youre not the owner", helper.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	id, errID := strconv.ParseUint(userID, 10, 64)

	if errID == nil {
		bookUpdateDTO.UserID = uint16(id)
	}

	result := c.bookService.UpdateBook(bookUpdateDTO)
	response := helper.BuildResponse(true, "OK!", result)

	ctx.JSON(http.StatusOK, response)
}

func (c bookController) Delete(ctx *gin.Context) {
	var book entity.Book
	idParam := ctx.Param("id")

	id, err := strconv.ParseUint(idParam, 0, 0)

	if err != nil {
		response := helper.BuildErrorResponse("failed to get book id", "no param were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	book.ID = int64(id)

	authHeader := ctx.GetHeader("Authorization")

	token, errToken := c.jwtSerice.ValidateToken(authHeader)

	if errToken != nil {
		response := helper.BuildErrorResponse("failed to precess request", errToken.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	if !c.bookService.IsAllowedToEdit(userID, uint64(book.ID)) {
		response := helper.BuildErrorResponse("you dont have the permission", "youre not the owner", helper.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	c.bookService.DeleteBook(book)

	response := helper.BuildResponse(true, "DELETED!", book)
	ctx.JSON(http.StatusOK, response)
}

func (c bookController) All(ctx *gin.Context) {
	var books [] entity.Book = c.bookService.AllBooks()

	response := helper.BuildResponse(true, "OK!", books)

	ctx.JSON(http.StatusOK, response)
}

func (c bookController) FindById(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.ParseUint(paramId, 10, 16)

	if err != nil {
		response := helper.BuildErrorResponse("no id param was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	var book entity.Book = c.bookService.FindById(id)

	if (book == entity.Book{}) {
		response := helper.BuildErrorResponse("data not found", "no data with the given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.BuildResponse(true, "Ok!", book)
	ctx.JSON(http.StatusOK, response)
}

func (c bookController) getUserIDByToken(token string) string {
	aToken, err := c.jwtSerice.ValidateToken(token)

	if err != nil {
		panic(err.Error())
	}

	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])

	return id
}