package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"fmt"
	"net/http"
)

func GenerateUUID(c *gin.Context) {
	u, err := uuid.NewRandom()
	if err != nil{
		fmt.Println(err)
	}

	c.Redirect(http.StatusFound, "/download/" + u.String())
	c.Abort()
}