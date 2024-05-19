package commands

import (
	"log"

	"github.com/Amabeusz/vcs/pkg/common"
	"github.com/Amabeusz/vcs/pkg/file"
	"github.com/Amabeusz/vcs/pkg/objects"
)

func Add(arg string) {
	if arg[:1] == "." {
		log.Fatal("Not implemented yet")
		return
	}

	addFile(arg)
}

func addFile(filePath string) {
	addBlobFile(filePath)
}

func addBlobFile(filePath string) {
	content := file.Read(filePath)
	fileSha := common.FileSha(content)

	objects.CreateBlobFile(fileSha, content)
	objects.AddBlobToIndex(fileSha, filePath)
}
