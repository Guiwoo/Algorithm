package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		writer = bufio.NewWriter(os.Stdout)
		oneK   = 1000
		tenK   = oneK * 10
		cal    func(i, j, total int) int
		arr    = make([]int, 0, 9000)
	)

	cal = func(i, j, total int) int {
		if i < j {
			return total + i
		}
		total += i % j
		return cal(i/j, j, total)
	}
	defer writer.Flush()

	for i := oneK; i < tenK; i++ {
		decimal := cal(i, 10, 0)
		duo := cal(i, 12, 0)
		if decimal != duo {
			continue
		}
		hexa := cal(i, 16, 0)
		if duo != hexa {
			continue
		}

		arr = append(arr, i)
	}

	for _, _num := range arr {
		fmt.Fprintln(writer, _num)
	}
}
