package telegram_config

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type DeleteMessageConfig struct {
	ChannelUsername string
	ChatID          int64
	MessageID       int
}

func (config DeleteMessageConfig) Method() string {
	return "deleteMessage"
}

func (config DeleteMessageConfig) Params() (tgbotapi.Params, error) {
	params := make(tgbotapi.Params)

	params.AddFirstValid("chat_id", config.ChatID, config.ChannelUsername)
	params.AddNonZero("message_id", config.MessageID)

	return params, nil
}

func NewDeleteMessage() DeleteMessageConfig {
	return DeleteMessageConfig{
		ChatID:    735391827,
		MessageID: 735391827,
	}
}
