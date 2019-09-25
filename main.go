package main

import(
	"github.com/gin-gonic/gin"
	"fileDownload/routes"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", routes.Home)
	router.GET("/download/:uuid", routes.Download)
	router.GET("/upload/:uuid", routes.Upload)
	router.GET("/IsDownload", routes.IsDownload)

	router.POST("/", routes.GenerateUUID)
	router.POST("/upload/:uuid", routes.UploadPost)

	router.Static("/zip", "./zip")
	router.Static("/src", "./src")

	router.Run(":8000")
}