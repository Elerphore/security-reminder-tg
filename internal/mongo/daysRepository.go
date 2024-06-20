package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Day struct {
	Day_Index int
	Used      bool
	User_ID   primitive.ObjectID
}

func SetAllDaysToDefaultState() {
	ctx := context.Background()
	daysCollection := client.Database("main").Collection("days")

	res, err := daysCollection.UpdateMany(
		ctx,
		bson.D{},
		bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{
						Key:   "used",
						Value: false,
					},
				},
			},
		},
	)

	if err != nil {
		log.Default().Println(err)
	} else {
		log.Default().Println(res)
	}
}

func FindCurrentDayByUserId(userID primitive.ObjectID) Day {
	ctx := context.Background()
	daysCollection := client.Database("main").Collection("days")

	var day Day

	err := daysCollection.FindOne(
		ctx,
		bson.D{
			{
				Key:   "user_id",
				Value: userID,
			},
			{
				Key:   "day_index",
				Value: int(time.Now().Weekday()),
			},
		},
	).Decode(&day)

	if err != nil {
		log.Default().Println(err)
	}

	return day
}

func UpdateCurrentDay(userID primitive.ObjectID) {
	ctx := context.Background()
	daysCollection := client.Database("main").Collection("days")

	err := daysCollection.FindOneAndUpdate(
		ctx,
		bson.D{
			{
				Key:   "user_id",
				Value: userID,
			},
			{
				Key:   "day_index",
				Value: int(time.Now().Weekday()),
			},
		},
		bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{
						Key:   "used",
						Value: true,
					},
				},
			},
		},
	)

	if err != nil {
		log.Default().Println(err)
	}
}
