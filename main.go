package main

import (
	"backend/server"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server.NewServer()
	// res, err := http.Get("https://backend-production-f609.up.railway.app/guru/all")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// body, _ := io.ReadAll(res.Body)
	// fmt.Println(body)
}
