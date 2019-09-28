package main

import(
	"fmt"
	"bufio"
	"os"
	"io"
	"io/ioutil"
	"archive/zip"
)

func main(){
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()

	files, err := ioutil.ReadDir("./" + text)
	if err != nil {
		fmt.Println(err)
	}

	path := "./zip/" + text + ".zip"
	dest, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer dest.Close()


	zipWriter := zip.NewWriter(dest)
	defer zipWriter.Close()

	for _, file := range files{
		path := "./" + text + "/" + file.Name()
		addToZip(path, zipWriter)
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