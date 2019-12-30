package dup

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func dup() {
	// dup function count the line that appear more than once in the standard input,preceded by its count
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	fmt.Println("result:\n")
	for line, n := range counts { //line 为键值， n为出现次数
		if n > 1 {
			fmt.Printf("键值:%s\t出现次数：%d\n", line, n)
		}
	}
}

/*****************Dividing line*****************/

func dup2() {
	//dup2 print the count and text of lines that appear more than once in the input.
	// It read from stdin or from list of named files.
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				_, _ = fmt.Fprint(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			_ = f.Close()
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
	}
}

/*****************Dividing line*****************/

func dup3() {

	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			_, _ = fmt.Fprint(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n\t%s", n, line)
		}
	}
}
