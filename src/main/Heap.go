package main

/**
 * 二叉堆
 */
type Heap struct {
	a []int
	N int // 数组a的大小
}

func (heap Heap) exch(i, j int) { // 交换a[i]和a[j]
	heap.a[i], heap.a[j] = heap.a[j], heap.a[i]
}

func (heap Heap) swim(i int) { // 上浮
	for i > 1 && heap.a[i] > heap.a[i/2] {
		heap.exch(i, i/2)
		i /= 2
	}
}

func (heap Heap) sink(i int) { // 下潜
	for i*2 <= heap.N {
		j := 2 * i
		if j < heap.N && heap.a[j] < heap.a[j+1] {
			j++
		}
		if heap.a[i] < heap.a[j] {
			break
		}
		heap.exch(i, j)
		i = j
	}
}

func (heap Heap) getTop() int {
	return heap.a[0]
}

func (heap Heap) delTop() int {
	top := heap.a[0]
	heap.a[0] = heap.a[heap.N]
	heap.N--
	heap.sink(0)
	return top
}

func (heap Heap) add(value int) {
	heap.a = append(heap.a, value)
	heap.N++
	heap.sink(heap.N)
}
