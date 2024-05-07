package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	var n int
	_, _ = fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var s string
		_, _ = fmt.Fscan(in, &s)
		sAsSlice := strings.Split(s, "")

		key := sAsSlice[0]
		if sAsSlice[len(sAsSlice)-1] != key {
			_, _ = fmt.Fprintln(out, "NO")
			continue
		}

		ok := true
		for j := 1; j < len(sAsSlice)-1; j++ {
			if sAsSlice[j] != key {
				if sAsSlice[j-1] != key || sAsSlice[j+1] != key {
					_, _ = fmt.Fprintln(out, "NO")
					ok = false
					break
				}
			}
		}

		if ok {
			_, _ = fmt.Fprintln(out, "YES")
		}
	}
}
