package mongo

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageDTO struct {
	MessageId int `json:"message_id"`
}

type Message struct {
	MessageId int
	ChatID    int64
	Deleted   bool
}

func InsertNewMessage(resp tgbotapi.Message) {

	message := Message{
		MessageId: resp.MessageID,
		ChatID:    resp.Chat.ID,
		Deleted:   false,
	}

	res, err := messages_collection.InsertOne(ctx, message)

	if err != nil {
		log.Default().Println(err)
	} else {
		log.Default().Println(res)
	}
}

func UpdateMessage() {
	messages_cursor, err := messages_collection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatalln(err)
	}

	var msgs []Message

	err = messages_cursor.All(ctx, &msgs)

	if err != nil {
		log.Fatalln(err)
	}

	// return msgs
}

func GetMessages() []Message {
	messages_cursor, err := messages_collection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatalln(err)
	}

	var msgs []Message

	err = messages_cursor.All(ctx, &msgs)

	if err != nil {
		log.Fatalln(err)
	}

	return msgs
}

func GetMessagesById(chatid int64) []Message {
	messages_cursor, err := messages_collection.Find(ctx, bson.D{{Key: "chatid", Value: chatid}})

	if err != nil {
		log.Fatalln(err)
	}

	var msgs []Message

	err = messages_cursor.All(ctx, &msgs)

	if err != nil {
		log.Fatalln(err)
	}

	return msgs
}

func UpdateMessagesByChatID(chatid int64, update primitive.D) {
	result, err := messages_collection.UpdateMany(ctx, bson.D{{Key: "chatid", Value: chatid}}, update)

	if err != nil {
		log.Default().Println(err)
	} else {
		log.Default().Println(result)
	}
}
