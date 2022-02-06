package middleware

import (
	"log"
	"net/http"
	
	"new-proj/helper"
	"new-proj/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func (c *gin.Context) {
		authHeader := c.GetHeader("x-auth-token")

		if authHeader == "" {
			response := helper.BuildErrorResponse("failed to process request", "no token found", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}

		token, err := jwtService.ValidateToken(authHeader)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)

			log.Println("user id", claims["user_id"])
			log.Println("user id", claims["issuer"])
		} else {
			log.Println(err);

			response := helper.BuildErrorResponse("invalid token", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
	}
}