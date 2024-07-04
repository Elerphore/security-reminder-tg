package web

import (
	"net/http"
	"os"

	"elerphore.com/flower-journal/internal/mongo"
	"elerphore.com/flower-journal/internal/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func executeScheduling(w http.ResponseWriter, r *http.Request) {
	users := mongo.GetUsers()

	for _, user := range users {

		day := mongo.FindCurrentDayByUserId(user.ID)

		if !day.Used {
			msg := tgbotapi.NewMessage(user.Telegram_Chat_ID, os.Getenv("QUESTION_TO_ASK"))

			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Да", "1"),
					tgbotapi.NewInlineKeyboardButtonData("Нет", "2"),
				))

			resp := telegram.SendMessageNative(msg)
			mongo.UpdateUserMessageID(user.ID, resp.MessageID)
		}
	}
}

func updateDays(w http.ResponseWriter, r *http.Request) {
	mongo.SetAllDaysToDefaultState()
}
