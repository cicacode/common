package response

import (
	"net/http"

	"github.com/cicacode/common/model"

	"github.com/gin-gonic/gin"
)

type Wrapper interface {
	Write(code int, message model.Lang)
	Error(code int, message model.Lang)
	Token(token string)
}

type wrapper struct {
	c *gin.Context
}

func New(c *gin.Context) Wrapper {
	return &wrapper{c: c}
}

func (w *wrapper) Write(code int, message model.Lang) {
	w.c.JSON(code, model.Response{Code: code, Message: message})
}

func (w *wrapper) Error(code int, message model.Lang) {
	w.c.JSON(code, model.Response{Code: code, Message: message})
}

func (w *wrapper) Token(token string) {
	w.c.JSON(http.StatusOK, model.Token{Token: token})
}
