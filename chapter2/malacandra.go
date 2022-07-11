package main

import (
	"fmt"
)

func main() {
	var distance = 56000000
	var days = 28
	var speed = distance / days
	fmt.Print("マラカンドラに", days, "日で到着するためには、", speed, "km/hで宇宙船を飛ばす必要があります。")
}
