package twlib

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

// GenerateNonce generates a random nonce of the given number of bytes and returns it as a hex encoded string.
func GenerateNonce(bytes int) (string, error) {
	if bytes < 1 {
		return "", fmt.Errorf("bytes must be greater than 0")
	}

	randomBytes := make([]byte, bytes)

	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(randomBytes), nil
}

// WithNonce returns a middleware for a http.Handler that generates random nonces of the given number of bytes and sets
// the Content-Security-Policy response header in addition to setting it in the context for automatic usage by templ.
func WithNonce(bytes int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			nonce, err := GenerateNonce(bytes)
			if err != nil {
				log.Printf("nonce could not be generated: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
			}

			w.Header().Add("Content-Security-Policy", fmt.Sprintf("script-src 'nonce-%s'", nonce))
			next.ServeHTTP(w, r.WithContext(templ.WithNonce(r.Context(), nonce)))
		})
	}
}
