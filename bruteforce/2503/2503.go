package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
4
123 1 1
356 1 0
327 2 0
489 0 1
*/

type NumBaseBall struct {
	Number       string
	Strike, Ball int
}

func main() {
	var (
		reader     = bufio.NewReader(os.Stdin)
		writer     = bufio.NewWriter(os.Stdout)
		N, answer  int
		isPossible = func(arr []NumBaseBall, nums []int) bool {
			for _, v := range arr {
				strike := 0
				ball := 0
				for idx, vv := range v.Number {
					for numIdx, num := range nums {
						target := int(vv-'0') == num
						if target && idx == numIdx {
							strike++
							break
						}
						if target {
							ball++
							break
						}
					}
				}
				if v.Strike != strike || v.Ball != ball {
					return false
				}
			}
			return true
		}
	)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	arr := make([]NumBaseBall, N)
	for i := 0; i < N; i++ {
		var (
			number       string
			strike, ball int
		)
		fmt.Fscanln(reader, &number, &strike, &ball)
		arr[i] = NumBaseBall{
			Number: number,
			Strike: strike,
			Ball:   ball,
		}
	}

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			for k := 1; k <= 9; k++ {
				if i == j || j == k || i == k {
					continue
				}
				if isPossible(arr, []int{i, j, k}) {
					answer++
				}
			}
		}
	}

	fmt.Fprintln(writer, answer)
}
