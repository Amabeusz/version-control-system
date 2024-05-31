package commands

import (
	"log"
	"os"

	"github.com/Amabeusz/vcs/pkg/common"
	"github.com/Amabeusz/vcs/pkg/objects"
)

func Add(arg string) {
	if arg[:1] == "." {
		log.Fatal("Not implemented yet")
		return
	}

	if _, err := os.Stat(common.GetRootPath() + "\\" + arg); err != nil {
		headFiles := objects.GetHeadFiles()
		_, ok := headFiles[common.GetRootPath()+"\\"+arg]
		if ok {
			objects.AddBlobToIndex("-", arg)
			return
		}

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
