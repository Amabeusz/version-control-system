package objects

import (
	"github.com/Amabeusz/vcs/pkg/common"
	"github.com/Amabeusz/vcs/pkg/file"
	"github.com/Amabeusz/vcs/pkg/global"
)

func ReadObject(objectSha string) []byte {
	content := file.ReadRoot(global.OBJECTS_PATH + objectSha)
	return common.Decompress(content)
}

func SaveObject(content []byte) string {
	objectSha := common.FileSha(content)
	file.Create(global.OBJECTS_PATH+objectSha, common.Compress(content))
	return objectSha
}
