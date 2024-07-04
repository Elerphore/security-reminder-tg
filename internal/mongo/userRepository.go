package mongo

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                  primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Telegram_User_ID    int64
	Telegram_Chat_ID    int64
	Telegram_Message_ID int
}

func InsertNewUser(update tgbotapi.Update, msg tgbotapi.Message) {

	var user = User{
		ID:                  primitive.NewObjectID(),
		Telegram_User_ID:    update.SentFrom().ID,
		Telegram_Chat_ID:    update.FromChat().ID,
		Telegram_Message_ID: msg.MessageID,
	}

	insertNewUser(user)
}

func UpdateUserMessageID(user_id primitive.ObjectID, message_id int) {
	_, err := user_collection.UpdateOne(ctx,
		bson.D{{
			Key:   "_id",
			Value: user_id,
		}},
		bson.D{
			{
				Key: "$set",
				Value: bson.D{{
					Key:   "telegram_message_id",
					Value: message_id,
				}},
			},
		},
	)

	if err != nil {
		log.Fatalln(err)
	}
}

func insertNewUser(user User) {

	result := user_collection.FindOne(ctx, bson.D{{Key: "telegram_user_id", Value: user.Telegram_User_ID}})

	if result.Err() != nil {
		res, err := user_collection.InsertOne(ctx, user)

		if err != nil {
			log.Default().Println(err)
		}

		user_id := res.InsertedID.(primitive.ObjectID)

		days := []interface{}{}

		for i := 0; i < 7; i++ {
			days = append(days,
				Day{
					Day_Index: i,
					Used:      false,
					User_ID:   user_id,
				})
		}

		days_collection.InsertMany(ctx, days)

		if err != nil {
			log.Default().Println(err)
		}

		defaultLog.Println("New User added: {}", res)
	}

}

func GetUsers() []User {

	users_cursor, err := user_collection.Find(ctx, bson.D{})

	var users []User

	if err != nil {
		log.Fatalln(err)
	}

	err = users_cursor.All(ctx, &users)

	if err != nil {
		log.Fatalln(err)
	}

	return users
}

func GetUserByTelegramUserId(telegramUsedID int64) User {
	var user User

	err := user_collection.FindOne(
		ctx,
		bson.D{
			{
				Key:   "telegram_user_id",
				Value: telegramUsedID,
			},
		},
	).Decode(&user)

	if err != nil {
		log.Fatalln(err)
	}

	return user
}
