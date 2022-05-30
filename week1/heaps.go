package main

import (
	"fmt"
)

type MaxHeap []int
type MinHeap []int

// this is a min heap!
func main() {
	minHeap := make(MinHeap, 0)
	maxHeap := make(MaxHeap, 0)
	nums := make([]int, 0)
	input := make([]int, 0)
	for i := 1; i <= 50000; i++ {
		input = append(input, i)
	}
	// addition of an element to min and max heap for a data stream of medians
	for j := 0; j < 50000; j++ {
		// rand.Intn(20000)
		num := input[j]
		fmt.Printf("--> New Number %d\n", num)
		nums = append(nums, num)

		// by default push to maxHeap and then check the balance
		maxHeap.push(&maxHeap, num)
		// check balance
		balance := checkBalance(maxHeap, minHeap)
		if balance > 1 {
			minHeap.push(&minHeap, maxHeap.pop(&maxHeap))
		} else if balance < 0 {
			maxHeap.push(&maxHeap, minHeap.pop(&minHeap))
		} else {
			if len(minHeap) > 0 && (maxHeap)[0] > (minHeap)[0] {
				temp := minHeap[0]
				temp2 := maxHeap[0]
				minHeap[0] = temp2
				maxHeap[0] = temp
			}
		}

		//fmt.Print("MaxHeap ---> ")
		//for x := 0; x < len(maxHeap); x++ {
		//	fmt.Printf("%d ", maxHeap[x])
		//}
		//fmt.Println()
		//
		//fmt.Print("MinHeap ===> ")
		//for x := 0; x < len(minHeap); x++ {
		//	fmt.Printf("%d ", minHeap[x])
		//}
		//fmt.Println()

		median := 0.0
		if len(maxHeap) > len(minHeap) {
			median = float64((maxHeap)[0])
		} else if len(maxHeap) < len(minHeap) {
			median = float64((minHeap)[0])
		} else {
			median = float64((maxHeap)[0]+(minHeap)[0]) / 2.0
		}
		fmt.Printf(" For Index = %d The Median is ===> %f\n", j+1, median)
	}

	//fmt.Print("Input array ===> ")
	//for x := 0; x < len(nums); x++ {
	//	fmt.Printf("%d ", nums[x])
	//}
	//fmt.Println()
}

func checkBalance(maxHeap MaxHeap, minHeap MinHeap) int {
	return len(maxHeap) - len(minHeap)
}

// ---------Min Heap-------
func (minHeap *MinHeap) pop(heap *MinHeap) int {
	// smallest
	root := (*heap)[0]
	n := len(*heap)
	lastElement := (*heap)[n-1]
	(*heap)[0] = lastElement
	*heap = (*heap)[:n-1]
	minHeap.heapify(heap, 0)

	return root
}

func (minHeap *MinHeap) push(heap *MinHeap, num int) {
	// smallest
	*heap = append(*heap, num)
	//bottom up approach - start from the non-leaf nodes and traverse upwards
	for i := len(*heap)/2 - 1; i >= 0; i-- {
		minHeap.heapify(heap, i)
	}

	return
}

// pointer to array and the second param is the index we want to heapify
func (minHeap *MinHeap) heapify(heap *MinHeap, i int) {
	smallest := i
	lChild := 2*i + 1
	rChild := 2*i + 2

	if lChild < len(*heap) && (*heap)[lChild] < (*heap)[smallest] {
		smallest = lChild
	}

	if rChild < len(*heap) && (*heap)[rChild] < (*heap)[smallest] {
		smallest = rChild
	}

	if smallest != i {
		(*heap)[smallest], (*heap)[i] = (*heap)[i], (*heap)[smallest]
		minHeap.heapify(heap, smallest)
	}
}

// ---------Max Heap-------
func (maxHeap *MaxHeap) pop(heap *MaxHeap) int {
	// smallest
	root := (*heap)[0]
	n := len(*heap)
	lastElement := (*heap)[n-1]
	(*heap)[0] = lastElement
	*heap = (*heap)[:n-1]
	maxHeap.heapify(heap, 0)

	return root
}

func (maxHeap *MaxHeap) push(heap *MaxHeap, num int) {
	// smallest
	*heap = append(*heap, num)
	//bottom up approach - start from the non-leaf nodes and traverse upwards
	for i := len(*heap)/2 - 1; i >= 0; i-- {
		maxHeap.heapify(heap, i)
	}

	return
}

// pointer to array and the second param is the index we want to heapify
func (maxHeap *MaxHeap) heapify(heap *MaxHeap, i int) {
	largest := i
	lChild := 2*i + 1
	rChild := 2*i + 2

	if lChild < len(*heap) && (*heap)[lChild] > (*heap)[largest] {
		largest = lChild
	}

	if rChild < len(*heap) && (*heap)[rChild] > (*heap)[largest] {
		largest = rChild
	}

	if largest != i {
		(*heap)[largest], (*heap)[i] = (*heap)[i], (*heap)[largest]
		maxHeap.heapify(heap, largest)
	}
}
