package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type GitIgnore struct {
	patterns []string
}

func NewGitIgnore(dir string) *GitIgnore {
	gi := &GitIgnore{patterns: []string{}}
	
	current := dir
	for {
		gitignorePath := filepath.Join(current, ".gitignore")
		if file, err := os.Open(gitignorePath); err == nil {
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if line != "" && !strings.HasPrefix(line, "#") {
					gi.patterns = append(gi.patterns, line)
				}
			}
			file.Close()
		}
		
		parent := filepath.Dir(current)
		if parent == current {
			break
		}
		current = parent
	}
	
	return gi
}

func (gi *GitIgnore) IsIgnored(path string, isDir bool) bool {
	for _, pattern := range gi.patterns {
		if matched, _ := filepath.Match(pattern, filepath.Base(path)); matched {
			return true
		}
		if strings.HasSuffix(pattern, "/") && isDir {
			if matched, _ := filepath.Match(strings.TrimSuffix(pattern, "/"), filepath.Base(path)); matched {
				return true
			}
		}
	}
	return false
}

func printTree(dir string, indent string, dirsOnly bool, gitignore *GitIgnore) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	var filteredEntries []os.DirEntry
	for _, entry := range entries {
		if entry.Name() == ".git" {
			continue
		}
		path := filepath.Join(dir, entry.Name())
		if !gitignore.IsIgnored(path, entry.IsDir()) {
			if !dirsOnly || entry.IsDir() {
				filteredEntries = append(filteredEntries, entry)
			}
		}
	}

	sort.Slice(filteredEntries, func(i, j int) bool {
		return filteredEntries[i].Name() < filteredEntries[j].Name()
	})

	for _, entry := range filteredEntries {
		fmt.Printf("%s%s\n", indent, entry.Name())
		
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			printTree(subDir, indent+"  ", dirsOnly, gitignore)
		}
	}

	return nil
}

func main() {
	var dirsOnly bool
	flag.BoolVar(&dirsOnly, "f", false, "Show directories only")
	flag.Parse()

	targetDir := "."
	if flag.NArg() > 0 {
		targetDir = flag.Arg(0)
	}

	absDir, err := filepath.Abs(targetDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error resolving path: %v\n", err)
		os.Exit(1)
	}

	if _, err := os.Stat(absDir); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Directory does not exist: %s\n", absDir)
		os.Exit(1)
	}

	gitignore := NewGitIgnore(absDir)
	
	fmt.Println(filepath.Base(absDir))
	if err := printTree(absDir, "  ", dirsOnly, gitignore); err != nil {
		fmt.Fprintf(os.Stderr, "Error traversing directory: %v\n", err)
		os.Exit(1)
	}
}