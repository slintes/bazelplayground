package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Hello World, listing files in /:")

	err := filepath.Walk("/", func(path string, info os.FileInfo, err error) error {
		if info != nil && info.IsDir() && strings.HasPrefix(path, "/test") {
			fmt.Println(path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
