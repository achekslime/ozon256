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

	var t int
	_, _ = fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		_, _ = fmt.Fprintln(out, solve(in))
	}
}

const (
	bankCount      = 3
	operationCount = 6
)

const (
	rToD = iota
	rToE
	dToR
	dToE
	eToR
	eToD
)

type dpStr struct {
	dp           []float64
	usage        [][]bool
	defaultValue float64
	defaultUsage []bool
	costMatrix   [][]float64
}

func NewDp(defaultValue float64, costMatrix [][]float64) *dpStr {
	dp := &dpStr{}
	dp.dp = make([]float64, bankCount)
	dp.usage = make([][]bool, bankCount)
	for i := range dp.usage {
		dp.usage[i] = make([]bool, bankCount)
	}
	dp.costMatrix = costMatrix
	dp.defaultValue = defaultValue
	dp.defaultUsage = make([]bool, bankCount)
	return dp
}

func (dp dpStr) GetI(i int) float64 {
	if i < 0 {
		return dp.defaultValue
	}
	return dp.dp[i]
}

func (dp dpStr) GetUsage(i int) []bool {
	if i < 0 {
		return dp.defaultUsage
	}
	return dp.usage[i]
}

func (dp dpStr) Next(i int, aDp *dpStr, aOperation int, bDp *dpStr, bOperation int) {
	fromAValue, fromAIndex := aDp.findBestOperation(i-1, aOperation)
	fromBValue, fromBIndex := bDp.findBestOperation(i-1, bOperation)
	fromSelf := dp.GetI(i - 1)

	if fromAValue > fromBValue {
		if fromAValue > fromSelf {
			dp.dp[i] = fromAValue
			copy(dp.usage[i], aDp.GetUsage(i-1))
			dp.usage[i][fromAIndex] = true
			return
		}
	}
	if fromBValue > fromAValue {
		if fromBValue > fromSelf {
			dp.dp[i] = fromBValue
			copy(dp.usage[i], bDp.GetUsage(i-1))
			dp.usage[i][fromBIndex] = true
			return
		}
	}
	dp.dp[i] = dp.GetI(i - 1)
	dp.usage[i] = dp.GetUsage(i - 1)

}

func (dp dpStr) findBestOperation(i, operation int) (float64, int) {
	bestValue := float64(0)
	bestI := -1
	for j := 0; j < bankCount; j++ {
		if dp.GetUsage(i)[j] != true {
			if dp.GetI(i)*dp.costMatrix[j][operation] > bestValue {
				bestValue = dp.GetI(i) * dp.costMatrix[j][operation]
				bestI = j
			}
		}
	}
	return bestValue, bestI
}

func solve(in *bufio.Reader) float64 {
	costMatrix := make([][]float64, bankCount)
	for j := 0; j < bankCount; j++ {
		costMatrix[j] = make([]float64, operationCount)
	}

	for i := 0; i < bankCount; i++ {
		for j := 0; j < operationCount; j++ {
			var a, b float64
			_, _ = fmt.Fscan(in, &a, &b)
			costMatrix[i][j] = b / a
		}
	}

	dpD := NewDp(0, costMatrix)
	dpE := NewDp(0, costMatrix)
	dpR := NewDp(1, costMatrix)

	for i := 0; i < bankCount; i++ {
		dpR.Next(i, dpD, dToR, dpE, eToR)
		dpE.Next(i, dpD, dToE, dpR, rToE)
		dpD.Next(i, dpE, eToD, dpR, rToD)
	}

	return dpD.dp[bankCount-1]
}
