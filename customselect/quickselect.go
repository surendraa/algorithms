package customselect

func QuickSelect(arr []int, nth int) int {
	n := nth - 1
	if n >= len(arr) {
		return -1
	}
	quickselect(0, len(arr)-1, &arr, n)
	return arr[n]
}

func quickselect(low int, high int, arr *[]int, n int) {
	if low < high {
		partition := partition(low, high, arr)
		if partition == n {
			return
		} else if partition > n {
			quickselect(low, partition-1, arr, n)
		} else {
			quickselect(partition+1, high, arr, n)
		}
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
