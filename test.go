package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	test := bufio.NewReader(os.Stdin)
	tt, _ := test.ReadString('\n')
	fmt.Print(tt)
}
