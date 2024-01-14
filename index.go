package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)

	// 入力タブから文字列データを受け取る
	sc.Scan()
	name := sc.Text()
	fmt.Println("Hello" + name)

	// 入力タブから数字を受け取る
	sc.Scan()
	count, _ := strconv.Atoi(sc.Text())
	fmt.Println(count)

	greeting := "Hello World"

	for i := 0; i < count; i++ {
		fmt.Println(greeting)
	}

}
