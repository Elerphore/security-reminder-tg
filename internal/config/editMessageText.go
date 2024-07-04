package telegram_config

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type EditMessageTextConfig struct {
	ChatID      int64
	Text        string
	MessageID   int
	ReplyMarkup *tgbotapi.InlineKeyboardMarkup
}

func (config EditMessageTextConfig) Params() (tgbotapi.Params, error) {
	params := make(tgbotapi.Params)
	params.AddFirstValid("chat_id", config.ChatID)
	params.AddNonZero("message_id", config.MessageID)
	params.AddNonEmpty("text", config.Text)

	if config.ReplyMarkup != nil {
		params.AddFirstValid("reply_markup", config.ReplyMarkup)
	}

	return params, nil
}

func (config EditMessageTextConfig) Method() string {
	return "editMessageText"
}

func NewEditTextMessageReply(chat_id int64, message_id int, text string) EditMessageTextConfig {

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Да", "1"),
			tgbotapi.NewInlineKeyboardButtonData("Нет", "2"),
		),
	)

	return EditMessageTextConfig{
		ChatID:      chat_id,
		MessageID:   message_id,
		Text:        text,
		ReplyMarkup: &keyboard,
	}
}

func NewEditTextMessage(chat_id int64, message_id int, text string) EditMessageTextConfig {
	return EditMessageTextConfig{
		ChatID:      chat_id,
		MessageID:   message_id,
		Text:        text,
		ReplyMarkup: nil,
	}
}
