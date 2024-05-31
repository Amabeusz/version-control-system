package objects

import "github.com/Amabeusz/vcs/pkg/file"

func SaveBlob(filePath string) string {
	return SaveObject(append([]byte("blob\n"), file.ReadRoot(filePath)...))
}
