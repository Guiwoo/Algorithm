package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		reader   = bufio.NewReader(os.Stdin)
		writer   = bufio.NewWriter(os.Stdout)
		colCheck = func(row, col int, str string, arr [][]string) bool {
			for i := 0; i < len(arr); i++ {
				if arr[i][col] == str {
					return false
				}
			}
			return true
		}
		rowCheck = func(row, col int, str string, arr [][]string) bool {
			for i := 0; i < len(arr); i++ {
				if arr[row][i] == str {
					return false
				}
			}
			return true
		}
		crossCheck = func(row, col int, str string, arr [][]string) bool {
			nRow := (row / 3) * 3
			nCol := (col / 3) * 3

			for i := nRow; i < nRow+3; i++ {
				for j := nCol; j < nCol+3; j++ {
					if arr[i][j] == str {
						return false
					}
				}
			}
			return true
		}
		recur = func(cur int) {}
	)
	defer writer.Flush()

	arr := make([][]string, 9)
	for i := range arr {
		x, _ := reader.ReadString('\n')
		arr[i] = strings.Split(strings.TrimSpace(x), " ")
	}

	recur = func(cur int) {
		if cur == 81 {
			for _, v := range arr {
				fmt.Printf("%+v\n", v)
			}
			os.Exit(1)
			return
		}
		row := cur / 9
		col := cur % 9

		if arr[row][col] == "0" {
			for _, v := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"} {
				if colCheck(row, col, v, arr) && rowCheck(row, col, v, arr) && crossCheck(row, col, v, arr) {
					arr[row][col] = v
					recur(cur + 1)
					arr[row][col] = "0"
				}
			}
		} else {
			recur(cur + 1)
		}
	}
	recur(0)
}
