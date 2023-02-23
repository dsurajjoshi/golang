package main

import (
	"fmt"
	"restful_api/getApi"
	"restful_api/postApi"
	"restful_api/updateApi"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Main File Activated")
	server := gin.Default()
	getApi.GetApi(server)
	postApi.PostApi(server)
	updateApi.UpdateApi(server)
	server.Run("localhost:9000")
}
