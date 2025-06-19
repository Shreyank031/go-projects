# Disk Usage CLI (Go)

A simple command-line tool written in Go to display disk usage statistics for a given file system path.

## 📦 Features

- Fetches and displays:
  - Total disk space
  - Used disk space
  - Available disk space
  - Percentage of disk used
- Human-readable byte formatting (e.g., MB, GB, TB)
- Error handling for invalid paths

## 🛠️ Installation

First, clone the repository:

```bash
git clone https://github.com/Shreyank031/go-projects.git
cd disk-usage-go
```
## Then build the binary:
```bash
go build -o diskusage main.go
```

## Usage
```bash
./diskusage
./diskusage --path=/
```

## Sample Output:

```bash
Disk usage of the path '/'
Total: '100.00 GB'
Used: '72.34 GB' (72.34%)
Available: '27.66 GB'
```
