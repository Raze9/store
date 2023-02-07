package v1_test

import (
	"GOproject/GIT/mail/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserRegister(t *testing.T) {
	r := gin.Default()
	r.POST("/register", v1.UserRegister)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status OK but got %v", w.Code)
	}
}

func TestUserLogin(t *testing.T) {
	r := gin.Default()
	r.POST("/login", v1.UserLogin)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status OK but got %v", w.Code)
	}
}

func TestUserUpdate(t *testing.T) {
	r := gin.Default()
	r.POST("/update", v1.UserUpdate)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/update", nil)
	req.Header.Add("Authorization", "")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status OK but got %v", w.Code)
	}
}

func TestUploadAvatar(t *testing.T) {
	r := gin.Default()
	r.POST("/uploadAvatar", v1.UploadAvatar)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/uploadAvatar", nil)
	req.Header.Add("Authorization", "")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status OK but got %v", w.Code)
	}
}

func TestSendEmail(t *testing.T) {
	r := gin.Default()
	r.POST("/sendemail", v1.Sendemail)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sendemail", nil)
	req.Header.Add("Authorization", "")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status OK but got %v", w.Code)
	}
}
