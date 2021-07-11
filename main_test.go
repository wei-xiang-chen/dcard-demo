package main

import (
	"bytes"
	"dcard/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string, body interface{}) *httptest.ResponseRecorder {
	var req *http.Request
	if body != nil {
		jsonByte, _ := json.Marshal(body)
		req, _ = http.NewRequest(method, path, bytes.NewReader(jsonByte))
	} else {
		req, _ = http.NewRequest(method, path, nil)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestWrapper(t *testing.T) {
	originalUrl := "https://www.google.com.tw"
	expireAt := time.Now().Add(8 * time.Hour)

	urlId := test_Url(t, &originalUrl, &expireAt)
	test_GetOriginal(t, urlId, &originalUrl)
}

func test_Url(t *testing.T, originalUrl *string, expireAt *time.Time) *string {
	router := initializeRoutes()

	reqBody := model.UrlInput{Url: originalUrl, ExpireAt: expireAt}

	w := performRequest(router, http.MethodPost, "/api/v1/urls/", reqBody)

	body, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Printf("response:%v\n", string(body))

	var respBody model.UrlOutput
	json.Unmarshal(body, &respBody)
	return &respBody.Id
}

func test_GetOriginal(t *testing.T, urlId *string, originalUrl *string) {
	router := initializeRoutes()

	w := performRequest(router, http.MethodGet, "/"+*urlId, nil)

	body, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, http.StatusMovedPermanently, w.Code)
	assert.Equal(t, true, strings.Contains(string(body), *originalUrl))

	fmt.Printf("response:%v\n", string(body))
}
