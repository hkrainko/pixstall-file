package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type APIClaims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

type JWTPayloadsExtractor struct {
	Key       string
	SecretKey []byte
}

func NewJWTPayloadsExtractor(key string) *JWTPayloadsExtractor {
	return &JWTPayloadsExtractor{
		Key:      key,
		SecretKey: []byte("nBcWcVKTRiiUT0iaahFBFskAlugkP5GX"),
	}
}

func (j *JWTPayloadsExtractor) ExtractPayloadsFromJWTInHeader(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")
	if bearerToken == "" {
		c.Next()
		return
	}
	jwtToken := strings.TrimPrefix(bearerToken, "Bearer ")
	token, err := jwt.ParseWithClaims(
		jwtToken,
		&APIClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return j.SecretKey, nil
		},
	)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(*APIClaims); ok && token.Valid {
		c.Set(j.Key, claims.UserID)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
