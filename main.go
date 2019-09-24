package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"fileDownload/routes"
)

func main() {
	fmt.Println("helloworld")
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", routes.Home)
	router.GET("/download/:uuid", routes.Download)
	router.GET("upload/:uuid")

	router.POST("/", routes.GenerateUUID)

	router.Static("/assets", "./assets")

	router.Run(":8000")
}