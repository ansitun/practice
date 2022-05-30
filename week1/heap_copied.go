package main

import (
	"container/heap"
	"sort"
)

type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

type IntHeap struct {
	sort.IntSlice
}

func (h *IntHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	x := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[0 : h.Len()-1]
	return x
}

type MaxHeap struct {
	IntHeap
}

func (h *MaxHeap) Less(i, j int) bool {
	return h.IntHeap.IntSlice[i] > h.IntHeap.IntSlice[j]
}

func (h *IntHeap) Peak() int {
	return h.IntSlice[0]
}

type MedianFinder struct {
	min *IntHeap
	max *MaxHeap
}

func Constructor() MedianFinder {
	min := &IntHeap{}
	max := &MaxHeap{}
	return MedianFinder{min, max}
}

func (this *MedianFinder) AddNum(num int) {

	if this.max.Len() == 0 {
		heap.Push(this.max, num)
		return
	}

	if this.max.Peak() > num {
		heap.Push(this.max, num)
	} else {
		heap.Push(this.min, num)
	}

	d := this.max.Len() - this.min.Len()
	if d > 1 {
		x := heap.Pop(this.max)
		heap.Push(this.min, x)
	} else if d < -1 {
		x := heap.Pop(this.min)
		heap.Push(this.max, x)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	d := this.max.Len() - this.min.Len()
	switch d {
	case 1:
		return float64(this.max.Peak())
	case -1:
		return float64(this.min.Peak())
	}
	return float64(this.max.Peak()+this.min.Peak()) / 2.0
}

/** My Code

type MinHeap []int

type MaxHeap []int

type MedianFinder struct {
   minHeap *MinHeap
   maxHeap *MaxHeap
}


func Constructor() MedianFinder {
   minHeap := &MinHeap{}
   maxHeap := &MaxHeap{}

   return MedianFinder {minHeap,maxHeap}
}

// ----Min Heap---
func (minHeap *MinHeap)heapify(heap *MinHeap, i int) {
   smallest := i
   lChild := 2*i+1
   rChild := 2*i+2

   if lChild < len(*heap) && (*heap)[lChild] < (*heap)[smallest] {
       smallest = lChild
   }


   if rChild < len(*heap) && (*heap)[rChild] < (*heap)[smallest] {
       smallest = rChild
   }

   if i != smallest {
       (*heap)[smallest], (*heap)[i] = (*heap)[i], (*heap)[smallest]
       minHeap.heapify(heap, smallest)
   }
}

func (minHeap *MinHeap) push(heap *MinHeap, num int) {
   *heap = append(*heap, num)
   // after appending the number we need to heapify the entire array to reset the max element correctly
   for i:= len(*heap)/2 - 1; i >=0; i-- {
       minHeap.heapify(heap, i)
   }

   return
}

func (minHeap *MinHeap) pop(heap *MinHeap) int {
   // take the root element - min in this case
   root := (*heap)[0]
   n := len(*heap)
   lastElement := (*heap)[n-1]
   // since the root element will be popped out, replace the 0th element with the last element of heap
   (*heap)[0] = lastElement
   // resize the heap to ignore the last element
   *heap = (*heap)[:n-1]
   // heapify to find the correct placemnt of the 0th element
   minHeap.heapify(heap, 0)

   return root
}

// ----- Max Heap-----

func (maxHeap *MaxHeap)heapify(heap *MaxHeap, i int) {
   largest := i
   lChild := 2*i+1
   rChild := 2*i+2

   if lChild < len(*heap) && (*heap)[lChild] > (*heap)[largest] {
       largest = lChild
   }


   if rChild < len(*heap) && (*heap)[rChild] > (*heap)[largest] {
       largest = rChild
   }

   if i != largest {
       (*heap)[largest], (*heap)[i] = (*heap)[i], (*heap)[largest]
       maxHeap.heapify(heap, largest)
   }
}

// the below func is same for as min heap
func (maxHeap *MaxHeap) push(heap *MaxHeap, num int) {
   *heap = append(*heap, num)
   // after appending the number we need to heapify the entire array to reset the max element correctly
   for i:= len(*heap)/2 - 1; i >=0; i-- {
       maxHeap.heapify(heap, i)
   }

   return
}

// the below func is same for as min heap
func (maxHeap *MaxHeap) pop(heap *MaxHeap) int {
  // take the root element - max in this case
   root := (*heap)[0]
   n := len(*heap)
   lastElement := (*heap)[n-1]
   // since the root element will be popped out, replace the 0th element with the last element of heap
   (*heap)[0] = lastElement
   // resize the heap to ignore the last element
   *heap = (*heap)[:n-1]
   // heapify to find the correct placemnt of the 0th element
   maxHeap.heapify(heap, 0)

   return root
}

func checkBalance(minHeap *MinHeap, maxHeap *MaxHeap) int {
   return len(*maxHeap) - len(*minHeap)
}

func (this *MedianFinder) AddNum(num int)  {
   // add to max heap
   // check balance between heaps
   // rebalance the max and min heap
   this.maxHeap.push(this.maxHeap, num)

   balance := checkBalance(this.minHeap, this.maxHeap)
   if  balance > 1 {
       this.minHeap.push(this.minHeap, this.maxHeap.pop(this.maxHeap))
   } else if balance < 0 {
       this.maxHeap.push(this.maxHeap, this.minHeap.pop(this.minHeap))
   } else {
       if len(*this.minHeap) > 0 && (*this.maxHeap)[0] > (*this.minHeap)[0] {
           // swap the top of min and max heaps
           topMinHeap := (*this.minHeap)[0]
           topMaxHeap := (*this.maxHeap)[0]
           (*this.minHeap)[0] = topMaxHeap
           (*this.maxHeap)[0] = topMinHeap
       }
	}
}

func (this *MedianFinder) FindMedian() float64 {
   median := 0.0

   if len(*this.maxHeap) > len(*this.minHeap) {
       median = float64((*this.maxHeap)[0])
   } else if len(*this.maxHeap) < len(*this.minHeap) {
       median = float64((*this.minHeap)[0])
   } else {
       median = float64((*this.maxHeap)[0] + (*this.minHeap)[0])/2.0
   }

   return median
}

*/
