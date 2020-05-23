package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"url_shortener/app/helpers"
	"url_shortener/app/models"
)

// GetUrl godoc
// @Summary Get Destination URL from Short Url
// @Description get Destination URL from Short Url, params example: http://localhost:8080/api/get/2SHcWFg
// @Param short_url path string true "short URL"
// @Success 200 {object} models.Url
// @Router /api/get/{shortUrl} [get]
func GetUrl(w http.ResponseWriter, r *http.Request) {
	url 			:= models.Url{}
	url.ShortUrl 	= chi.URLParam(r, "shortUrl")
	models.GetByShortUrl(&url)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(url)
}

// DeleteUrl godoc
// @Summary Delete Url
// @Description delete URL from Id, params example: http://localhost:8080/api/delete/2SHcWFg
// @Param id path int true "1"
// @Success 200 {object} helpers.DeletedResponse
// @Router /api/delete [delete]
func DeleteUrl(w http.ResponseWriter, r *http.Request) {
	url 			:= models.Url{}
	url.ID,_ 		= strconv.Atoi(chi.URLParam(r, "id"))
	w.Header().Set("Content-Type", "application/json")
	resp := helpers.DeletedResponse{models.DeleteById(&url)}
	json.NewEncoder(w).Encode(resp)
}

// CreateUrl godoc
// @Summary Create Short URL
// @Description get short URL from dest URL, params example: {"destination_url":"example"}
// @Param destination_url body string true "destination URL"
// @Success 200 {object} models.Url
// @Failure 200 {object} helpers.ErrorRequest
// @Router /api/create [post]
func CreateUrl(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	var url models.Url
	err = json.NewDecoder(r.Body).Decode(&url)
	if err != nil || url.DestinationUrl == ""{
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(helpers.ErrorRequest{Error: "Bad Request"})
	}

	models.CreateNewShortUrl(&url)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(url)

}

// RedirectShortUrl godoc
// @Summary Redirect Url to DestinationUrl
// @Description Redirect to Destination URL from Short Url, params example: http://localhost:8080/2SHcWFg
// @Param short_url path string true "short URL"
// @Success 301
// @Failure 200 {object} helpers.ErrorRequest
// @Router /{shortUrl} [get]
func RedirectShortUrl(w http.ResponseWriter, r *http.Request){
	url 			:= models.Url{}
	url.ShortUrl 	= chi.URLParam(r, "shortUrl")
	models.GetByShortUrl(&url)
	if url.DestinationUrl != "" {
		http.Redirect(w, r, url.DestinationUrl, 301)
	}else{
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(helpers.ErrorRequest{Error: "Not Found"})
	}
}