package objects

import (
	"github.com/Amabeusz/vcs/pkg/file"
	"github.com/Amabeusz/vcs/pkg/global"
)

func createTree() string {
	return SaveObject(append([]byte("tree\n"), file.Read(global.INDEX_FILE)...))
}
