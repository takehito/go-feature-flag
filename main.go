package main

import (
	"flag"
	"fmt"
)

var feature bool

func sumNew(a int, b int) int {
	fmt.Println("fixing sum is not completed!")
	return a + b
}

func sum(a int, b int) int {
	if feature {
		return sumNew(a, b)
	}
	return a + b
}

func main() {
	flag.BoolVar(&feature, "feature", false, "開発中の機能を実行するためのフラグ")
	flag.Parse()

	fmt.Println(sum(2, 7))
}
