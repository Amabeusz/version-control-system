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

	repoNew := []string{}
	repoModified := []string{}
	repoDeleted := []string{}
	indexNew := []string{}
	indexModified := []string{}
	indexDeleted := []string{}
	notChanged := []string{}

	for _, file := range files {
		processFile(file, headFiles, indexFiles, &repoNew, &repoModified, &indexNew, &indexModified, &notChanged)
	}

	processRemainingIndexFiles(indexFiles, headFiles, &repoDeleted, &indexDeleted)

	for headFile := range headFiles {
		repoDeleted = append(repoDeleted, headFile)
	}

	printStatus("Not staged", repoNew, repoModified, repoDeleted)
	printStatus("Staged", indexNew, indexModified, indexDeleted)
	printUnchangedFiles(notChanged)
}

func processFile(f string, headFiles, indexFiles map[string]string, repoNew, repoModified, indexNew, indexModified, notChanged *[]string) {
	indexValue, inIndex := indexFiles[f]
	headValue, inHead := headFiles[f]

	if inIndex {
		delete(indexFiles, f)
		if inHead {
			delete(headFiles, f)
			fileSha := common.FileSha(file.Read(f))

			if fileSha != indexValue {
				*repoModified = append(*repoModified, f)
			} else if fileSha != headValue {
				*indexModified = append(*indexModified, f)
			} else {
				*notChanged = append(*notChanged, f)
			}
		} else {
			if common.FileSha(file.Read(f)) != indexValue {
				*repoNew = append(*repoNew, f)
			} else {
				*indexNew = append(*indexNew, f)
			}
		}
	} else if inHead {
		delete(headFiles, f)

		if common.FileSha(file.Read(f)) != headValue {
			*repoModified = append(*repoModified, f)
		} else {
			*notChanged = append(*notChanged, f)
		}
	} else {
		*repoNew = append(*repoNew, f)
	}
}

func processRemainingIndexFiles(indexFiles, headFiles map[string]string, repoDeleted, indexDeleted *[]string) {
	for file, indexValue := range indexFiles {
		if indexValue == "-" {
			*indexDeleted = append(*indexDeleted, file)
		} else {
			*repoDeleted = append(*repoDeleted, file)
		}
		delete(headFiles, file)
	}
}

func printStatus(header string, newFiles, modifiedFiles, deletedFiles []string) {
	fmt.Println(header)
	fmt.Println("\tNew files:")
	for _, file := range newFiles {
		fmt.Printf("\033[32m\t\t%s\033[0m\n", file)
	}
	fmt.Println("\tModified files:")
	for _, file := range modifiedFiles {
		fmt.Printf("\033[33m\t\t%s\033[0m\n", file)
	}
	fmt.Println("\tDeleted files:")
	for _, file := range deletedFiles {
		fmt.Printf("\033[31m\t\t%s\033[0m\n", file)
	}
}

func printUnchangedFiles(files []string) {
	fmt.Println("\nNot changed:")
	for _, file := range files {
		fmt.Printf("\t%s\n", file)
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
