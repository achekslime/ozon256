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

const (
	sendMsg = iota + 1
	printMsg
)

func main() {
	in, out := initBuffers()
	defer out.Flush()

	var n, q int
	_, _ = fmt.Fscan(in, &n, &q)
	userMessages := make([]int, n)
	messageNumber := 1
	globalMessage := 0

	for i := 0; i < q; i++ {
		var operationType, userID int
		_, _ = fmt.Fscan(in, &operationType, &userID)

		if operationType == sendMsg {
			if userID != 0 {
				userMessages[userID-1] = messageNumber
			} else {
				globalMessage = messageNumber
			}
			messageNumber += 1
		} else if operationType == printMsg {
			if userMessages[userID-1] < globalMessage {
				_, _ = fmt.Fprintln(out, globalMessage)
			} else {
				_, _ = fmt.Fprintln(out, userMessages[userID-1])
			}
		}
	}
}
