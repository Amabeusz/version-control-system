package commands

import (
	"fmt"
	"strings"

	"github.com/Amabeusz/vcs/pkg/objects"
)

func Log() {
	head := objects.Head()

	refSha := objects.RefSha(string(head))

	if refSha == nil {
		fmt.Println("No commits yet")
	}

	printLog(string(refSha))
}

func printLog(commitSha string) {
	for commitSha != "" {
		content := string(objects.ReadObject(string(commitSha)))

		fmt.Println(content[strings.Index(content, "\n"):])

		commitSha = getParentSha(content)
	}
}

func getParentSha(s string) string {
	startIndex := strings.Index(s, "parent: ")
	if startIndex == -1 {
		return ""
	}

	startIndex += len("parent: ")

	remainingStr := s[startIndex:]

	endIndex := strings.Index(remainingStr, "\n")
	if endIndex == -1 {
		return strings.TrimSpace(remainingStr)
	}

	return strings.TrimSpace(remainingStr[:endIndex])
}
