package main

import "fmt"

func main() {
	a := []int{2, 1, 5, 7, 9, 0, 6, 4, 3, 8}
	fmt.Println(a)
	sort(a, 0, 9)
	fmt.Println(a)
}

func sort(a []int, lo, hi int) {
	if (lo >= hi) {
		return
	}
	j := partition(a, lo, hi)
	sort(a, lo, j-1)
	sort(a, j+1, hi)
}

func partition(a []int, lo, hi int) (int) {
	i, j := lo+1, hi
	v := a[lo]
	for true {
		for a[i] < v {
			if (i == hi) {
				break
			}
			i++
		}
		for a[j] > v {
			if (j == lo) {
				break
			}
			j--
		}
		if (i >= j) {
			break
		}
		tmp := a[i]
		a[i] = a[j]
		a[j] = tmp
	}
	tmp := a[lo]
	a[lo] = a[j]
	a[j] = tmp
	return j
}
