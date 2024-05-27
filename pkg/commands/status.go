package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Amabeusz/vcs/pkg/common"
	"github.com/Amabeusz/vcs/pkg/file"
	"github.com/Amabeusz/vcs/pkg/objects"
)

func Status() {
	files := getAllFiles(common.GetRootPath())

	indexFiles := getFiles()

	stagedFiles := make([]string, 0)

	for _, v := range files {

		if _, ok := indexFiles[v]; !ok {
			fmt.Printf("new %v\n", v)
			continue
		}

		// change to sha on last commit
		fileSha := common.FileSha(file.Read(lastElement(v)))

		// if string(indexFiles[v]) == fileSha {
		// dont show - not changed

		if string(indexFiles[v]) != fileSha {
			stagedFiles = append(stagedFiles, "updated "+v)
		}
	}

	fmt.Println("Staged")

	for _, v := range stagedFiles {
		fmt.Println(v)
	}
}

func lastElement(str string) string {
	if str == "" {
		log.Fatal("The string is empty")
	}

	parts := strings.Split(str, "\\")
	if len(parts) == 0 {
		log.Fatal("No parts found")
	}

	last := parts[len(parts)-1]
	return last
}

func getAllFiles(dir string) []string {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if info.Name() == ".git" || info.Name() == ".vcs" {
				return filepath.SkipDir
			}
		}

		files = append(files, path)
		return nil
	})
	common.Check(err)

	return files
}

func getFiles() map[string]string {
	root := common.GetRootPath()
	status := objects.ReadIndex()

	s := strings.Fields(string(status))

	files := make(map[string]string, 0)

	for i, v := range s {
		if i%2 == 0 {
			continue
		}

		files[root+"\\"+v] = s[i-1]
	}

	return files
}
