package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// TokenAuthMiddleware Middleware Login
func TokenAuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, valid := tokenValid(r)
		if !valid {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error":{"status": true,"msg":"Tidak ada akses ke halaman ini","code":"401"}}`))
			return
		}

		ctx := context.WithValue(r.Context(), "id", uint64(claims["user_id"].(float64)))
		ctx = context.WithValue(ctx, "role", claims["role"].(string))

		handler.ServeHTTP(w, r.WithContext(ctx))
	}
}

// TokenForgotPassMiddleware Middleware Forgot Password
func TokenForgotPassMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, valid := tokenValid(r)
		if !valid {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error":{"status": true,"msg":"Tidak ada akses ke halaman ini","code":"401"}}`))
			return
		}

		ctx := context.WithValue(r.Context(), "email", claims["email"].(string))

		handler.ServeHTTP(w, r.WithContext(ctx))
	}
}

// TokenValid tes
func tokenValid(r *http.Request) (jwt.MapClaims, bool) {
	var (
		claims jwt.MapClaims
		ok     bool
	)

	token, err := verifyToken(r)
	if err != nil {
		return claims, false
	}

	if claims, ok = token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return claims, false
	}

	return claims, true
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
