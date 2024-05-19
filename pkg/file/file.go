package file

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Amabeusz/vcs/pkg/common"
)

func Read(filePath string) []byte {
	path := common.GetRootPath() + "\\" + filePath
	fmt.Println(path)
	file, err := os.Open(path)
	common.Check(err)
	defer file.Close()

	content, err := io.ReadAll(file)
	common.Check(err)

	return content
}

func Create(filePath string, content []byte) {
	path := common.GetRootPath() + "\\" + filePath
	fmt.Println(path)
	file, err := os.Create(filePath)
	common.Check(err)
	defer file.Close()

	_, err = file.Write(content)
	common.Check(err)

	file.Sync()

	log.Printf("File created successfully: %v", filePath)
}

func FindAndWrite(filePath string, content []byte) {
	path := common.GetRootPath() + "\\" + filePath
	fmt.Println(path)
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	common.Check(err)
	defer file.Close()

	Write(file, content)
}

func Write(file *os.File, content []byte) {
	_, err := file.Write(content)
	common.Check(err)

	file.Sync()
}
