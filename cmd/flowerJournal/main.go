package main

// 7314449266:AAFl490U2NEHvLuWopDJ4SxLD1cxek6asiU

import (
	"log"
	"sync"

	"elerphore.com/flower-journal/internal/mongo"
	"elerphore.com/flower-journal/internal/telegram"
	"elerphore.com/flower-journal/internal/web"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go telegram.InitTelegrammBot(&wg)

	wg.Add(2)
	go mongo.MongoDBConncetion(&wg)

	wg.Add(3)
	go web.WebServer(&wg)

	wg.Wait()
}
