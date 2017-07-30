package main

import (
	"fmt"
	"math/rand"
	"time"

	dutch "github.com/valentinarho/go-dutch-flag/dutch"
)

// get an integer and return -1 if is even, 1 if is odd and 0 if is zero
func evenOddPartitioning(elem int) int {
	if elem == 0 {
		return 0
	} else if elem%2 == 0 {
		return -1
	} else {
		return 1
	}
}

func main() {
	// first example, using a 1-2-3 array and the basic partitioning function
	fa := []int{2, 1, 2, 3, 2, 3, 1, 1, 3, 1, 3, 2, 2, 2}
	fmt.Println("First array before partitioning:")
	fmt.Println(fa)
	dutch.DutchFlag(fa, dutch.BasicOneTwoThreeGetPartition)
	fmt.Println("First array after partitioning:")
	fmt.Println(fa)

	// second example, using a random int array and a custom partitioning function
	// that partition the array as <even, zeros, odd>
	rand.Seed(time.Now().Unix())
	var sa []int = make([]int, 10)
	sa[0] = 0
	for i := 1; i < 10; i++ {
		sa[i] = rand.Intn(1000)
	}

	fmt.Println("Second array before partitioning:")
	fmt.Println(sa)
	dutch.DutchFlag(sa, evenOddPartitioning)
	fmt.Println("Second array after partitioning:")
	fmt.Println(sa)
}
