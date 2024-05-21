package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Amabeusz/vcs/pkg/common"
	"github.com/Amabeusz/vcs/pkg/objects"
)

func Status() {
	files := getAllFiles(common.GetRootPath())

	//for _, v := range files {
	//	fmt.Println(v)
	//}

	indexFiles := getFiles()

	fmt.Println(indexFiles)

	for _, v := range files {
		// fmt.Printf("Comparing %v to %v\n", v, indexFiles[v])
		if indexFiles[v] != 1 {
			fmt.Printf("New file %v\n", v)
			continue
		}

		// if string(indexFiles[v]) != common.FileSha(file.Read(v[14:])) {
		fmt.Printf("File changed %v\n", v)
		//}

	}
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

func getFiles() map[string]int {
	root := common.GetRootPath()
	status := objects.ReadIndex()

	s := strings.Fields(string(status))

	files := make(map[string]int, 0)

	for i, v := range s {
		if i%2 == 0 {
			continue
		}

		files[root+"\\"+v] = 1
	}

	return files
}
