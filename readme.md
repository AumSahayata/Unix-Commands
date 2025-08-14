# Unix Commands
A collection of Unix-inspired command-line tools implemented in Go

- `wc` - Counts the lines, words, and characters in a file.
- `grep` - Searches for specified text or patterns within files.
- `gotree` - Displays all files and directories in a given path in a tree-like format.
  
## Setup
1. Clone the repository
```bash
 git clone https://github.com/AumSahayata/Unix-Commands.git
```
2. Navigate to the repository
```bash
cd Unix-Commands
```
3. Run a tool – you can either build & install it, or run it directly.

- Option 1: Build & Install
```bash
cd wc
go build .
go install
```
After installation, you can run the command from anywhere:
```bash
wc test.txt
```
- Option 2: Run Directly
``` bash
cd wc
go run main.go <arguments>
```

## Command Flags (gotree)

`gotree` supports the following flags to customize its output:

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--fullPath` | `-f` | `false` | Print full paths of directories and files. |
| `--printDir` | `-d` | `false` | Show only directories (no files). |
| `--printPerm` | `-p` | `false` | Display file and directory permissions. |
| `--sort` | `-t` | `false` | Sort results by modification time. |
| `--levels` | `-L` | `-1` | Limit traversal depth to the specified number of levels (`-1` means no limit). |

### Example
```bash
# Print only directories with full paths, limited to 2 levels deep from test directory
gotree -d -f -L 2 test
```
## Command Flags (grep)
`grep` supports the following flags to customize its search behavior:

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--case_insensitive` | `-i` | `false` | Perform a case-insensitive search. |
| `--output` | `-o` | `""` | Write matching lines to the specified output file instead of stdout. |
| `--nlines_after` | `-A` | `0` | Print *n* lines after each match. |
| `--nlines_before` | `-B` | `0` | Print *n* lines before each match. |
| `--count` | `-C` | `false` | Only print the number of matches, not the matching lines. |

### Example
```bash
# Case-insensitive search, printing 2 lines after each match
grep -i -A 2 "error" logfile.txt

# Search and save results to a file
grep -o results.txt "TODO" source.go
```
## Command Flags (wc)

`wc` supports the following flags to control what is counted:

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--lines` | `-l` | `false` | Count the number of lines in the file. |
| `--words` | `-w` | `false` | Count the number of words in the file. |
| `--char` | `-c` | `false` | Count the number of characters in the file. |

### Example
```bash
# Count lines in a file
wc -l file.txt

# Count lines, words, and characters
wc -l -w -c file.txt
