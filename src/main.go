package main

import (
	"fmt"
	"os"

	"github.com/ief2i-florent/go-f22/src/fs"
)

func main() {
	currentDir, _ := os.Getwd()
	fmt.Println(currentDir)
	fmt.Println("source", currentDir+"/test/sources")
	fmt.Println("target", currentDir+"/test/targets")

	err := fs.Copy(currentDir+"/test/sources", currentDir+"/test/targets")
	fmt.Println(err)
}
