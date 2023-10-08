package main

import (
	"strings"

	"github.com/Dparty/common/utils"
	"github.com/gin-gonic/gin"
)

type Headers struct {
	Authorization string
}

type AuthenticationStatus string

const (
	Authorized   = "Authorized"
	Unauthorized = "Unauthorized"
)

type Authentication struct {
	Status    AuthenticationStatus
	AccountId uint
	Email     string
}

func Authorize(c *gin.Context) Authentication {
	var headers Headers
	c.ShouldBindHeader(&headers)
	authorization := headers.Authorization
	splited := strings.Split(authorization, " ")
	if authorization == "" || len(splited) != 2 {
		return Authentication{
			Status: Unauthorized,
		}
	}
	return AuthorizeByJWT(splited[1])
}

func AuthorizeByJWT(token string) Authentication {
	claims, err := utils.VerifyJwt(token)
	if err != nil {
		return Authentication{
			Status: Unauthorized,
		}
	}
	return Authentication{
		Status:    Authorized,
		Email:     claims["email"].(string),
		AccountId: utils.StringToUint(claims["id"].(string)),
	}
}
