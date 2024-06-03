package objects

import "github.com/Amabeusz/vcs/pkg/file"

func SaveBlob(filePath string) string {
	return SaveObject(file.ReadRoot(filePath))
	// return SaveObject(append([]byte("blob\n"), file.ReadRoot(filePath)...))
}
