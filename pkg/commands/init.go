package commands

import (
	"fmt"
	"os"

	"github.com/Amabeusz/vcs/pkg/common"
	"github.com/Amabeusz/vcs/pkg/file"
	"github.com/Amabeusz/vcs/pkg/global"
)

func Init() {
	// create struct
	// .vcs
	// -- objects
	// -- refs/master
	// -- INDEX
	// -- HEAD [master]

	rootPath := common.GetRootPath()

	mkdir(rootPath + "\\.vcs")
	file.Create(rootPath+"\\"+global.INDEX_FILE, nil)
	file.Create(rootPath+"\\"+global.HEAD_FILE, []byte("master"))

	mkdir(rootPath + "\\" + global.VCS_PATH + "refs")
	file.Create(rootPath+"\\"+global.REFS_PATH+"master", nil)

	mkdir(rootPath + "\\" + global.VCS_PATH + "objects")
}

func mkdir(dirPath string) {
	err := os.Mkdir(dirPath, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	fmt.Println("Directory " + dirPath + " created successfully.")
}
