package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/erwindosianipar/common/http/middleware"
	"github.com/erwindosianipar/common/http/request"
	"github.com/erwindosianipar/common/util/token"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	secretKey = "v3ryvery53cr3tk3y"
	authToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImVyd2luZG8iLCJyb2xlIjoiVVNFUiIsImV4cGlyZWQiOiIyMDIzLTAyLTIyVDIzOjM0OjE4KzA3OjAwIn0.70ls_wGpKiUdjzRdLbd-A80KH91Llsh4kwjyW-b26MY"
)

func TestCORS(t *testing.T) {
	t.Run("test normal case cors", func(t *testing.T) {
		gin := gin.New()
		rec := httptest.NewRecorder()
		h := request.DefaultHandler()

		gin.Use(middleware.NewMiddleware(secretKey).CORS("*"))
		gin.GET("/", h.Index)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		gin.ServeHTTP(rec, req)

		t.Run("test status code and access allow origin", func(t *testing.T) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "*", rec.Header().Get("Access-Control-Allow-Origin"))
		})
	})
}

func TestAUTH(t *testing.T) {
	t.Run("test normal case auth", func(t *testing.T) {
		gin := gin.New()
		rec := httptest.NewRecorder()
		h := request.DefaultHandler()

		gin.Use(middleware.NewMiddleware(secretKey).AUTH("USER"))
		gin.GET("/", h.Index)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", authToken)

		gin.ServeHTTP(rec, req)

		authorization := req.Header.Get("Authorization")

		t.Run("test status code and validate authorization token", func(t *testing.T) {
			if len(authorization) > 0 {
				token, err := token.NewToken(secretKey).ValidateToken(authorization, "USER")
				if err != nil {
					t.Fatal(err)
				}

				assert.Equal(t, true, token.Valid)
				assert.Equal(t, http.StatusOK, rec.Code)
			} else {
				t.Fatal("test failed")
			}
		})
	})
}
