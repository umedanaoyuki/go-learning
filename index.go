package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 入力タブから文字列データを受け取る
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	name := sc.Text()
	fmt.Println("Hello" + name)
}
