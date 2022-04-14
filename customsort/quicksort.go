package customsort

//Quicksort: Given an int array use quicksort and return sorted array
func QuickSort(arr []int) []int {
	low := 0
	high := len(arr) - 1
	quick(low, high, &arr)
	return arr
}

func quick(low, high int, arr *[]int) {
	if low < high {
		pivot_index := partition(low, high, arr)
		quick(low, pivot_index-1, arr)
		quick(pivot_index+1, high, arr)
	}
}

func partition(low, high int, arr *[]int) int {
	pivot := (*arr)[high]
	i := low
	j := high - 1

	for {
		if i >= j {
			break
		}

		if (*arr)[i] < pivot {
			i++
		}

		if (*arr)[j] > pivot {
			j--
		}

		if (*arr)[i] > (*arr)[j] {
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		}
	}

	if (*arr)[i] > pivot {
		(*arr)[i], (*arr)[high] = (*arr)[high], (*arr)[i]
	}

	return i

}
