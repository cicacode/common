package middleware

import (
	"net/http"

	"github.com/cicacode/common/http/response"
	"github.com/cicacode/common/model"
	"github.com/cicacode/common/util/token"

	"github.com/gin-gonic/gin"
)

type Middleware interface {
	CORS(host string) gin.HandlerFunc
	AUTH(role string) gin.HandlerFunc
}

type middleware struct {
	secretKey string
}

func NewMiddleware(secretKey string) Middleware {
	return &middleware{secretKey: secretKey}
}

func (m *middleware) CORS(host string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", host)
		c.Next()
	}
}

func (m *middleware) AUTH(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		if len(authorization) > 1 {
			token, err := token.NewToken(m.secretKey).ValidateToken(authorization, role)
			if err != nil {
				response.New(c).Error(http.StatusInternalServerError, model.Lang{Id: "Internal server error" + err.Error(), En: "Terjadi kesalahan server"})
				c.Abort()
			} else if token.Valid {
				c.Next()
			} else {
				response.New(c).Error(http.StatusUnauthorized, model.Lang{Id: "Token otorisasi tidak valid", En: "Invalid authorization token"})
				c.Abort()
			}
		} else {
			response.New(c).Error(http.StatusUnauthorized, model.Lang{Id: "Diperlukan token otorisasi", En: "Authorization token required"})
			c.Abort()
		}
	}
}
