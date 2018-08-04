package main

import (
	"fmt"
)

func sub(x int) {
	fmt.Println("share by arguments: ", x*x)
}

func main() {
	c := 5
	go sub(c)

	go func() {
		fmt.Println("share by capture: ", c*c)
	}()
	// キャプチャによって渡すことはポインタを引数で渡した場合と同じ
	// ==
	// go func(c *int) {
	// 	fmt.Println("share by capture: ", c*c)
	// }(&c)

	// deferはブロッキング的に動作する
	defer func() {
		fmt.Println("share by capture(defer): ", c*c)
	}()
	c = 10
}
