package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func Connect() error {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/parser"))
	if err != nil {
		return err
	}
	Collection = client.Database("parser").Collection("parser")

	//res, err := Collection.InsertOne(ctx, structs.Line{
	//	PersonalAccount: "1",
	//	Name:            "2",
	//	Address:         "3",
	//	AccrualPeriod:   "4",
	//	Count:           "5",
	//	Number:          "6",
	//	Testimony:       "7",
	//})
	//if err != nil {
	//	return err
	//}

	//fmt.Println("res.InsertedID", res.InsertedID)

	return nil
}
