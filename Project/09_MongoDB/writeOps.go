package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func InsertADocument(coll *mongo.Collection) {
	// new instance of the NewRestaurant struct in models.go
	newRestaurant := NewRestaurant{Name: "8299", Cuisine: "Korean"}

	result, err := coll.InsertOne(context.TODO(), newRestaurant)
	if err != nil {
		panic(err)
	}

	log.Println(result.InsertedID)
}

func InsertManyDocument(coll *mongo.Collection){
	// new instance of the NewRestaurant struct in models.go
	newRestaurants := []interface{}{
		NewRestaurant{Name: "Rule of Thirds", Cuisine: "Japanese"},
		NewRestaurant{Name: "Madame Vo", Cuisine: "Vietnamese"},
	}

	// Inserts sample documents into the collection
	result, err := coll.InsertMany(context.TODO(), newRestaurants)
	if err != nil {
		panic(err)
	}

	log.Println(result.InsertedIDs...)


}

func UpdateADocument(coll *mongo.Collection,uid string){
	id, _ := primitive.ObjectIDFromHex(uid)
	filter := bson.D{{"_id", id}}
	// Creates instructions to add the "avg_rating" field to documents
	update := bson.D{{"$set", bson.D{{"avg_rating", 4.4}}}}
	// Updates the first document that has the specified "_id" value
	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	log.Println("Modified count: ",result.ModifiedCount)
}

func UpdateManyDocument(coll *mongo.Collection){
	// filter := bson.D{{"address.market", "Sydney"}}
	filter := bson.D{{"cuisine", "Bakery"}}
	// Creates instructions to update the values of the "price" field
	update := bson.D{{"$set", bson.D{{"restaurant_id", 1.15}}}}
	// Updates documents in which the value of the "address.market"
	// field is "Sydney"
	result, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	log.Println(result.ModifiedCount)
}

