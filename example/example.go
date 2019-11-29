package main

import (
	"fmt"

	"github.com/ppreeper/pad"
)

func main() {
	fmt.Println(pad.Right("Hello", "!", 20))
	fmt.Println(pad.Left("Exit now", "→", 20))
	fmt.Println("-----")
	fmt.Println(pad.Left("D", "a", 4))
	fmt.Println(pad.Right("D", "a", 4))
	fmt.Println(pad.ZFill("D", 4))
	fmt.Println(pad.LJust("D", 4))
	fmt.Println(pad.RJust("D", 4))
	fmt.Println("-----")
	fmt.Println(pad.Left("ABCD", "a", 5))
	fmt.Println(pad.Right("ABCD", "a", 5))
	fmt.Println(pad.ZFill("ABCD", 5))
	fmt.Println(pad.LJust("ABCD", 5))
	fmt.Println(pad.RJust("ABCD", 5))
	fmt.Println("-----")
	fmt.Println(pad.LeftLen("D", "a", 4))
	fmt.Println(pad.RightLen("D", "a", 4))
	fmt.Println(pad.ZFillLen("D", 4))
	fmt.Println(pad.LJustLen("D", 4))
	fmt.Println(pad.RJustLen("D", 4))
	fmt.Println("-----")
	fmt.Println(pad.LeftLen("ABCD", "a", 5))
	fmt.Println(pad.RightLen("ABCD", "a", 5))
	fmt.Println(pad.ZFillLen("ABCD", 5))
	fmt.Println(pad.LJustLen("ABCD", 5))
	fmt.Println(pad.RJustLen("ABCD", 5))
}
