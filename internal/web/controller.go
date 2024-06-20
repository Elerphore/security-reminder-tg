package web

import (
	"log"
	"net/http"
	"os"
	"sync"
)

func WebServer(wg *sync.WaitGroup) {
	defer wg.Done()

	http.HandleFunc("/api/executeScheduling", executeScheduling)

	http.HandleFunc("/api/setAllDaysToDefaultState", updateDays)

	if err := http.ListenAndServe(os.Getenv("PORT"), nil); err != nil {
		log.Default().Println(err)
	}
}
