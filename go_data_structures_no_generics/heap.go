package go_data_structures_no_generics

type MaxHeap struct {
	items []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) Peek() int {
	if len(h.items) == 0 {
		return 0 // Assuming 0 as the default "zero" value for int
	}
	return h.items[0]
}

func (h *MaxHeap) RemoveTop() {
	if len(h.items) == 0 {
		return
	}
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.siftDown(0)
}

func (h *MaxHeap) Insert(val int) {
	h.items = append(h.items, val)
	h.siftUp(len(h.items) - 1)
}

func (h *MaxHeap) Remove(val int) {
	for i, item := range h.items {
		if item == val {
			h.items[i] = h.items[len(h.items)-1]
			h.items = h.items[:len(h.items)-1]
			h.siftDown(i)
			return
		}
	}
}

func (h *MaxHeap) Empty() bool {
	return len(h.items) == 0
}

func (h *MaxHeap) Size() int {
	return len(h.items)
}

func (h *MaxHeap) Sorted() []int {
	originalItems := make([]int, len(h.items))
	copy(originalItems, h.items)

	var sorted []int
	for len(h.items) > 0 {
		sorted = append(sorted, h.Peek())
		h.RemoveTop()
	}

	h.items = originalItems
	return sorted
}

func (h *MaxHeap) siftUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.items[index] <= h.items[parentIndex] {
			break
		}
		h.items[index], h.items[parentIndex] = h.items[parentIndex], h.items[index]
		index = parentIndex
	}
}

func (h *MaxHeap) siftDown(index int) {
	lastIndex := len(h.items) - 1
	for index < lastIndex {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		largestIndex := index

		if leftChildIndex <= lastIndex && h.items[leftChildIndex] > h.items[largestIndex] {
			largestIndex = leftChildIndex
		}
		if rightChildIndex <= lastIndex && h.items[rightChildIndex] > h.items[largestIndex] {
			largestIndex = rightChildIndex
		}
		if largestIndex == index {
			break
		}
		h.items[index], h.items[largestIndex] = h.items[largestIndex], h.items[index]
		index = largestIndex
	}
}
