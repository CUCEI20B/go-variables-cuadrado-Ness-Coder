package main

import "fmt"

func main()  {
	var lade uint64

	fmt.Scanln(&lade)

	area := lade * lade

	fmt.Println(area)
}