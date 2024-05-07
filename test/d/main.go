package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func initBuffers() (*bufio.Reader, *bufio.Writer) {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	return in, out
}

type Pair struct {
	Value    int
	Sequence int
}

func (p Pair) String() string {
	return fmt.Sprintf("%d: %d", p.Sequence, p.Value)
}

type ByTime []Pair

func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool { return a[i].Value < a[j].Value }

type BySequence []Pair

func (a BySequence) Len() int           { return len(a) }
func (a BySequence) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySequence) Less(i, j int) bool { return a[i].Sequence < a[j].Sequence }

func main() {
	in, out := initBuffers()
	defer out.Flush()

	var t int
	_, _ = fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Fscan(in, &n)
		times := make([]Pair, n)
		var tmpTime int
		for j := 0; j < n; j++ {
			_, _ = fmt.Fscan(in, &tmpTime)
			times[j].Value = tmpTime
			times[j].Sequence = j
		}

		sort.Sort(ByTime(times))

		currentPrizePlace := 1
		numberOfDividers := 1
		previousTime := times[0].Value
		prizes := make([]Pair, n)
		prizes[0].Value = currentPrizePlace
		prizes[0].Sequence = times[0].Sequence
		for j := 1; j < len(times); j++ {
			if times[j].Value-previousTime <= 1 {
				prizes[j].Value = currentPrizePlace
				prizes[j].Sequence = times[j].Sequence
				numberOfDividers++
			} else {
				currentPrizePlace = currentPrizePlace + numberOfDividers
				prizes[j].Value = currentPrizePlace
				prizes[j].Sequence = times[j].Sequence
				numberOfDividers = 1
			}
			previousTime = times[j].Value
		}

		sort.Sort(BySequence(prizes))

		for _, v := range prizes {
			_, _ = fmt.Fprint(out, fmt.Sprintf("%d ", v.Value))
		}
		_, _ = fmt.Fprintln(out, "")
	}
}
