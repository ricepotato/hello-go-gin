package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ricepotato/hello-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func TestPing(t *testing.T) {
	router := gin.Default()
	router.GET("/ping", controllers.Ping)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Error("Status code should be 200, but got", w.Code)
	}
	if w.Body.String() != "{\"message\":\"pong\"}" {
		t.Error("Response body should be {\"message\":\"pong\"}, but got", w.Body.String())
	}

}
