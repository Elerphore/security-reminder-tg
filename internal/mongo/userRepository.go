package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID               primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Telegram_User_ID int64
	Telegram_Chat_ID int64
	Send_Time        string
}

func InsertNewUser(user User) {
	ctx := context.Background()
	user_collection := client.Database("main").Collection("user")

	result := user_collection.FindOne(ctx, bson.D{{Key: "telegram_user_id", Value: user.Telegram_User_ID}})

	if result.Err() != nil {
		days_collection := client.Database("main").Collection("days")
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
	ctx := context.Background()
	user_collection := client.Database("main").Collection("user")
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
	ctx := context.Background()
	user_collection := client.Database("main").Collection("user")

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
