package herramientas

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"
)

func GenerarToken(cant int) (string, error) {
	caracteres := make([]byte, cant)
	_, err := rand.Read(caracteres)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(caracteres), nil
}

func ValidarToken(cookie *http.Cookie) bool {
	if cookie.Value == "" || cookie.Expires.Before(time.Now()) {
		return false
	}
	return true
}

func Middleware(SigManejador http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "No autorizado", http.StatusUnauthorized) // no son correctas las credenciales
				return
			}
			http.Error(w, "Error interno del servidor", http.StatusInternalServerError) // fallo al obtener la cookie
			return
		}
		if !ValidarToken(cookie) {
			http.Error(w, "No autorizado", http.StatusUnauthorized)
			return
		}
		SigManejador.ServeHTTP(w, r)
	})
}
