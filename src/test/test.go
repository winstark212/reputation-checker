package main
import (
    "fmt"
		"math/rand"
		"time"
)


func main() {
	start := time.Now()
	var n = 100000

	for num := 1; num < n; num++ {
		printNumber(num)
	}

	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s", elapsed)
}

func printNumber(num int) {
	fmt.Print(rand.Intn(100))
}