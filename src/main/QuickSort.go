package main

/**
 * 快速排序
 */
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var a []int
	for i := 0; i < 10; i++ {
		a = append(a, rand.Intn(100))
	}
	fmt.Println(a)
	quicoSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func quicoSort(a []int, lo, hi int) { // 递归调用
	if lo >= hi {
		return
	}
	j := partition(a, lo, hi)
	quicoSort(a, lo, j-1)
	quicoSort(a, j+1, hi)
}

func partition(a []int, lo, hi int) int { // 将a[lo+1]至a[hi]中所有小于a[lo]的数放左边, 大雨a[lo]的数放右边
	i, j := lo+1, hi
	v := a[lo]
	for true {
		for a[i] < v {
			if i == hi {
				break
			}
			i++
		}
		for a[j] > v {
			if j == lo {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	a[j], a[lo] = a[lo], a[j]
	return j
}
