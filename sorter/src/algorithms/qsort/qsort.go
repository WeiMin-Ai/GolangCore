package qsort

func QuickSort(values []int) {
	length := len(values)
	baseValue := values[0]
	var leftIndex int
	var rightIndex int = length - 1
	var flag bool = true

	for {
		if flag {
			for j := rightIndex; j > 0; j-- {
				if baseValue > values[j] {
					values[leftIndex] = values[j]
					rightIndex = j
					leftIndex++
					flag = false
					break
				}
			}
		} else {
			for x := leftIndex; x < length; x++ {
				if baseValue < values[x] {
					values[rightIndex] = values[x]
					leftIndex = x
					rightIndex--
					flag = true
					break
				}
			}
		}

		if leftIndex == rightIndex {
			values[leftIndex] = baseValue
			break
		}
	}

}
