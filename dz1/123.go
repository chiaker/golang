package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("go run . <temp>")
		os.Exit(1)
	}

	temp, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		os.Exit(1)
	}
	//F = C * 9/5 + 32
	fmt.Printf("%.1f\n", temp*9/5+32)
}
