package request

import (
	"net/http"

	"github.com/cicacode/common/http/response"
	"github.com/cicacode/common/model"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	NoRoute(c *gin.Context)
	Index(c *gin.Context)
}

type handler struct {
	// Stuff maybe needed for handler
}

func DefaultHandler() Handler {
	return &handler{}
}

func (h *handler) NoRoute(c *gin.Context) {
	response.New(c).Error(http.StatusNotFound, model.Lang{Id: "Rute tidak ditemukan", En: "Route not found"})
}

func (h *handler) Index(c *gin.Context) {
	response.New(c).Write(http.StatusOK, model.Lang{Id: "Authentication API", En: "Authentication API"})
}
