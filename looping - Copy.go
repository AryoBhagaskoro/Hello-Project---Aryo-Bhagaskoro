package main

import "fmt"

func dectobin(num int) {
	var biner []int

	for num != 0 {
		biner = append(biner, num%2)
		num = num / 2

	}
	if len(biner) == 0 {
		fmt.Print("0")

	} else {
		for i := len(biner) - 1; i > 0; i-- {
			fmt.Printf("%d", biner[i])
		}
	}

}
func main() {
	var desimal int
	fmt.Scan(&desimal)
	dectobin(desimal)
}
