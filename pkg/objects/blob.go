package objects

import (
	"github.com/Amabeusz/vcs/pkg/common"
	"github.com/Amabeusz/vcs/pkg/file"
	"github.com/Amabeusz/vcs/pkg/global"
)

func createBlob(fileContent []byte) []byte {
	return common.Compress(append([]byte("blob\n"), fileContent...))
}

func CreateBlobFile(fileSha string, content []byte) {
	file.Create(global.OBJECTS_PATH+fileSha, createBlob(content))
}
