package objects

import (
	"fmt"
	"os"

	"github.com/Amabeusz/vcs/pkg/common"
	"github.com/Amabeusz/vcs/pkg/file"
	"github.com/Amabeusz/vcs/pkg/global"
)

func RefSha(ref string) []byte {
	return file.ReadRoot(global.REFS_PATH + "master")
}

func UpdateRef(commitSha []byte) {
	headSha := Head()
	head := common.GetRootPath() + "\\" + global.REFS_PATH + string(headSha)

	err := os.WriteFile(head, commitSha, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
