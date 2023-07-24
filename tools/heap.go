package tools

import (
	"container/heap"
	"github.com/dadongc/aggregation/dto"
)

type Float64Heap []dto.CosineLabel

func (f *Float64Heap) Len() int {
	return len(*f)
}

func (f *Float64Heap) Less(i, j int) bool {
	return (*f)[i].Cosine < (*f)[j].Cosine
}

// Swap swaps the elements with indexes i and j.
func (f *Float64Heap) Swap(i, j int) {
	(*f)[i], (*f)[j] = (*f)[j], (*f)[i]
}

func (f *Float64Heap) Push(x any) {
	*f = append(*f, x.(dto.CosineLabel))
}

func (f *Float64Heap) Pop() any {
	tmp := (*f)[len(*f)-1]
	*f = (*f)[:len(*f)-1]
	return tmp
}

func GetMaxNumsOfCosine(arr []dto.CosineLabel, k int) []dto.CosineLabel {
	h := &Float64Heap{}
	heap.Init(h)
	for _, a := range arr {
		heap.Push(h, a)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return *h
}
