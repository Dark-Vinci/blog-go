package controller

import (
	"net/http"
	"strconv"

	"new-proj/dto"
	"new-proj/helper"
	"new-proj/service"
	entity "new-proj/entities"

	"github.com/gin-gonic/gin"
)


type AuthController interface {
	Login (ctx *gin.Context)
	Register (ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		jwtService: jwtService,
		authService: authService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)

	if errDTO != nil {
		response := helper.BuildErrorResponse("failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)

	if v, ok := authResult.(entity.User); ok {
		generateToken := c.jwtService.GenerateToken(strconv.FormatInt(v.ID, 10))
		v.Token = generateToken

		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.BuildErrorResponse("plase check your credential", "invalid credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO

	errorDTO := ctx.ShouldBind(&registerDTO)

	if errorDTO != nil {
		response := helper.BuildErrorResponse("failed to process request", errorDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("failed to process respinse", "duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(uint64(createdUser.ID), 10))
		
		createdUser.Token = token

		response := helper.BuildResponse(true, "ok!", createdUser)

		ctx.JSON(http.StatusOK, response)
	}
}