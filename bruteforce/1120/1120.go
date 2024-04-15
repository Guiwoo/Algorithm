package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var (
		reader = bufio.NewReader(os.Stdin)
		writer = bufio.NewWriter(os.Stdout)
		x, y   string
		answer = math.MaxInt
		cal    = func(idx int) int {
			var rs int
			for i := idx; i < idx+len(x); i++ {
				if x[i-idx] != y[i] {
					rs++
				}
			}
			return rs
		}
	)
	defer writer.Flush()

	fmt.Fscanln(reader, &x, &y)

	for i := 0; i < len(y)-len(x)+1; i++ {
		answer = min(answer, cal(i))
		if answer == 0 {
			fmt.Fprintln(writer, "0")
			return
		}
	}

	fmt.Fprintln(writer, answer)
}
