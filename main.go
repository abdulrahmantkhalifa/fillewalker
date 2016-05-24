package main

import (
	"fmt"
	"os"

	"github.com/abdulrahmantkhalifa/fillewalker/walker"
)

func main() {
	err := walker.Walk("test/", printFile)
	if err != nil {
		fmt.Println(err)
	}
}

func printFile(path string, file os.FileInfo) error {
	fmt.Println("this is : ", file.Name())
	return nil
}
