package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDir(dir string, fileSizes chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			go walkDir(filepath.Join(dir, entry.Name()), fileSizes, wg)
		} else {
			info, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dir info: %v\n", err)
			}
			fileSizes <- info.Size()
		}
	}
}

var cache = make(chan struct{}, 20)

func dirents(dir string) []os.DirEntry {
	cache <- struct{}{}
	defer func() {
		<-cache
	}()
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
	}
	return entries
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		fmt.Println("true")
		return true
	default:
		return false
	}
}

func main() {
	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var wg sync.WaitGroup
	for _, dir := range roots {
		wg.Add(1)
		go walkDir(dir, fileSizes, &wg)
	}
	go func() {
		wg.Wait()
		close(fileSizes)
	}()
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		case <-done:
			// 此时还有 go程在后台执行，接收他们但是不处理（避免他们阻塞泄露）
			for range fileSizes {
			}
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
