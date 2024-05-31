package objects

import (
	"fmt"
	"os"
	"strings"

	"github.com/Amabeusz/vcs/pkg/common"
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

	AddToIndex(sha, filePath)
}

func AddToIndex(sha string, filePath string) {
	file.FindAndWrite(global.INDEX_FILE, []byte(sha+" "+filePath+"\n"))
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
	return file.ReadRoot(global.INDEX_FILE)
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

func GetIndexFiles() map[string]string {
	root := common.GetRootPath()
	content := strings.Fields(string(ReadIndex()))

	files := make(map[string]string, 0)

	for i := 0; i < len(content); i += 2 {
		files[root+"\\"+content[i+1]] = content[i]
	}

	return files
}
