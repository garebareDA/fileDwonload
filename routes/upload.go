package routes

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"os"
	"os/exec"
	"io"
	"time"
	"fmt"
)

func Upload(c *gin.Context) {
	uuid := c.Param("uuid")
	c.HTML(http.StatusOK,"upload.html",gin.H{
		"uuid":uuid,
	})
}

func UploadPost(c *gin.Context) {
	uuid := c.Param("uuid")
	path := "./zip/" + uuid + ".zip"

	form, err := c.MultipartForm()
	if err != nil {
		panic(err)
	}

	os.Mkdir(uuid, 0755)

	files := form.File["files"]

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		path := "./" + uuid + "/" + filename

		err := c.SaveUploadedFile(file, path)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	cmd := exec.Command("go", "run", "./child/zip.go")
	stdin, _ := cmd.StdinPipe()
	io.WriteString(stdin, uuid)
	stdin.Close()

	cmd.Run()
	cmd.Wait()
	cmd.Process.Kill()

	err = os.RemoveAll("./" + uuid)
	if err != nil {
		fmt.Println(err)
	}

	go RemoveHour(path)

	c.String(201, "アップロードが完了しました")
}

func RemoveHour(path string) {
	time.Sleep(1 * time.Hour)
	os.Remove(path)
}