package web

import (
	"log"
	"net/http"

	cofig "elerphore.com/flower-journal/internal/config"
	"elerphore.com/flower-journal/internal/mongo"
	"elerphore.com/flower-journal/internal/telegram"
)

func executeScheduling(w http.ResponseWriter, r *http.Request) {
	users := mongo.GetUsers()

	for index, user := range users {

		day := mongo.FindCurrentDayByUserId(user.ID)

		if !day.Used {
			telegram.DeleteMessages(user.Telegram_Chat_ID)

			var request = cofig.NewSendMessage(users[index])
			telegram.SendMessage(request)
		} else {
			log.Default().Println("User:", user.ID, "already did it")
		}
	}
}

func updateDays(w http.ResponseWriter, r *http.Request) {
	mongo.SetAllDaysToDefaultState()
}
