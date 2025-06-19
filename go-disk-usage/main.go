package main

import (
	"fmt"
	"os"
	"syscall"
)

func getDiskUsage(path string) {

	var stat syscall.Statfs_t
	err := syscall.Statfs(path, &stat)
	if err != nil {
		fmt.Printf("Error fetching disk usage for '%s': '%v'\n", path, err)
	}

	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bfree * uint64(stat.Bsize)
	used := total - free

	percentUsed := (float64(used) / float64(total)) * 100

	fmt.Printf("Disk usage of path '%s'\n", path)
	fmt.Printf("Total: %d GB\n", total/1e9)
	fmt.Printf("Used: %d GB (%.2f%%)\n", used/1e9, percentUsed)
	fmt.Printf("Free: `%d` GB\n", free/1e9)
}

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

	getDiskUsage(path)

}
