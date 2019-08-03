package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/thedevsaddam/renderer"
)

// ErrorResponseObject can be used to send back error responses
type ErrorResponseObject struct {
	Status int    `json:status`
	Error  string `json:error`
}

// RequestLogger hijacks a request and logs it for viewing
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(`%s %s`, r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

// Encrypt plaintext with a 32 byte key using AES in GCM mode.
func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// Decrypt cipher text to plain text with a 32 byte key using AES in GCM mode.
func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

//
func RequestResponder(w http.ResponseWriter, req *http.Request, statusCode int, data interface{}) {
	// create a renderer object
	rnd := renderer.New()

	switch req.Header.Get("Accept") {
	case renderer.ContentJSON:
		rnd.JSON(w, statusCode, data)
	case renderer.ContentXML:
		rnd.XML(w, statusCode, data)
	case renderer.ContentYAML:
		rnd.YAML(w, statusCode, data)
	default:
		rnd.JSON(w, statusCode, data)
	}
}
