package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 文字列を1行入力
func StrStdin() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput = scanner.Text()

	stringInput = strings.TrimSpace(stringInput)
	return
}

func main() {
	p := StrStdin()
	fmt.Println("結果 "+ p)
}
