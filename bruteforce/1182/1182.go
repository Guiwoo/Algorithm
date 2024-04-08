package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		reader                        = bufio.NewReader(os.Stdin)
		writer                        = bufio.NewWriter(os.Stdout)
		length, target, total, answer int
		recur                         func(cur, idx, flag int)
	)
	defer writer.Flush()

	fmt.Fscanln(reader, &length, &target)
	arr := make([]int, length)

	for i := 0; i < length; i++ {
		var x int
		fmt.Fscan(reader, &x)
		arr[i] = x
		total += x
	}
	recur = func(cur, idx, flag int) {
		if idx >= length {
			if cur == target && flag > 0 {
				answer++
			}
			return
		}
		recur(cur+arr[idx], idx+1, flag|(1<<idx))
		recur(cur, idx+1, flag)
	}

	recur(0, 0, 0)
	fmt.Fprintln(writer, answer)
}
