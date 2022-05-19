package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

const (
	HEADER_AUTH = "X-Poll-Token"
)

func NewAuth(key string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := strings.Split(r.Header.Get(HEADER_AUTH), "Bearer ")
			if len(authHeader) != 2 {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Malformed Token"))
			} else {
				jwtToken := authHeader[1]
				token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}
					return []byte(key), nil
				})
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					ctx := context.WithValue(r.Context(), "props", claims)
					// ctx can provide the pollId so we can validate that
					// a single vote has been submitted per poll per user token
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					fmt.Println(err)
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Unauthorized"))
				}
			}
		})
	}
}
