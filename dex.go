package main

import (
	"fmt"
	"os"
)

func PrintDirRecussive(p string, level int) {

	entries, err := os.ReadDir(p)
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		tabSpace := ""
		for i := 0; i < level; i++ {
			tabSpace += "-"
		}
		info, err := entry.Info()
		if err != nil {
			panic(err)
		}
		if entry.IsDir() {
			fmt.Printf("%s Directory %s %v\n", tabSpace, entry.Name(), info.Mode())
			PrintDirRecussive(p+"/"+entry.Name(), level+1)
		} else {
			fmt.Printf("%s File %s %v\n", tabSpace, entry.Name(), info.Mode())
		}
	}
}

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("2 Args needed")
	}
	path := args[1]
	level := 0

	PrintDirRecussive(path, level)
}
