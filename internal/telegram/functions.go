package telegram

import (
	"encoding/json"
	"log"

	"elerphore.com/flower-journal/internal/mongo"
	telegram_config "elerphore.com/flower-journal/internal/telegramm_config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func SendMessage(request telegram_config.SendMessageConfig) {

	params, _ := request.Params()
	response, err := telegramm_bot.MakeRequest(request.Method(), params)

	if err != nil {
		log.Default().Println(err)
	}

	response_bytes, err := response.Result.MarshalJSON()

	if err != nil {
		log.Default().Println(err)
	}

	var message mongo.MessageDTO
	json.Unmarshal(response_bytes, &message)

	mongo.InsertNewMessage(mongo.Message{
		MessageId: message.MessageId,
		ChatID:    request.ChatID,
		Deleted:   false,
	})
}

func SendMessageNative(message tgbotapi.MessageConfig) {

	responseMessage, err := telegramm_bot.Send(message)

	if err != nil {
		log.Default().Println(err)
	}

	mongo.InsertNewMessage(mongo.Message{
		MessageId: responseMessage.MessageID,
		ChatID:    responseMessage.Chat.ID,
		Deleted:   false,
	})

}

func DeleteMessages(chatid int64) {
	messages := mongo.GetMessagesById(chatid)

	for _, item := range messages {
		deleteMessageConfig := telegram_config.DeleteMessageConfig{
			ChatID:    item.ChatID,
			MessageID: item.MessageId,
		}

		params, _ := deleteMessageConfig.Params()
		response, err := telegramm_bot.MakeRequest(deleteMessageConfig.Method(), params)

		if err != nil {
			log.Default().Println(err)
		}

		response_bytes, err := response.Result.MarshalJSON()

		if err != nil {
			log.Default().Println(err)
		}

		var message mongo.MessageDTO
		json.Unmarshal(response_bytes, &message)
	}

	mongo.UpdateMessagesByChatID(chatid,
		bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{
						Key:   "delete",
						Value: true,
					},
				},
			}},
	)
}
