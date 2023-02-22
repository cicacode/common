package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

type Token interface {
	GenerateToken(username string, role string) string
	ValidateToken(token string, role string) (*jwt.Token, error)
}

type token struct {
	secretKey string
}

func NewToken(secretKey string) Token {
	return &token{secretKey: secretKey}
}

type authClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func (t *token) GenerateToken(username string, role string) string {
	claims := &authClaims{
		username,
		role,
		jwt.StandardClaims{
			Issuer:    "authentication",
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
		},
	}

	ctx := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := ctx.SignedString([]byte(t.secretKey))
	if err != nil {
		logrus.Panic(err)
	}

	return token
}

func (t *token) ValidateToken(encodedToken string, role string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, valid := token.Method.(*jwt.SigningMethodHMAC); !valid {
			return nil, fmt.Errorf("invalid token %v", token.Header["alg"])
		}
		if token.Claims.(jwt.MapClaims)["role"].(string) != role {
			return nil, fmt.Errorf("invalid role")
		}

		return []byte(t.secretKey), nil
	})
}
