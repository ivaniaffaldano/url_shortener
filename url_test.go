package main

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
	"url_shortener/app/controllers"
	"url_shortener/app/helpers"
	"url_shortener/app/models"
)

var url models.Url = models.Url{}

func TestCreateUrl(t *testing.T) {
	helpers.CreateDB()
	url := getUrl()
	destinationUrl := "https://www.google.it/?test=" + strconv.Itoa(getRandomNumber())
	url.DestinationUrl = destinationUrl
	body,_ := json.Marshal(url)
	req, err := http.NewRequest("POST", "/api/create", strings.NewReader(string(body)))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.CreateUrl)
	handler.ServeHTTP(rr, req)
	assert.Equal(t,rr.Code,http.StatusOK,"Error HTTP Request")
	err = json.NewDecoder(rr.Body).Decode(&url)
	assert.Equal(t, url.DestinationUrl, destinationUrl, "The two url should be the same. ID: " + strconv.Itoa(url.ID))
	assert.NotEqual(t, url.ID, 0, "ID not created")
}

func TestGetUrl(t *testing.T) {
	url := *getUrl()
	id := url.ID
	req, err := http.NewRequest("POST", "/api/get/" + url.ShortUrl, strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetUrl)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("shortUrl", url.ShortUrl)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	handler.ServeHTTP(rr, req)
	assert.Equal(t,rr.Code,http.StatusOK,"Error HTTP Request")
	err = json.NewDecoder(rr.Body).Decode(&url)
	assert.Equal(t, url.ID, id, "The two id should be the same. ID: " + strconv.Itoa(url.ID))
	assert.NotEqual(t, url.ID, 0, "ID not found. ID: " + strconv.Itoa(url.ID))
}

func TestDeleteUrl(t *testing.T) {
	url := *getUrl()
	req, err := http.NewRequest("DELETE", "/api/delete/" + strconv.Itoa(url.ID), strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.DeleteUrl)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", strconv.Itoa(url.ID))
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	handler.ServeHTTP(rr, req)
	assert.Equal(t,rr.Code,http.StatusOK,"Error HTTP Request")
	deleted := helpers.DeletedResponse{}
	err = json.NewDecoder(rr.Body).Decode(&deleted)
	assert.Equal(t, deleted.Deleted, true, "Not Deleted.")
}

func TestGetDeletedUrl(t *testing.T) {
	url := *getUrl()
	id := url.ID
	req, err := http.NewRequest("POST", "/api/get/" + url.ShortUrl, strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetUrl)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("shortUrl", url.ShortUrl)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	handler.ServeHTTP(rr, req)
	assert.Equal(t,rr.Code,http.StatusOK,"Error HTTP Request")
	err = json.NewDecoder(rr.Body).Decode(&url)
	assert.Equal(t, url.ID, 0, "The two id should be the same. ID: " + strconv.Itoa(url.ID))
	assert.NotEqual(t, url.ID, id, "ID not found. ID: " + strconv.Itoa(url.ID))
}

func getRandomNumber() int{
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 1000000
	return int(rand.Intn(max - min + 1) + min)
}

func getUrl() *models.Url{
	return &url
}
