package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func ReadOneInDB(coll *mongo.Collection) {

	// Creating a filter
	filter := bson.D{{"name", "Bagels N Buns"}}
	
	// instance of struct Restaurant in the model
	var result Restaurant

	err:=coll.FindOne(context.TODO(),filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		panic(err)
	}

	log.Println("Result : ", result)


}


func ReadManyInDB(coll *mongo.Collection){

	//filter
	filter := bson.D{{"cuisine", "Italian"}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	// Iterate over the cursor and print the results
	var results []Restaurant
	for cursor.Next(context.TODO()) {
		var result Restaurant
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	// Check for errors from iterating over the cursor
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once done
	cursor.Close(context.TODO())

	// Print the results
	log.Println(results)

}