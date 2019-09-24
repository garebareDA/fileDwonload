package routes

import(
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"encoding/base64"
	"fmt"
)

func Download(c *gin.Context) {
	uuidPram := c.Param("uuid")

	u, err := uuid.NewRandom()
	if err != nil{
		fmt.Println(err)
	}

	var png []byte
	png, err = qrcode.Encode("http://localhost:8000/upload/" + u.String(), qrcode.Medium, 256)
	encoded := base64.StdEncoding.EncodeToString(png)

	fmt.Println(uuidPram)
	fmt.Println(encoded)
}