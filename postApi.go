package postApi

import (
	"fmt"
	"net/http"
	"restful_api/customer_list"

	"restful_api/mongo_connect"
	"restful_api/packs_list"

	"github.com/gin-gonic/gin"
)

func PostApi(server *gin.Engine) {
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

	server.POST("/customers", func(c *gin.Context) {
		var customer customer_list.Customer
		if err := c.BindJSON(&customer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		mongo_connect.Insert_Data(database, collection_customers, customer)
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	server.POST("/packs", func(c *gin.Context) {
		var pack packs_list.Packs
		if err := c.BindJSON(&pack); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		mongo_connect.Insert_Data(database, collection_packs, pack)
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})
}
