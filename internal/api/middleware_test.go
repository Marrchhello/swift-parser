package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLogger(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(Logger())
	router.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestErrorHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(ErrorHandler())

	// Route that panics
	router.GET("/panic", func(c *gin.Context) {
		panic("test panic")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/panic", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}
}

func TestValidateSwiftCode(t *testing.T) {
	tests := []struct {
		name       string
		swiftCode  string
		wantStatus int
	}{
		{
			name:       "Valid SWIFT code",
			swiftCode:  "TESTTR00XXX",
			wantStatus: http.StatusOK,
		},
		{
			name:       "Invalid length",
			swiftCode:  "SHORT",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Empty SWIFT code",
			swiftCode:  "",
			wantStatus: http.StatusBadRequest,
		},
	}

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(ValidateSwiftCode())
	router.GET("/:swiftCode", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/"+tt.swiftCode, nil)
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("ValidateSwiftCode() status = %d, want %d", w.Code, tt.wantStatus)
			}
		})
	}
}
