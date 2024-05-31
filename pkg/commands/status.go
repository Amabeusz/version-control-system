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

func PrintStatus() {
	files := getAllFiles(common.GetRootPath())
	headFiles := objects.GetHeadFiles()
	indexFiles := objects.GetIndexFiles()

	fmt.Println("head: ")
	for k := range headFiles {
		fmt.Println(k)
	}

	fmt.Println("index: ")
	for k := range indexFiles {
		fmt.Println(k)
	}

	repoNew := make([]string, 0)
	repoUpdated := make([]string, 0)
	repoDeleted := make([]string, 0)
	indexNew := make([]string, 0)
	indexUpdated := make([]string, 0)
	indexDeleted := make([]string, 0)

	notChanged := make([]string, 0)

	for _, f := range files {
		fmt.Println("file: " + f)
		if indexValue, ok := indexFiles[f]; ok {
			// staged
			delete(indexFiles, f)
			if _, ok := headFiles[f]; ok {
				delete(headFiles, f)

				if common.FileSha(file.Read(f)) != indexValue {
					indexUpdated = append(indexUpdated, f)
				} else {
					notChanged = append(notChanged, f)
				}
			} else {
				if common.FileSha(file.Read(f)) != indexValue {
					repoUpdated = append(repoUpdated, f)
				} else {
					indexNew = append(indexNew, f)
				}
			}
		} else {
			// repo
			if headValue, ok := headFiles[f]; ok {
				delete(headFiles, f)

				if common.FileSha(file.Read(f)) != headValue {
					repoUpdated = append(repoUpdated, f)
				} else {
					notChanged = append(notChanged, f)
				}
			} else {
				repoNew = append(repoNew, f)
			}
		}
	}

	for k := range indexFiles {
		indexDeleted = append(indexDeleted, k)
		delete(headFiles, k)
	}

	for k := range headFiles {
		repoDeleted = append(repoDeleted, k)
	}

	fmt.Println("Not staged")

	fmt.Println("\tNew files:")
	for _, v := range repoNew {
		fmt.Println("\t\t" + v)
	}

	fmt.Println("\tUpdated files:")
	for _, v := range repoUpdated {
		fmt.Println("\t\t" + v)
	}

	fmt.Println("\tDeleted files:")
	for _, v := range repoDeleted {
		fmt.Println("\t\t" + v)
	}

	fmt.Println("Staged")

	fmt.Println("\tNew files:")
	for _, v := range indexNew {
		fmt.Println("\t\t" + v)
	}

	fmt.Println("\tUpdated files:")
	for _, v := range indexUpdated {
		fmt.Println("\t\t" + v)
	}

	fmt.Println("\tDeleted files:")
	for _, v := range indexDeleted {
		fmt.Println("\t\t" + v)
	}

	fmt.Println("\nNot chenged:")
	for _, v := range notChanged {
		fmt.Println("\t" + v)
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
			return nil
		}

		files = append(files, path)
		return nil
	})
	common.Check(err)

	return files
}
