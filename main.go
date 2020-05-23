package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"url_shortener/app/helpers"
	"url_shortener/app/route"
)

// fun main()
func main() {

	// @title URL Shortener API
	// @version 1.0
	// @description This is a test.

	// @contact.name Ivan Iaffaldano
	// @contact.email i.iaffaldano@gmail.com

	helpers.CreateDB()

	// load env vars
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	appEnv := os.Getenv("app_env")
	log.Println(appEnv)

	//creating instance of mux
	routes := route.GetRoutes()
	http.Handle("/", routes)

	// just for test different application environments
	if appEnv == "local" {
		log.Fatal(http.ListenAndServe(":8080", routes))
	}else{
		log.Fatal(http.ListenAndServe(":8080", routes))
	}
}


