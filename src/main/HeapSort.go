package main

import "fmt"

func main() {
	a := []int{2, 1, 10, 11, 9, 12, 6, 4, 3, 8, 5, 7, 0}
	fmt.Println(a)
	heapSort(a, len(a))
	fmt.Println(a)
}

func heapSort(a []int, N int) {
	for i := N / 2; i >= 0; i-- { // 实现堆有序
		sink(a, i, N)
	}
	fmt.Println(a)

	for i := N - 1; i >= 0; i-- { // 排序
		exch(a, 0, i)
		sink(a, 0, i)
	}
}

func sink(a []int, i, N int) { //将小的数字下沉
	for i*2+1 < N {
		j := i*2 + 1
		if j < N-1 && a[j] < a[j+1] {
			j++
		}
		if a[i] > a[j] {
			break
		}
		exch(a, i, j)
		i = j
	}
}

func exch(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}
