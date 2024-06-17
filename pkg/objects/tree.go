package objects

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Amabeusz/vcs/pkg/common"
	"github.com/Amabeusz/vcs/pkg/file"
	"github.com/Amabeusz/vcs/pkg/global"
)

func createTree() string {
	content := file.ReadRoot(global.INDEX_FILE)

	newContent := []byte(removeLinesContainingHyphen(string(content)))

	fmt.Println("newContent: " + string(newContent))
	if len(content) != len(newContent) {
		file, err := os.OpenFile(common.GetRootPath()+"\\"+global.INDEX_FILE, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = file.WriteString(string(newContent))
		if err != nil {
			log.Fatal(err)
		}
	}

	return SaveObject(append([]byte("tree\n"), newContent...))
}

func removeLinesContainingHyphen(input string) string {
	lines := strings.Split(input, "\n")
	var result []string

	for _, line := range lines {
		if !strings.Contains(line, "-") {
			result = append(result, line)
		}
	}

	return strings.Join(result, "\n")
}

func ReadTree(sha string) map[string]string {
	content := strings.Fields(string(ReadObject(sha)))

	files := make(map[string]string, 0)

	if len(content) == 1 {
		return files
	}

	root := common.GetRootPath()
	for i := 1; i < len(content); i += 2 {
		if content[i] != "-" {
			files[root+"\\"+content[i+1]] = content[i]
		}
	}

	return files
}

func getTreeSha(s string) string {
	startIndex := strings.Index(s, "tree: ")
	if startIndex == -1 {
		return ""
	}

	startIndex += len("tree: ")

	remainingStr := s[startIndex:]

	endIndex := strings.Index(remainingStr, "\n")
	if endIndex == -1 {
		return strings.TrimSpace(remainingStr)
	}

	return strings.TrimSpace(remainingStr[:endIndex])
}
