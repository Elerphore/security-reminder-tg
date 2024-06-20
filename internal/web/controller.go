package web

import (
	"log"
	"net/http"
	"sync"
)

func WebServer(wg *sync.WaitGroup) {
	defer wg.Done()

	http.HandleFunc("/api/executeScheduling", executeScheduling)

	http.HandleFunc("/api/setAllDaysToDefaultState", updateDays)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Default().Println(err)
	}
}
