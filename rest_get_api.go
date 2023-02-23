package getApi

import (
	"context"
	"fmt"
	"net/http"
	_ "restful_api/customer_list"
	"restful_api/mongo_connect"
	_ "restful_api/packs_list"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetApi(server *gin.Engine) {
	fmt.Println("Initiating Restful GET API")

	//server := gin.Default()

	mongo_client, _ := mongo_connect.Connect_to_mongo()

	if mongo_client != nil {
		fmt.Println("Connected to the MongoDB")
	} else {
		fmt.Println("Error connecting to MongoDB")
	}

	// Providing Database and Collection Name:
	database := "dth_logs"
	collection_customers := "customer_list"
	collection_packs := "packs_list"

	// Fetching Data from MongoDB:
	fetched_customers, _ := mongo_connect.Fetch_Data(database, collection_customers)
	fetched_packs, _ := mongo_connect.Fetch_Data(database, collection_packs)
	defer fetched_customers.Close(context.Background())
	defer fetched_packs.Close(context.Background())

	// Getting customers details:
	for fetched_customers.Next(context.Background()) {
		var cust_document bson.M
		fetched_customers.Decode(&cust_document)
		server.GET("/customers", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H(cust_document))
		})

		fmt.Println(cust_document)
	}

	// Getting packs details:
	for fetched_packs.Next(context.Background()) {
		var pac_document bson.M
		fetched_packs.Decode(&pac_document)
		server.GET("/packs", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H(pac_document))
		})

		fmt.Println(pac_document)
	}

}
