package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
)

type DiskUsage struct {
	Path      string
	Total     float64
	Available float64
	Used      float64
	UsedPact  float64
}

func getDiskUsage(path string) (*DiskUsage, error) {
	var stat syscall.Statfs_t

	if err := syscall.Statfs(path, &stat); err != nil {
		return nil, fmt.Errorf("failed to stat read '%s': %v", path, err)
	}

	total := float64(stat.Blocks) * float64(stat.Bsize)
	if total == 0 {
		return nil, fmt.Errorf("Disk total size is zero for path '%s'", path)
	}

	free := float64(stat.Bavail) * float64(stat.Bsize)
	used := total - free
	percentUsed := (used / total) * 100

	return &DiskUsage{
		Path:      path,
		Total:     total,
		Available: free,
		Used:      used,
		UsedPact:  percentUsed,
	}, nil
}

func humanize(bytes float64) string {

	const (
		KB = 1 << 10 // 2^10 -> 1024 (right shift operator)
		MB = 1 << 20 // 2^20
		GB = 1 << 30 // 2^30
		TB = 1 << 40 // 2^40
	)

	switch {
	case bytes >= TB:
		return fmt.Sprintf("%.2f TB", bytes/TB)
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", bytes/GB)
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", bytes/MB)
	default:
		return fmt.Sprintf("%.2f KB", bytes)
	}
}

func printDiskUsage(du *DiskUsage) {
	fmt.Printf("Disk usage of the path '%s'\n", du.Path)
	fmt.Printf("Total: '%s'\n", humanize(du.Total))
	fmt.Printf("Used: '%s' (%.2f%%)\n", humanize(du.Used), du.UsedPact)
	fmt.Printf("Available: '%s'\n", humanize(du.Available))
}

func main() {
	path := flag.String("path", "/", "Path to check disk usage")
	flag.Parse()

	if _, err := os.Stat(*path); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error path '%s' does not exist\n", *path)
		os.Exit(1)
	}

	du, err := getDiskUsage(*path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	printDiskUsage(du)
}
