package main

// #include <stdio.h>
import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {

	fmt.Println(s)

	const n = 500_000_000
	fmt.Println(n)

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
