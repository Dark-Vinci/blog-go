package controller

import (
	"fmt"
	"net/http"
	"new-proj/dto"
	"new-proj/helper"
	"new-proj/service"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService: jwtService,
	}
}

func (c *userController) Update(ctx *gin.Context) {
	var userUpdateDTO  dto.UserUpdateDTO

	err := ctx.ShouldBind(&userUpdateDTO)

	if err != nil {
		response := helper.BuildErrorResponse("failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")

	token, errToken := c.jwtService.ValidateToken(authHeader)

	if errToken != nil {
		fmt.Println("panic")
		panic(errToken.Error())
	}

	fmt.Println(token)
	claims := token.Claims.(jwt.MapClaims)

	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	fmt.Println(id)

	if err != nil {
		fmt.Println("panic 2")
		panic(err.Error())
	}

	fmt.Println("hereeee")
	userUpdateDTO.ID = id
	// fmt.Println(userUpdateDTO)
	u := c.userService.Update(userUpdateDTO)
	// fmt.Println("we dey")
	response := helper.BuildResponse(true, "Ok!", u)

	ctx.JSON(http.StatusOK, response)
}

func (c *userController) Profile(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)

	if err != nil {
		panic(err.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])

	user := c.userService.Profile(id)

	response := helper.BuildResponse(true, "Ok!", user)

	ctx.JSON(http.StatusOK, response)
}