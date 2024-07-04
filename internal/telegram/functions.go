package telegram

import (
	"encoding/json"
	"log"

	config "elerphore.com/flower-journal/internal/config"
	"elerphore.com/flower-journal/internal/mongo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func SendMessageNative(message tgbotapi.MessageConfig) *tgbotapi.Message {

	responseMessage, err := telegramm_bot.Send(message)

	if err != nil {
		log.Default().Println(err)
	}

	return &responseMessage

}

func DeleteMessages(chatid int64) {
	messages := mongo.GetMessagesById(chatid)

	for _, item := range messages {
		deleteMessageConfig := config.DeleteMessageConfig{
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
