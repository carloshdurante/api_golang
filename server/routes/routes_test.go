package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestShowAllBooksRoute(t *testing.T) {
	r := ConfigRoutes(gin.Default())
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/books/", nil)
	r.ServeHTTP(w, req)
	checkResponseCode(t, http.StatusOK, w.Code)
}

// func TestShowBooksRoute(t *testing.T) {
// 	router := ConfigRoutes(gin.Default())
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/api/v1/books/1", nil)
// 	router.ServeHTTP(w, req)
// 	checkResponseCode(t, http.StatusOK, w.Code)
// }
