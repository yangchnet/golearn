package main

import "fmt"

type ByteCount int

func (c *ByteCount) Write(p []byte) (int, error) {
	*c += ByteCount(len(p))
	return len(p), nil
}

func main() {
	var c ByteCount
	_, _ = c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	var name = "Dolly"
	_, _ = fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}