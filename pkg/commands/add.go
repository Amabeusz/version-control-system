package commands

import (
	"log"

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
	fileSha := objects.SaveBlob(filePath)
	objects.AddBlobToIndex(fileSha, filePath)
}
