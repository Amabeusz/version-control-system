package commands

import (
	"fmt"
	"strings"

	"github.com/Amabeusz/vcs/pkg/common"
	"github.com/Amabeusz/vcs/pkg/file"
	"github.com/Amabeusz/vcs/pkg/objects"
)

func Diff() {
	files := getAllFiles(common.GetRootPath())
	headFiles := objects.GetHeadFiles()

	for _, f := range files {
		fileContent := string(file.Read(f))

		headSha := headFiles[f]

		if headSha == "" {
			continue
		}

		differences := findLineDifferences(string(fileContent), string(objects.ReadObject(headSha)))

		if len(differences) > 0 {
			fmt.Println(f)
			for _, diff := range differences {
				fmt.Println(diff)
			}
		}
	}
}

func findLineDifferences(str1, str2 string) []string {
	lines1 := strings.Split(str1, "\n")
	lines2 := strings.Split(str2, "\n")

	maxLen := len(lines1)
	if len(lines2) > maxLen {
		maxLen = len(lines2)
	}

	differences := []string{}
	for i := 0; i < maxLen; i++ {
		line1 := ""
		line2 := ""

		if i < len(lines1) {
			line1 = lines1[i]
		}

		if i < len(lines2) {
			line2 = lines2[i]
		}

		if line1 != line2 {
			if line1 != "" {
				differences = append(differences, fmt.Sprintf("\033[32m+\t\t%v\033[0m", line1))
			}
			if line2 != "" {
				differences = append(differences, fmt.Sprintf("\033[31m-\t\t%v\033[0m", line2))
			}
		}
	}

	return differences
}
