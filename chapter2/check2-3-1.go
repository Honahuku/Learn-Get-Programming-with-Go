package main

import "fmt"

func main() {
	const spaceXSpeed = 100800 //km/s
	var distance = 96300000    //km

	fmt.Println(distance/spaceXSpeed/60, "時間")
}
