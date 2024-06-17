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

		lines := strings.Split(content, "\n")

		fmt.Printf("\033[33m%s\033[0m\n", commitSha)
		fmt.Printf("%v\n", lines[3])
		fmt.Printf("message: %v\n\n", lines[5])

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
