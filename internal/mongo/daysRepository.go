package mongo

import (
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
	res, err := days_collection.UpdateMany(
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
	var day Day

	err := days_collection.FindOne(
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
	err := days_collection.FindOneAndUpdate(
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
