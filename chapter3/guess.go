package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var num = rand.Intn(100) + 1 //0を含めたくないため+1する

	var pNum = 50 // Predict Num, 予測値の意
	for {
		fmt.Printf("現在の予測値は%vです\n", pNum)
		if pNum < num {
			fmt.Println("numは予測値より大きい。処理を続行します。")
			pNum++
		} else if pNum > num {
			fmt.Println("numは予測値より小さい。処理を続行します。")
			pNum--
		} else if pNum == num {
			fmt.Println("numは%vでした。", pNum)
			break
		}
	}
}
