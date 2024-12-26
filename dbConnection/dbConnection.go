package dbconnection

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "userDb"
const collectionName = "userList"
const newsDb = "newsDb"
const newsCollection = "newsCollection"

var collection *mongo.Collection
var client *mongo.Client
var NewsCollection *mongo.Collection
var UserCollection *mongo.Collection

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("some error occur in env", err)
	}

	dt := os.Getenv("MONGO_URI")
	if dt != "" {
		//todo make connection with mongo db here
		connectionString := dt
		clientOption := options.Client().ApplyURI(connectionString)
		client, err = mongo.Connect(context.TODO(), clientOption)
		if err != nil {
			log.Fatal(err)
		}
		client.Database(dbName).Collection(collectionName)
		client.Database(newsDb).Collection(newsCollection)
		NewsCollection = client.Database(newsDb).Collection(newsCollection)
		UserCollection = client.Database(dbName).Collection(collectionName)
		fmt.Println("collection instanace is ready")
		fmt.Println("collection", collection)
		fmt.Println("connection success", client)
		// Defer the disconnect operation to ensure it's always called when the main function exits
		// defer func() {
		// 	if err := client.Disconnect(context.Background()); err != nil {
		// 		log.Fatal(err)
		// 	}
		// }()

	} else {
		panic("dt is null")
	}

}
