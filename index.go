package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// 複数データを受け取って、分類する
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	count, _ := strconv.Atoi(sc.Text())
	fmt.Println(count)

	for i := 0; i < count; i++ {
		sc.Scan()
		number, _ := strconv.Atoi(sc.Text())
		fmt.Println(number)

		if number == 10 {
			fmt.Println(strconv.Itoa(number) + "は10に等しい")
		} else if number > 10 {
			fmt.Println(strconv.Itoa(number) + "は10より大きい")
		} else {
			fmt.Println(strconv.Itoa(number) + "は10未満")
		}
	}


}
