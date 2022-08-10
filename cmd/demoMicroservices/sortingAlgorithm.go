package demoMicroservices

// sorting Algorithms

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
			j--
		}
	}
	return arr
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	result = append(result, left...)
	result = append(result, right...)
	return result
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[0]
	left, right := partition(arr[1:], pivot)
	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func partition(arr []int, pivot int) ([]int, []int) {
	var left, right []int
	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return left, right
}

func heapSort(arr []int) []int {
	for i := len(arr) / 2; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
	return arr
}

func heapify(arr []int, i, n int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2
	if l < n && arr[l] > arr[largest] {
		largest = l
	}
	if r < n && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, largest, n)
	}
}

func countingSort(arr []int) []int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	count := make([]int, max+1)
	for _, v := range arr {
		count[v]++
	}
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	result := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		result[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}
	return result
}

func radixSort(arr []int) []int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	buckets := make([][]int, 10)
	for i := 0; i < 10; i++ {
		buckets[i] = make([]int, 0)
	}
	for _, v := range arr {
		buckets[v%10] = append(buckets[v%10], v)
	}
	result := make([]int, 0)
	for _, v := range buckets {
		result = append(result, radixSort(v)...)
	}
	return result
}

// what time complexity is each sorting algorithm?
// bubble sort: O(n^2)
// selection sort: O(n^2)
// insertion sort: O(n^2)
// merge sort: O(n log n)
// quick sort: O(n log n)
// heap sort: O(n log n)
// counting sort: O(n)
// radix sort: O(n)
