package objects

import (
	"fmt"

	"github.com/Amabeusz/vcs/pkg/config"
)

func CreateCommit(msg string) {
	content := []byte("commit\n")

	head := string(Head())
	parent := RefSha(head)
	tree := createTree()

	content = append(content, []byte(fmt.Sprintf("tree: %v\n", tree))...)
	if parent != nil {
		content = append(content, []byte(fmt.Sprintf("parent: %v\n", string(parent)))...)
	}
	content = append(content, []byte(fmt.Sprintf("author: %v\n\n", config.User()))...)
	content = append(content, []byte(msg)...)

	sha := SaveObject(content)
	UpdateRef([]byte(sha))
}

func ReadCommitTree(sha string) map[string]string {
	headCommitContent := string(ReadObject(sha))
	treeSha := getTreeSha(headCommitContent)

	return ReadTree(treeSha)
}
