package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/Amabeusz/vcs/pkg/objects"
)

func CatObjectType(flag string, objectSha string) {
	if flag == "-t" {
		content := objects.ReadObject(objectSha)

		fmt.Println(getFirstLine(string(content)))
		return
	}

	log.Fatal("Flag not recognized")
}

func CatObject(objectSha string) {
	content := objects.ReadObject(objectSha)

	fmt.Println(removeFirstLine(string(content)))
}

func getFirstLine(s string) string {
	lines := strings.SplitN(s, "\n", 2)
	return lines[0]
}

func removeFirstLine(s string) string {
	lines := strings.SplitN(s, "\n", 2)

	if len(lines) < 2 {
		return ""
	}

	return lines[1]
}
