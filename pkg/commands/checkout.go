package commands

import (
	"fmt"
	"strings"

	"github.com/Amabeusz/vcs/pkg/file"
	"github.com/Amabeusz/vcs/pkg/objects"
)

func Checkout(arg string) {
	content := objects.ReadObject(arg)

	lines := strings.Split(string(content), "\n")

	if lines[0] != "commit" {
		fmt.Println("Only commits can be checked out")
		return
	}

	treeSha := lines[1][6:]

	tree := objects.ReadTree(treeSha)

	for k, v := range tree {
		fmt.Printf("k: %v, v: %v\n", k, v)
		file.CreateOrOverwrite(k, objects.ReadObject(v))
	}
}
