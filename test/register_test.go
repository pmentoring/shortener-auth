package test

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"shortener-auth/internal/routing"
	"strings"
	"testing"
)

func TestRegister(t *testing.T) {
	router := gin.Default()

	routing.Register(router)

	w := httptest.NewRecorder()
	request := ShortenUrlRequest{
		URL:   "https://google.com/hui",
		Title: "sukaaa",
	}
	jsonBody, _ := json.Marshal(request)

	req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(string(jsonBody)))
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	fmt.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}
