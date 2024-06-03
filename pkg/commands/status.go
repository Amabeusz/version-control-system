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

	//	fmt.Println("head: ")
	//	for k := range headFiles {
	//		fmt.Println(k)
	//	}
	//
	//	fmt.Println("index: ")
	//	for k := range indexFiles {
	//		fmt.Println(k)
	//	}

	repoNew := make([]string, 0)
	repoModified := make([]string, 0)
	repoDeleted := make([]string, 0)
	indexNew := make([]string, 0)
	indexModified := make([]string, 0)
	indexDeleted := make([]string, 0)

	notChanged := make([]string, 0)

	for _, f := range files {
		fmt.Println(f)
		if indexValue, ok := indexFiles[f]; ok {
			fmt.Println("FOUND IN INDEX")
			delete(indexFiles, f)
			if headValue, ok := headFiles[f]; ok {
				fmt.Println("FOUND IN HEAD")
				delete(headFiles, f)
				fileSha := common.FileSha(file.Read(f))
				fmt.Println("fileSHa " + fileSha)
				fmt.Println("indexValue " + indexValue)
				fmt.Println("headValue " + headValue)

				if fileSha != indexValue {
					repoModified = append(repoModified, f)
				} else {
					if fileSha != headValue {
						indexModified = append(indexModified, f)
					} else {
						notChanged = append(notChanged, f)
					}
				}
			} else {
				if common.FileSha(file.Read(f)) != indexValue {
					repoNew = append(repoNew, f)
				} else {
					indexNew = append(indexNew, f)
				}
			}
		} else {
			if headValue, ok := headFiles[f]; ok {
				delete(headFiles, f)

				fmt.Println(f)
				fmt.Println(common.FileSha(file.Read(f)))
				fmt.Println(headValue)
				if common.FileSha(file.Read(f)) != headValue {
					repoModified = append(repoModified, f)
				} else {
					notChanged = append(notChanged, f)
				}
			} else {
				repoNew = append(repoNew, f)
			}
		}
	}

	for k, v := range indexFiles {
		if v == "-" {
			indexDeleted = append(indexDeleted, k)
			delete(headFiles, k)
		} else {
			repoDeleted = append(repoDeleted, k)
			delete(headFiles, k)
		}
	}

	for k := range headFiles {
		repoDeleted = append(repoDeleted, k)
	}

	fmt.Println("Not staged")

	fmt.Println("\tNew files:")
	for _, v := range repoNew {
		fmt.Println("\033[32m\t\t" + v + "\033[0m")
	}

	fmt.Println("\tModified files:")
	for _, v := range repoModified {
		fmt.Println("\033[33m\t\t" + v + "\033[0m")
	}

	fmt.Println("\tDeleted files:")
	for _, v := range repoDeleted {
		fmt.Println("\033[31m\t\t" + v + "\033[0m")
	}

	fmt.Println("Staged")

	fmt.Println("\tNew files:")
	for _, v := range indexNew {
		fmt.Println("\033[32m\t\t" + v + "\033[0m")
	}

	fmt.Println("\tModified files:")
	for _, v := range indexModified {
		fmt.Println("\033[33m\t\t" + v + "\033[0m")
	}

	fmt.Println("\tDeleted files:")
	for _, v := range indexDeleted {
		fmt.Println("\033[31m\t\t" + v + "\033[0m")
	}

	fmt.Println("\nNot changed:")
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
