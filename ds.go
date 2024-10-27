package main

import (
	"flag"
	"fmt"
	ep "github.com/PeterHickman/expand_path"
	humanize "github.com/dustin/go-humanize"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
)

var raw bool
var swap bool
var args []string

func fileArgument() string {
	if len(args) != 1 {
		fmt.Println("No file argument given")
		os.Exit(1)
	}

	return args[0]
}

func init() {
	r := flag.Bool("raw", false, "Report sizes as bytes")
	s := flag.Bool("swap", false, "Display size before directory name")

	flag.Parse()

	raw = *r
	swap = *s

	args = flag.Args()
}

func main() {
	sizes := map[string]int64{}

	root, _ := ep.ExpandPath(fileArgument())

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if info == nil {
			return nil
		}

		fh, err := os.Stat(path)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if !fh.IsDir() {
			sizes[filepath.Dir(path)] += fh.Size()
		}

		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %v\n", err)
		return
	}

	keys := []string{}
	for k, _ := range sizes {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	for _, k := range keys {
		if raw {
			if swap {
				fmt.Printf("%d %s\n", sizes[k], k)
			} else {
				fmt.Printf("%s %d\n", k, sizes[k])
			}
		} else {
			if swap {
				fmt.Printf("%s %s\n", humanize.Bytes(uint64(sizes[k])), k)
			} else {
				fmt.Printf("%s %s\n", k, humanize.Bytes(uint64(sizes[k])))
			}
		}
	}
}
