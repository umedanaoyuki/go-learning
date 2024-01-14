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
	number, _ := strconv.Atoi(sc.Text())
	fmt.Println(number)

	// 条件文
	if name == "go" {
		fmt.Println("welcome")
	}

	if number <= 100 {
		// 整数を文字列と連結するには、strconv.Itoa()で変換してから、+で連結します。
		fmt.Println(strconv.Itoa(number) + "は10に等しい")
	}

}
