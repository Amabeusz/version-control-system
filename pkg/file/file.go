package file

import (
	"io"
	"log"
	"os"

	"github.com/Amabeusz/vcs/pkg/common"
)

func ReadRoot(filePath string) []byte {
	path := common.GetRootPath() + "\\" + filePath
	return Read(path)
}

func Read(filePath string) []byte {
	file, err := os.Open(filePath)
	common.Check(err)
	defer file.Close()

	content, err := io.ReadAll(file)
	common.Check(err)

	return content
}

func Create(filePath string, content []byte) {
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
