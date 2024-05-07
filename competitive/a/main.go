package main

import (
	"bufio"
	"fmt"
	"os"
)

func initBuffers() (*bufio.Reader, *bufio.Writer) {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	return in, out
}

func main() {
	in, out := initBuffers()
	defer out.Flush()

	var a, b int
	_, _ = fmt.Fscan(in, &a, &b)
	_, _ = fmt.Fprintln(out, a-b)
}
