package mongo_connect

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect_to_mongo() (*mongo.Client, error) {
	fmt.Println("Program Started..... ")
	// Set client Options:
	clientOptions := options.Client().ApplyURI("mongodb+srv://dth_sample_database:joshi311500@dth.jiqdzc5.mongodb.net/test")

	// Connect to mongodb:
	client, err := mongo.Connect(context.TODO(), clientOptions)

	// Checking for error while making client:
	if err != nil {
		log.Fatal("Error while connecting to mongodb: ", err)
	}

	// Checking the connection:
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Connection Testing Failed with error: ", err)
	}

	fmt.Println("Connected to the Mongo DB")

	return client, nil

}

// Fetching Data from MongoDB (GET API):
func Fetch_Data(database, collection string) (*mongo.Cursor, error) {
	findOptions := options.Find()
	fetch_client, _ := Connect_to_mongo()
	fetch_collection := fetch_client.Database(database).Collection(collection)
	cur, err := fetch_collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal("Error while fetching Data from the collection", err)
	}
	return cur, err
}

// Inserting Data into MongoDB (POST API):
func Insert_Data(database, collection string, insert_record interface{}) {
	insert_client, _ := Connect_to_mongo()
	insert_collection := insert_client.Database(database).Collection(collection)
	_, err := insert_collection.InsertOne(context.TODO(), insert_record)
	if err != nil {
		log.Fatal("Error while inserting Document in the collection with error: ", err)
	} else {
		fmt.Println("Document Inserted into the collection")
	}
}

// Updating Data into MongoDB (PUT API)
func Update_Data(id, database, collection string, update_record interface{}) {
	converted_id, _ := primitive.ObjectIDFromHex(id)
	update_client, _ := Connect_to_mongo()
	//id_var := "_id"
	filter := bson.M{"_id": converted_id}
	update_collection_record := update_client.Database(database).Collection(collection)
	_, err := update_collection_record.UpdateOne(context.TODO(), filter, update_record)

	if err != nil {
		log.Fatal("Error while updating record: ", err)
	} else {
		fmt.Println("Document Updation Successful")
	}

}

// Deleting Data from MongoDB (DELETE API):
func Delete_Data(id, database, collection string) {
	converted_id, _ := primitive.ObjectIDFromHex(id)
	delete_client, _ := Connect_to_mongo()
	filter := bson.M{"_id": converted_id}
	delete_collection_record := delete_client.Database(database).Collection(collection)
	_, err := delete_collection_record.DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Fatal("Error while deleting record: ", err)
	} else {
		fmt.Println("Document Deletion Successful")
	}
}
