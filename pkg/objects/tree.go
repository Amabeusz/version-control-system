package objects

import (
	"github.com/Amabeusz/vcs/pkg/common"
	"github.com/Amabeusz/vcs/pkg/file"
	"github.com/Amabeusz/vcs/pkg/global"
)

func createTree() []byte {
	content := append([]byte("tree\n"), file.Read(global.INDEX_FILE)...)

	sha := common.FileSha(content)

	file.Create(global.OBJECTS_PATH+sha, content)

	return []byte(sha)
}
