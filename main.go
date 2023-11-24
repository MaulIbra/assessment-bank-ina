package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MaulIbra/assessment-bank-ina/service"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file not FOUND")
		}
	}
	r := service.Init()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%v", os.Getenv("PORT")),
		Handler:        r,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
