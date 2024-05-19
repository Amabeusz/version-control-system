package objects

import (
	"fmt"
	"os"
	"strings"

	"github.com/Amabeusz/vcs/pkg/file"
	"github.com/Amabeusz/vcs/pkg/global"
)

func AddBlobToIndex(sha string, filePath string) {
	indexSha := shaInIndex(filePath)

	if indexSha == sha {
		fmt.Println("Already in index")
		return
	}

	if indexSha != "" && indexSha != sha {
		ReplaceInIndex(sha, filePath)
		fmt.Println("Replaced " + indexSha + " with " + sha)
		return
	}

	AddToIndex([]byte(sha + " " + filePath + "\n"))
}

func AddToIndex(s []byte) {
	file.FindAndWrite(global.INDEX_FILE, s)
}

func RemoveFromIndex(filePath string) {
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func ReplaceInIndex(sha string, filePath string) {
	indexContent := ReadIndex()

	rows := strings.Split(string(indexContent), "\n")

	for i, row := range rows {
		if strings.Contains(row, filePath) {
			rows[i] = sha + " " + filePath
		}
	}

	replaced := strings.Join(rows, "\n")

	os.WriteFile(global.INDEX_FILE, []byte(replaced), 0644)
}

func ReadIndex() []byte {
	return file.Read(global.INDEX_FILE)
}

func shaInIndex(filePath string) string {
	indexContent := ReadIndex()

	rows := strings.Split(string(indexContent), "\n")

	mapp := make(map[string]string)

	for _, row := range rows {
		r := strings.Split(row, " ")
		if len(r) > 1 {
			mapp[r[1]] = r[0]
		}
	}

	return mapp[filePath]
}
