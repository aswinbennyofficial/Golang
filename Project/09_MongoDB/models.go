package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Restaurant struct{
	ID primitive.ObjectID `bson:"_id"`
	Name string
	Restaurant_ID string `bson:"restaurant_id"`
	Cuisine      string
	Address      interface{}
	Borough      string
	Grades       []interface{}
} 