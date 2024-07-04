package telegram

import (
	"elerphore.com/flower-journal/internal/mongo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func onCommand(msg tgbotapi.MessageConfig, update tgbotapi.Update) {
	msg.Text = "Вы зарегестрированы в системе!"
	resp := SendMessageNative(msg)
	mongo.InsertNewMessage(*resp)
	mongo.InsertNewUser(update, *resp)
}

func onUnknown(msg tgbotapi.MessageConfig) {
	msg.Text = "Не известная команда"
	SendMessageNative(msg)
}
