package routes

import(
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"github.com/boombuler/barcode/qr"
	"fmt"
)

func Download(c *gin.Context) {
	uuidPram := c.Param("uuid")

	u, err := uuid.NewRandom()
	if err != nil{
		fmt.Println(err)
	}

	qrCode, _ := qr.Encode("http://localhost:8000/upload/" + u.String(), qr.L, qr.Auto)

	fmt.Println(qrCode)
	fmt.Println(uuidPram)
}