package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		reader        = bufio.NewReader(os.Stdin)
		writer        = bufio.NewWriter(os.Stdout)
		length, total int
		fnc           func(arr []int, out []int, depth, flag int)
		sum           = func(arr []int) int {
			var x int
			for i := 1; i < len(arr); i++ {
				t := arr[i] - arr[i-1]
				if t < 0 {
					t *= -1
				}
				x += t
			}
			return x
		}
	)
	defer writer.Flush()

	fnc = func(arr []int, out []int, depth, flag int) {
		if depth >= len(arr) {
			x := sum(out)
			if total < x {
				fmt.Println(out)
				total = x
			}
			return
		}
		for i := 0; i < len(arr); i++ {
			if flag&(1<<i) != 0 {
				continue
			}
			out[depth] = arr[i]
			fnc(arr, out, depth+1, flag|(1<<i))
		}
	}

	fmt.Fscanln(reader, &length)

	arr := make([]int, 0, length)

	for i := 0; i < length; i++ {
		var x int
		fmt.Fscan(reader, &x)
		arr = append(arr, x)
	}

	out := make([]int, length)
	fnc(arr, out, 0, 0)
	fmt.Fprintln(writer, total)
}
