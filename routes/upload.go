package routes

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"archive/zip"
	"io"
	"os"
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

	dest, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer dest.Close()

	zipWriter := zip.NewWriter(dest)
	defer zipWriter.Close()

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

		err = addToZip(path, zipWriter)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = os.RemoveAll("./" + uuid)
	if err != nil {
		fmt.Println(err)
	}
}

func addToZip(filename string, zipWriter *zip.Writer) error {
	src, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer src.Close()

	writer, err := zipWriter.Create(filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, src)
	if err != nil {
		return err
	}

	return nil
}