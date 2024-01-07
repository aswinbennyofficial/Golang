package main

import (
	//"fmt"
	//"log"
	//"os"
	"context"
	"fmt"
	//"log"

	//"github.com/joho/godotenv"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	
)

type Tea struct {
	Type   string
	Rating int32
	Vendor []string `bson:"vendor,omitempty" json:"vendor,omitempty"`
}

func mongoOfficialRead(DB_NAME string, DB_COLLECTION_NAME string, coll *mongo.Collection){

	matchStage := bson.D{{"$match", bson.D{{"rating", bson.D{{"$gt", 5}}}}}}
countStage := bson.D{{"$count", "counted_documents"}}

cursor, err := coll.Aggregate(context.TODO(), mongo.Pipeline{matchStage, countStage})
if err != nil {
   panic(err)
}

var results []bson.D
if err = cursor.All(context.TODO(), &results); err != nil {
   panic(err)
}
for _, result := range results {
   fmt.Println(result)
}
}