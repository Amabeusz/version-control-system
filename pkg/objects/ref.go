package objects

import (
	"os"

	"github.com/Amabeusz/vcs/pkg/file"
	"github.com/Amabeusz/vcs/pkg/global"
)

func RefSha(ref string) []byte {
	return file.ReadRoot(global.REFS_PATH + "master")
}

func UpdateRef(head string, commitSha []byte) {
	os.WriteFile(global.REFS_PATH+head, commitSha, 0644)
}
