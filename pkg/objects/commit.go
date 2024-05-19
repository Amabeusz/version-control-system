package objects

import (
	"fmt"

	"github.com/Amabeusz/vcs/pkg/file"
	"github.com/Amabeusz/vcs/pkg/global"
)

func createCommit(title string) []byte {
	content := []byte("commit\n")

	author := "author"

	parent := getHeadSha()
	tree := createTree()

	content = append(content, []byte(fmt.Sprintf("tree: %v\n", tree))...)
	if len(parent) != 0 {
		content = append(content, []byte(fmt.Sprintf("parent: %v\n", parent))...)
	}
	content = append(content, []byte(fmt.Sprintf("author: %v\n\n", author))...)
	content = append(content, []byte(fmt.Sprintf("\t%v", title))...)

	return content
}

func getHeadSha() []byte {
	headContent := file.Read(global.HEAD_FILE)

	return file.Read(global.VCS_PATH + "\\" + string(headContent[5:]))
}
