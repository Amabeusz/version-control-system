package objects

import (
	"log"

	"github.com/Amabeusz/vcs/pkg/file"
	"github.com/Amabeusz/vcs/pkg/global"
)

func Head() []byte {
	head := file.Read(global.HEAD_FILE)

	if head == nil || string(head) == "" {
		log.Fatal("Head required")
	}

	return head
}
