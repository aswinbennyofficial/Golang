package main

import (
	"fmt"
	"log"
	"os"
	"context"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Review struct{
	Name string
	Text string
}

func main(){
	//load env variables
	err:=godotenv.Load(".env")
	if err != nil {
        log.Println("Error loading environment variables file")
		return
    }
	DB_URI:=os.Getenv("DB_URI")
	DB_NAME:=os.Getenv("DB_NAME")
	DB_COLLECTION_NAME:=os.Getenv("DB_COLLECTION_NAME")

	
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	
	// Configue settings related Mongodb client behaviour
	opts := options.Client().ApplyURI(DB_URI).SetServerAPIOptions(serverAPI)
	
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Create MongoDB collection obj
	coll:=client.Database(DB_NAME).Collection(DB_COLLECTION_NAME)

	readDB(coll)

}

func readDB(coll *mongo.Collection){
	// Creating filter
	filter:=bson.D{{"name","Rodrik Cassel"}}
	
	// Fetching result as BSON
	var result bson.D
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err!=nil{
		log.Println(err)
	}
	fmt.Println(result)

	// Unmarshalling BSON into Review struct
	var reviewInfo Review
	if err := coll.FindOne(context.TODO(), filter).Decode(&reviewInfo); err != nil {
		log.Println(err)
	}
	fmt.Println(reviewInfo)
}