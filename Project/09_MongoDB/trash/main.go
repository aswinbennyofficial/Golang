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

type Inspection struct{
	Business_name string
	Result string
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

	//readOneDB(coll)
	//readManyDB(coll)

	// MongoDoc
	mongoOfficialRead(DB_NAME, DB_COLLECTION_NAME,coll)
}



func readManyDB(coll *mongo.Collection){

	// Creating filter
	filter:=bson.D{{"result","No Violation Issued"}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.TODO())

	// Create instance of struct
	var inspectionResults []Inspection

	// Using cursor All
	if err := cursor.All(context.TODO(), &inspectionResults); err != nil {
		panic(err)
	}

	// Using collection next
	for cursor.Next(context.TODO()) {
		var inspectionResult Inspection
		if err := cursor.Decode(&inspectionResult); err != nil {
			panic(err)
		}
		inspectionResults = append(inspectionResults, inspectionResult)
	}
	


	fmt.Println(inspectionResults)

}

func readOneDB(coll *mongo.Collection){
	

	// Creating filter
	filter:=bson.D{{"result","No Violation Issued"}}
	
	// Fetching result as BSON
	// var result bson.D
	// err := coll.FindOne(context.TODO(), filter).Decode(&result)
	// if err!=nil{
	// 	log.Println(err)
	// }
	// fmt.Println(result)

	// Unmarshalling BSON into Review struct
	var inspetionInfo Inspection
	if err := coll.FindOne(context.TODO(), filter).Decode(&inspetionInfo); err != nil {
		log.Println(err)
	}
	fmt.Println(inspetionInfo)
}