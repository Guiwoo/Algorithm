package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var (
		reader = bufio.NewReader(os.Stdin)
		writer = bufio.NewWriter(os.Stdout)
	)
	defer writer.Flush()

	arr := make([]int, 9)
	total := 0

	for i := 0; i < len(arr); i++ {
		var x int
		fmt.Fscanln(reader, &x)
		arr[i] = x
		total += x
	}

	sort.Ints(arr)

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if total-arr[i]-arr[j] == 100 {
				fmt.Println(i, j)
				for k := 0; k < len(arr); k++ {
					if k == i || k == j {
						continue
					}
					fmt.Println(arr[k])
				}
				return
			}
		}
	}
}
