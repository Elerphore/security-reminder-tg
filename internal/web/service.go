package web

import (
	"log"
	"net/http"

	"elerphore.com/flower-journal/internal/mongo"
	"elerphore.com/flower-journal/internal/telegram"
	telegram_config "elerphore.com/flower-journal/internal/telegramm_config"
)

func executeScheduling(w http.ResponseWriter, r *http.Request) {
	users := mongo.GetUsers()

	for index, user := range users {

		day := mongo.FindCurrentDayByUserId(user.ID)

		if !day.Used {
			telegram.DeleteMessages(user.Telegram_Chat_ID)

			var request = telegram_config.NewSendMessage(users[index])
			telegram.SendMessage(request)
		} else {
			log.Default().Println("User:", user.ID, "already did it")
		}
	}
}

func updateDays(w http.ResponseWriter, r *http.Request) {
	mongo.SetAllDaysToDefaultState()
}
