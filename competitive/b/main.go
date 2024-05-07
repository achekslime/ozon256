package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var n, t int
	_, _ = fmt.Fscan(in, &n, &t)

	var key = ""

	for i := 0; i < n; i++ {
		var ni string
		_, _ = fmt.Fscan(in, &ni)
		key += ni
	}

	keyAsArray := strings.Split(key, "")
	sort.Slice(keyAsArray, func(i, j int) bool {
		return keyAsArray[i] < keyAsArray[j]
	})
	keyResult := strings.Join(keyAsArray, "")

	for i := 0; i < t; i++ {
		var ti string
		_, _ = fmt.Fscan(in, &ti)

		tiAsArray := strings.Split(ti, "")
		sort.Slice(tiAsArray, func(i, j int) bool {
			//a, _ := strconv.Atoi(tiAsArray[i])
			//b, _ := strconv.Atoi(tiAsArray[j])
			return tiAsArray[i] < tiAsArray[j]
		})
		tiResult := strings.Join(tiAsArray, "")

		if tiResult != keyResult {
			_, _ = fmt.Fprintln(out, "NO")
		} else {
			_, _ = fmt.Fprintln(out, "YES")
		}
	}
}
