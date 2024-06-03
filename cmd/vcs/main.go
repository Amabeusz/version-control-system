package main

import (
	"fmt"
	"os"

	"github.com/Amabeusz/vcs/pkg/commands"
)

func main() {
	// initCommand := flag.NewFlagSet("init", flag.ExitOnError)
	// statusCommand := flag.NewFlagSet("status", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("usage: vcs <command> [<args>]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		commands.Init()
	case "status":
		commands.PrintStatus()
	//	case "help":
	//		fmt.Println("help [status | init]")
	case "add":
		checkArgs(3)
		commands.Add(os.Args[2])
	case "commit":
		checkArgs(3)
		commands.Commit(os.Args[2])
	case "cat-file":
		checkArgs(3)
		if len(os.Args) == 3 {
			commands.CatObject(os.Args[2])
		}
		if len(os.Args) == 4 {
			commands.CatObjectType(os.Args[2], os.Args[3])
		}
	default:
		os.Exit(1)
	}
}

func checkArgs(n int) {
	if len(os.Args) < n {
		fmt.Println("Specify file")
		os.Exit(1)
	}
}
