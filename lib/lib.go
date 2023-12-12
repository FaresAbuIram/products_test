package lib

func DeleteItemByValueFromSlice(slice []int, value int) []int {
	index := -1

	// Find the index of the element with the specified value
	for i, item := range slice {
		if item == value {
			index = i
			break
		}
	}

	// If the value was found, remove the element at that index
	if index != -1 {
		return append(slice[:index], slice[index+1:]...)
	}

	// Value not found, return the original slice unchanged
	return slice
}
