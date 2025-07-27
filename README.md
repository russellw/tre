# tre

A command line utility that displays directory structure with simplified output, similar to the Windows `tree` command but designed for programmatic use.

## Features

- Clean, indented output with no tree characters - just file and directory names
- Respects `.gitignore` files in target directory and parent directories
- Automatically skips `.git` directories
- Optional `-f` flag to show directories only
- Takes directory path as argument (defaults to current directory)

## Usage

```bash
# Show all files and directories in current directory
./tre

# Show all files and directories in specified directory  
./tre /path/to/directory

# Show only directories (with -f flag)
./tre -f

# Show only directories in specified path
./tre -f /path/to/directory
```

## Example Output

```
myproject
  src
    main.go
    utils.go
  docs
    README.md
  .gitignore
  go.mod
```

With `-f` flag (directories only):
```
myproject
  src
  docs
```

## Building

```bash
go build -o tre
```
