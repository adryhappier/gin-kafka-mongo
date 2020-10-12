package main

import (
	gin "github.com/adryhappier/gin-kafka-mongo/src/routes"
	"github.com/joho/godotenv"
)

func main() {
	// load .env environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Start Server Gin
	var s gin.Routes
	s.StartGin()
}

/*
	Karena variable s memiliki tipe gin.Server milik server.go
	maka s dapat menjalan kan fungsi StartGin
*/
