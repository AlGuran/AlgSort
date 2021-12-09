package BubbleSort

func BubbleSortInt(slice []int) []int {
	swap := false
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			slice[i], slice[i+1] = slice[i+1], slice[i]
			swap = true
		}
	}

	if swap {
		return BubbleSortInt(slice)
	}
	return slice
}
