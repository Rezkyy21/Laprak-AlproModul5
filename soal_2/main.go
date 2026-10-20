package main

import "fmt"

func main() {
	var saklar1, saklar2, saklar3 int
	var saklar bool

	fmt.Scan(&saklar1, &saklar2, &saklar3)
	saklar = (saklar1 == 1 && saklar2 == 1) || (saklar1 == 1 && saklar3 == 1) || (saklar2 == 1 && saklar3 == 1)

	fmt.Print(saklar)
}
