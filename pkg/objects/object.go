package objects

import (
	"github.com/Amabeusz/vcs/pkg/common"
	"github.com/Amabeusz/vcs/pkg/file"
	"github.com/Amabeusz/vcs/pkg/global"
)

func ReadObject(objectSha string) []byte {
	content := file.Read(global.OBJECTS_PATH + objectSha)
	return common.Decompress(content)
}
