package echo

import (
	"fmt"
	"os"
	"strings"
)

func hello() {
	fmt.Println("hello")
}

func echo() {
	s, sep := "", ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "

	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func echo4() {
	fmt.Println(strings.Join(os.Args[1:], "\n"))
}
