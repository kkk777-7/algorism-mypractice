package main

import "fmt"

type Heap struct {
	heap []int
}

func main() {
	var h Heap

	h.push(5)
	h.push(3)
	h.push(7)
	h.push(1)

	fmt.Println(h.top(), h)

	h.pop()
	fmt.Println(h.top(), h)

	h.push(11)
	fmt.Println(h.top(), h)
}

func (h *Heap) push(n int) {
	h.heap = append(h.heap, n)
	idx := len(h.heap) - 1
	for idx > 0 {
		if n <= h.heap[(idx-1)/2] {
			break
		}
		h.heap[idx] = h.heap[(idx-1)/2]
		idx = (idx - 1) / 2
	}
	h.heap[idx] = n
}

func (h *Heap) top() int {
	if len(h.heap) == 0 {
		return -1
	}
	return h.heap[0]
}

func (h *Heap) pop() {
	if len(h.heap) == 0 {
		return
	}
	n := h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	i := 0
	for 2*i+1 < len(h.heap) {
		child1, child2 := 2*i+1, 2*i+2
		if h.heap[child1] < h.heap[child2] && child2 < len(h.heap) {
			child1 = child2
		}
		if n >= h.heap[child1] {
			break
		}
		h.heap[i] = h.heap[child1]
		i = child1
	}
	h.heap[i] = n
}
