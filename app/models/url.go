package models

import (
	"log"
	"math/rand"
	"time"
	"url_shortener/app/helpers"
)

type Url struct {
	ID        		int    `json:"id"`
	DestinationUrl 	string `json:"destination_url"`
	ShortUrl		string `json:"short_url"`
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func generateShortUrl(length int) string {
	return stringWithCharset(length, charset)
}

func CreateNewShortUrl(url *Url){
	count := 1
	for count > 0 {
		url.ShortUrl = generateShortUrl(7)
		count = helpers.ExecCount("SELECT COUNT(*) FROM "+helpers.UrlTable+" WHERE short_url = ?;", url.ShortUrl)
	}
	url.ID = int(helpers.ExecInsert("INSERT INTO "+helpers.UrlTable+" (short_url, destination_url) VALUES (?, ?)",url.ShortUrl,url.DestinationUrl))
}

func GetByShortUrl(url *Url){
	rows := helpers.ExecSelect("SELECT id, destination_url FROM "+helpers.UrlTable+" WHERE short_url = ?;", url.ShortUrl)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&url.ID, &url.DestinationUrl)
		if err != nil {
			log.Println(err)
		}
		break
	}
}

func DeleteById(url *Url) bool{
	count := helpers.ExecDelete("DELETE FROM "+helpers.UrlTable+" WHERE id = ?;", url.ID)
	return count > 0
}