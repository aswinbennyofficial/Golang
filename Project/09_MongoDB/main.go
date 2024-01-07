package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"context"
	
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


	// Creating a mongodb client using Db() function in db.go
	client:=Db(DB_URI)
	
	// Create MongoDB collection obj
	coll:=client.Database(DB_NAME).Collection(DB_COLLECTION_NAME)
	
	
	// CRUD operations

	// FIND ops
	// Read one document
	//ReadOneInDB(coll)

	// Read multiple document
	//ReadManyInDB(coll)


	// WRITE ops
	// Insert a document
	//InsertADocument(coll)

	// Insert multiple documents
	//InsertManyDocument(coll)

	// Update a document
	// uid:="659a9b15aaf1d2c2408e0365"
	// UpdateADocument(coll,uid)

	// Update many document
	//UpdateManyDocument(coll)

	

	// Defer disconnecting from the MongoDB client
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}