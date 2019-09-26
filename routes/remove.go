package routes

import(
	"github.com/gin-gonic/gin"
	"os"
	"time"
	"fmt"
)

func Remove(c *gin.Context) {
	uuid := c.Query("uuid")
	path := "./zip/" + uuid + ".zip"
	fmt.Println(path)
	_, err := os.Stat(path)
	if err == nil {
		fmt.Println("remove start")
		removeZipSecond(path)
	}
}

func removeZipSecond(path string) {
	time.Sleep(1 * time.Second)
	err := os.Remove(path)
	if err != nil{
		fmt.Println(err)
		go removeZipSecond(path)
	}else{
		return
	}
}