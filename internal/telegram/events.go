package telegram

import (
	"log"

	"elerphore.com/flower-journal/internal/http_client"
	"elerphore.com/flower-journal/internal/mongo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func listenToUpdates() {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := telegramm_bot.GetUpdatesChan(u)

	handle(updates)
}

func handle(updates tgbotapi.UpdatesChannel) {
	for update := range updates {

		if update.Message != nil && update.Message.IsCommand() {
			commandHandle(update)
		}

		if update.CallbackQuery != nil {
			callBackQueryHandle(update)
		}
	}
}

func commandHandle(update tgbotapi.Update) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

	command := update.Message.Command()

	if command == "start" {
		onCommand(msg, update)
	}

	if command != "start" {
		onUnknown(msg)
	}

}

func callBackQueryHandle(update tgbotapi.Update) {
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)

	if _, err := telegramm_bot.Request(callback); err != nil {
		log.Default().Println(err)
	}

	var resp string

	user := mongo.GetUserByTelegramUserId(update.CallbackQuery.From.ID)

	if update.CallbackQuery.Data == "1" {

		mongo.UpdateCurrentDay(user.ID)
		resp = "Умница!"
	}

	if update.CallbackQuery.Data == "2" {
		resp = "Анекдот: " + http_client.GetJoke().Joke
	}

	msg := tgbotapi.NewEditMessageText(user.Telegram_Chat_ID, user.Telegram_Message_ID, resp)

	telegramm_bot.Send(msg)
}
