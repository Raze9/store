package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
	r := NewRouter()

	// Test the ping route
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/ping", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "success", w.Body.String())

	// Test the user register route
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/v1/user/register", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code) // or any other expected HTTP status code for this route

	// Test the user login route
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/v1/user/login", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code) // or any other expected HTTP status code for this route

	// Test the user update route (requires authentication)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/api/v1/user", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code) // expected, since JWT middleware is applied to this route

	// Test the avatar upload route (requires authentication)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/v1/avatar", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code) // expected, since JWT middleware is applied to this route

	// Test the send email route (requires authentication)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/v1/user/send_email", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
