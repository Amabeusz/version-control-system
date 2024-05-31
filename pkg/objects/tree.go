package objects

import (
	"strings"

	"github.com/Amabeusz/vcs/pkg/common"
	"github.com/Amabeusz/vcs/pkg/file"
	"github.com/Amabeusz/vcs/pkg/global"
)

func createTree() string {
	return SaveObject(append([]byte("tree\n"), file.ReadRoot(global.INDEX_FILE)...))
}

func ReadTree(sha string) map[string]string {
	root := common.GetRootPath()
	content := strings.Fields(string(ReadObject(sha)))

	files := make(map[string]string, 0)

	for i := 1; i < len(content); i += 2 {
		files[root+"\\"+content[i+1]] = content[i]
	}

	return files
}
