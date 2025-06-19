package main

import (
	"fmt"
	"os"
)

func main() {
	//	fmt.Println("Hello, World!")

	path := "/"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("Error: '%s' path doesn't exist.\n", path)
		return
	} else if err != nil {
		fmt.Printf("Error occured while accessing the path '%s': '%v'\n", path, err)
	}

}
