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

	var s string
	var n int
	_, _ = fmt.Fscan(in, &s, &n)

	resultSticker := strings.Split(s, "")

	for i := 0; i < n; i++ {
		var r string
		var start, end int
		_, _ = fmt.Fscan(in, &start, &end, &r)
		rAsArray := strings.Split(r, "")
		for j := start; j <= end; j++ {
			resultSticker[j-1] = rAsArray[j-start]
		}
	}

	_, _ = fmt.Fprintln(out, strings.Join(resultSticker, ""))
}
