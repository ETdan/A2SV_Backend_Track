package main

import (
	"TaskManagerWithMongoDB/data"
	"TaskManagerWithMongoDB/router"

	"github.com/gin-gonic/gin"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	/*
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
		} else {
			fmt.Println("Connected to MongoDB!")
		}
		// Check the connection
		if err = client.Ping(context.TODO(), nil); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
		}
		// disconnect with the db
		defer func() {
			if err = client.Disconnect(context.TODO()); err != nil {
				panic(err)
			} else {
				fmt.Println("Connection to MongoDB closed.")
			}
		}()

		// connect to db test and then represent trainers by colleciton
		Collection := client.Database("taskManager").Collection("tasks")
		// t := map[string]models.Task{
		// 	"1": {
		// 		ID:       "1",
		// 		Name:     "Task 1",
		// 		Detail:   "Detail for Task 1",
		// 		Start:    "2024-07-31T08:00:00Z",
		// 		Duration: "1h",
		// 	},
		// }
		// Collection.InsertOne(context.TODO(), t)
		fmt.Println(Collection.CountDocuments(context.TODO(), bson.D{{}}))
	*/
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
	data.StartMongoDB()
	r := gin.Default()
	r = router.Router(r)
	r.Run()

}
