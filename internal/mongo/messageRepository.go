package mongo

import (
	"context"
	"log"

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

func InsertNewMessage(message Message) {
	ctx := context.Background()
	messages := client.Database("main").Collection("messages")
	res, err := messages.InsertOne(ctx, message)

	if err != nil {
		log.Default().Println(err)
	} else {
		log.Default().Println(res)
	}
}

func UpdateMessage() {
	ctx := context.Background()
	messages := client.Database("main").Collection("messages")
	messages_cursor, err := messages.Find(ctx, bson.D{})

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
	ctx := context.Background()
	messages := client.Database("main").Collection("messages")
	messages_cursor, err := messages.Find(ctx, bson.D{})

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
	ctx := context.Background()
	messages := client.Database("main").Collection("messages")
	messages_cursor, err := messages.Find(ctx, bson.D{{"chatid", chatid}})

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
	ctx := context.Background()
	messages := client.Database("main").Collection("messages")
	result, err := messages.UpdateMany(ctx, bson.D{{"chatid", chatid}}, update)

	if err != nil {
		log.Default().Println(err)
	} else {
		log.Default().Println(result)
	}
}
