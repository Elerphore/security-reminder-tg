package telegram

import (
	"log"

	"elerphore.com/flower-journal/internal/mongo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	switch update.Message.Command() {
	case "start":
		msg.Text = "Вы зарегестрированы в системе!"
		insertNewUser(update)
	default:
		msg.Text = "It's not going to work :)"
	}

	SendMessageNative(msg)
}

func callBackQueryHandle(update tgbotapi.Update) {

	DeleteMessages(update.FromChat().ID)

	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)

	if _, err := telegramm_bot.Request(callback); err != nil {
		log.Default().Println(err)
	}

	var resp string

	if update.CallbackQuery.Data == "1" {
		user := mongo.GetUserByTelegramUserId(update.CallbackQuery.From.ID)
		mongo.UpdateCurrentDay(user.ID)
		resp = "Умница!"
	}

	if update.CallbackQuery.Data == "2" {
		resp = "Если бы я имплементировал сюда андектоты, то ты бы получила один, 100%!"
	}

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, resp)

	SendMessageNative(msg)
}

func insertNewUser(update tgbotapi.Update) {

	var user = mongo.User{
		ID:               primitive.NewObjectID(),
		Telegram_User_ID: update.SentFrom().ID,
		Telegram_Chat_ID: update.FromChat().ID,
		Send_Time:        "18:30",
	}

	mongo.InsertNewUser(user)
}
