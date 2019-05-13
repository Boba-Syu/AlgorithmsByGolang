package main

/**
 * 归并排序
 */
import "fmt"

var aux [10]int

func main() {
	a := []int{2, 1, 5, 7, 9, 0, 6, 4, 3, 8}
	fmt.Println(a)
	//top_down_sort(a, 0, 9)
	down_top_sort(a)
	fmt.Println(a)
}

func down_top_sort(a []int) { // 自底向上(递推)
	len := len(a)
	for size := 1; size < len; size += size {
		for lo := 0; lo < len-size; lo += size + size {
			var hi int
			if lo+size+size < len-1 {
				hi = lo + size + size
			} else {
				hi = len - 1
			}
			merge(a, lo, lo+size-1, hi)
		}
	}
}

func top_down_sort(a []int, lo, hi int) { // 自顶向下(递归)
	if lo >= hi {
		return
	}
	mid := (lo + hi) / 2
	top_down_sort(a, lo, mid)
	top_down_sort(a, mid+1, hi)
	merge(a, lo, mid, hi)
}

func merge(a []int, lo, mid, hi int) { // 归并操作
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}
	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if aux[i] > aux[j] {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}
