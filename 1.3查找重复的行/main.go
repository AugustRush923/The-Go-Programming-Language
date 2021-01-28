package main

import (
	"bufio"
	"fmt"
	"os"
)

func repeatLine() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		fmt.Println(input.Text())
		counts[input.Text()]++
		// 等价于
		// line := input.Text()
		// counts[line] = counts[line] + 1  ==> counts[input.Text()] = counts[input.Text()] + 1
		if input.Text() == "q" {
			break
		}
	}
	fmt.Println(counts)
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func repeatFile() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if input.Text() == "q" {
			break
		}
	}
}

func main() {
	repeatFile()
}
