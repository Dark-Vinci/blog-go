package service

import (
	"fmt"
	"time"

	"new-proj/helper"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(userId string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secreteKey string
	issuer string
}

func NewJWTService () JWTService {
	toReturn := &jwtService {
		issuer: "hdnjcej",
		secreteKey: helper.GetSecretKey(),
	}

	return toReturn
}

func (j *jwtService) GenerateToken(UserID string) string {
	claims := &jwtCustomClaim {
		UserID,
		jwt.StandardClaims {
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer: j.issuer,
			IssuedAt: time.Now().Unix(),
		},
	}

	// SigningMethodES256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([] byte(j.secreteKey))

	if err != nil {
		panic(err)
	}
	
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func (t_ *jwt.Token) (interface {}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}

		return []byte(j.secreteKey), nil
	})
}