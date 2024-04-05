package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		reader  = bufio.NewReader(os.Stdin)
		writer  = bufio.NewWriter(os.Stdout)
		builder = strings.Builder{}
	)
	defer writer.Flush()

	for {
		inputs, _ := reader.ReadString('\n')
		input := strings.Split(strings.TrimSpace(inputs), " ")
		if input[0] == "0" {
			fmt.Println(builder.String())
			break
		}

		output := make([]string, 0, 6)
		permutation(&builder, input, output, 1, 0|(0<<1))
		builder.WriteString("\n")
	}
}

func permutation(build *strings.Builder, arr, out []string, start, flag int) {
	if len(out) >= 6 {
		ans := ""
		for _, v := range out {
			ans += v + " "
		}
		build.WriteString(ans + "\n")
		return
	} else {
		for i := start; i < len(arr); i++ {
			if (flag & (1 << i)) != 0 {
				continue
			}
			permutation(build, arr, append(out, arr[i]), i+1, flag|(1<<i))
		}
	}
}
