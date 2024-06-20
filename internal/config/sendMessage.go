package telegram_config

import (
	"os"

	"elerphore.com/flower-journal/internal/mongo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SendMessageConfig struct {
	ChatID      int64
	Text        string
	ReplyMarkup tgbotapi.InlineKeyboardMarkup
}

func (config SendMessageConfig) Params() (tgbotapi.Params, error) {
	params := make(tgbotapi.Params)
	params.AddFirstValid("chat_id", config.ChatID)
	params.AddNonEmpty("text", config.Text)
	params.AddFirstValid("reply_markup", config.ReplyMarkup)
	return params, nil
}

func (config SendMessageConfig) Method() string {
	return "sendMessage"
}

func NewSendMessage(telegram_user mongo.User) SendMessageConfig {
	return SendMessageConfig{
		ChatID: telegram_user.Telegram_Chat_ID,
		Text:   os.Getenv("QUESTION_TO_ASK"),
		ReplyMarkup: tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Да", "1"),
				tgbotapi.NewInlineKeyboardButtonData("Нет", "2"),
			),
		),
	}
}
