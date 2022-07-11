package main

import "fmt"

// mainは最初に呼び出される関数
func main() {
	fmt.Print("火星の表面で、Nathanの体重は")
	fmt.Print(149.0 * 0.3783) // 56.3667
	fmt.Print("ポンド、年齢は、")
	fmt.Print(41 * 365 / 687) //21
	fmt.Print("歳になるでしょう。\n\n")

	fmt.Print("火星の表面で、ほなふくの体重は")
	fmt.Print(122.0 * 0.3783) // 46.1526
	fmt.Print("ポンド、年齢は、")
	fmt.Print(19 * 365 / 687) // 10
	fmt.Print("歳になるでしょう。")
}
