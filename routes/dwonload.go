package routes

import(
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"encoding/base64"
	"net/http"
	"os"
	"fmt"
)

func Download(c *gin.Context) {
	uuidPram := c.Param("uuid")
	fmt.Println(uuidPram)

	u, err := uuid.NewRandom()
	if err != nil{
		fmt.Println(err)
	}

	var png []byte
	png, err = qrcode.Encode("http://localhost:8000/upload/" + u.String(), qrcode.Medium, 256)
	encoded := base64.StdEncoding.EncodeToString(png)

	c.HTML(http.StatusOK, "download.html",gin.H{
		"qr":encoded,
		"uuid":u,
	})
}

func IsDownload(c *gin.Context) {
	uuid := c.Query("uuid")
	_, err := os.Stat("./zip/" + uuid + ".zip")

	if err == nil {
		fmt.Println("true")
		c.String(201,"true")
	}else{
		c.String(201,"false")
	}
}

func StartDwonload(c *gin.Context) {
	uuid := c.Query("uuid")
	path := "./zip/" + uuid + ".zip"

	c.Header("Access-Control-Allow-Origin", "http://localhost:8000")
	c.Header("Access-Control-Expose-Headers", "Content-Disposition")
	c.Header("Content-Disposition", "attachment; filename=" + uuid + ".zip" )
	c.Header("Content-Type", "application/octet-stream")
	c.File(path)
}