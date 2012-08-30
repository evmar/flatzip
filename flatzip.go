package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: flatzip srcdir dstdir\n" +
			"flatzip creates a mtime-equivalent mirror of srcdir\n" +
			"with all files filled with zeros for easy compression.\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
		return
	}
	srcroot := flag.Arg(0)
	dstroot := flag.Arg(1)

	visit := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}

		if !strings.HasPrefix(path, srcroot) {
			log.Fatalf("expected path %q to start with src %q", path, srcroot)
		}
		path = path[len(srcroot):]
		dst := filepath.Join(dstroot, path)

		if info.IsDir() {
			log.Printf("%s\n", path)
			check(os.MkdirAll(dst, 0777))
		} else {
			f, err := os.Create(dst)
			check(err)
			check(f.Truncate(info.Size()))
			check(f.Close())
			check(os.Chtimes(dst, info.ModTime(), info.ModTime()))
		}
		return nil
	}
	check(filepath.Walk(srcroot, visit))
}
