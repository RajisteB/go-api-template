package middleware

import (
	"net/http"
	"server/utils"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			utils.WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": "auth token required"})
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := utils.VerifyJWT(tokenString)
		if err != nil {
			utils.WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": "invalid token"})
			return
		}
		name := claims.(jwt.MapClaims)["name"].(string)
		role := claims.(jwt.MapClaims)["role"].(string)

		r.Header.Set("name", name)
		r.Header.Set("role", role)

		next.ServeHTTP(w, r)
	})
}
