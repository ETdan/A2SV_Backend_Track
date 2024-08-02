package main

import (
	// "TaskManagerWithMongoDB/router"
	"context"
	"fmt"
	"log"

	// "github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	// ash := Trainer{"Ash", 10, "Pallet Town"}
	// misty := Trainer{"Misty", 10, "Cerulean City"}
	// brock := Trainer{"Brock", 15, "Pewter City"}

	// trainers := []interface{}{misty, brock}

	// Set client options

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI("mongodb+srv://ETdan:kRPGzScrfbHSH4Gt@cluster0.f79ysrp.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetServerAPIOptions(serverAPI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// connect to db test and then represent trainers by colleciton
	// collection := client.Database("test").Collection("trainers")
	// fmt.Println(collection.CountDocuments(context.TODO(), bson.D{{}}))
	/*
		// insert one document(record)
		result, err := collection.InsertOne(context.TODO(), ash)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(result.InsertedID)
		}

		// insert many documents(records)
		resultMany, err := collection.InsertMany(context.TODO(), trainers)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(resultMany.InsertedIDs)
		}
	*/
	/*
		// update
			update := bson.D{
				{"$inc", bson.D{
					{"age", 1},
				},
				},
			}
			filter := bson.D{{"name", "Ash"}}

			result, err := collection.UpdateMany(context.TODO(), filter, update)
			if err == nil {
				fmt.Printf("%v %v %v %v", result.MatchedCount, result.ModifiedCount, result.UpsertedCount, result.UpsertedID)
			} else {
				log.Fatal(err)
			}*/
	/*
			// find
			// find one
			var result Trainer
			err = collection.FindOne(context.TODO(), bson.D{{"name", "Ash"}}).Decode(&result)
			if err != nil {
				log.Fatal("record not found")
			} else {
				fmt.Println(result)
			}

		// find many
		var results []Trainer
		findoptions := options.Find()
		findoptions.SetLimit(3)
		if cursur, err := collection.Find(context.TODO(), bson.D{{}}, findoptions); err != nil {
			log.Fatal("error in find")
		} else {

			for cursur.Next(context.TODO()) {
				var temp Trainer
				err := cursur.Decode(&temp)
				if err == nil {
					results = append(results, temp)
				} else {
					log.Fatal("cursur Error")
				}
			}
			fmt.Println(results)
		}
	*/
	// disconnect with the db
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")

	// r := gin.Default()
	// r = router.Router(r)
	// r.Run()

}
