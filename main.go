package main

import (
	ft "github.com/weaming/disk-analysis/filetree"
)

func main() {
	path := "."
	root := ft.NewFileNode(path, path, nil)
	ft.PrintFileNode(root, []string{}, 1, 2)
}
