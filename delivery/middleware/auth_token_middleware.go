package middleware

import (
	"fmt"
	"livecode-wmb-2/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

type AuthTokenMiddleware interface {
	RequireToken() gin.HandlerFunc
}

type authTokenMiddleware struct {
	acctToken utils.Token
}

// RequireToken implements AuthTokenMiddleware
func (a *authTokenMiddleware) RequireToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"massage": "unauthorized",
			})
			c.Abort()
		}

		tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)
		fmt.Println("token :", tokenString)

		if h.AuthorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"massage": "token invalid",
			})
			c.Abort()
			return
		}

		token, err := a.acctToken.VerrifyAccesToken(tokenString)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{
				"massage": "unauthorized",
			})
			c.Abort()
			return
		}
		fmt.Println("token :", token)
		if token != nil {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"massage": "unauthorized",
			})
			c.Abort()
			return
		}
	}
}

func NewAuthTokenValidator(acctToken utils.Token) AuthTokenMiddleware {
	return &authTokenMiddleware{acctToken: acctToken}
}
