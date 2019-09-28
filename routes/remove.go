package routes

import(
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func Remove(c *gin.Context) {
	uuid := c.Query("uuid")
	path := "./zip/" + uuid + ".zip"
	err := os.Remove(path)
	if err != nil {
		go removeZipSecond(path)
	}
}

func removeZipSecond(path string) {
	time.Sleep(1 * time.Second)
	err := os.Remove(path)
	if err == nil{
		removeZipSecond(path)
	}
}