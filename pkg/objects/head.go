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
	headRef := string(head[:len(head)-2])
	ref := ReadRef(headRef)

	return ReadTree(string(ref[:len(ref)-2]))
}

func ReadRef(ref string) []byte {
	path := global.REFS_PATH + ref
	return file.Read(path)
}
