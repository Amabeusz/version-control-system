package objects

import (
	"log"

	"github.com/Amabeusz/vcs/pkg/file"
	"github.com/Amabeusz/vcs/pkg/global"
)

func Head() []byte {
	head := file.ReadRoot(global.HEAD_FILE)

	if head == nil || string(head) == "" {
		log.Fatal("Head required")
	}

	return head
}

func GetHeadFiles() map[string]string {
	head := Head()

	ref := ReadRef(string(head))

	if len(ref) == 0 {
		return map[string]string{}
	}

	return ReadCommitTree(string(ref))
}

func ReadRef(ref string) []byte {
	path := global.REFS_PATH + ref
	return file.Read(path)
}
