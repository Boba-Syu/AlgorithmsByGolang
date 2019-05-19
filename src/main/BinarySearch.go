package main

/**
 * 二分查找
 */
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(time.Now().Unix())
	fmt.Println(a)
	n := rand.Intn(len(a))
	fmt.Printf("数组%d为数组a中的第%d项", n, binarySearch(a, n)+1)
}

func binarySearch(a []int, n int) int {
	lo, hi := 0, len(a)-1
	i := (lo + hi) / 2
	for lo < hi {
		if a[i] == n {
			return i
		}
		if a[i] < n {
			lo = i
			i = (lo + hi) / 2
		} else {
			hi = i
			i = (lo + hi) / 2
		}
	}
	return -1
}
