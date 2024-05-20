package commands

import (
	"github.com/Amabeusz/vcs/pkg/objects"
)

func Commit(msg string) {
	objects.CreateCommit(msg)
}
