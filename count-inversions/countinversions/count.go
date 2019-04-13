package countinversions

//BruteForce is a function to count the inversions using the first-thought way.
func BruteForce(slice []int) int {
	length := len(slice)
	count := 0
	for i := 0; i < length; i++ {
		num := slice[i]
		for j := i + 1; j < length; j++ {
			if slice[j] < num {
				count++
			}
		}
	}
	return count
}

//Count is a function to count the inversions using divid and conquer approach.
func Count(slice []int) int {
	length := len(slice)

	if length < 2 {
		return 0
	}
	mid := length / 2
	left := slice[0:mid]
	right := slice[mid:length]

	return Count(left) + Count(right) + mergeAndCount(slice, left, right)
}

func mergeAndCount(dest []int, left []int, right []int) int {
	length := len(dest)
	lenLeft := len(left)
	lenRight := len(right)

	tmp := make([]int, length)
	result := 0
	i := 0
	j := 0
	k := 0
	for ; i < lenLeft && j < lenRight; k++ {
		if left[i] <= right[j] {
			tmp[k] = left[i]
			i++
		} else {
			tmp[k] = right[j]
			result += lenLeft - i
			j++
		}
	}
	for ; i < lenLeft; i, k = i+1, k+1 {
		tmp[k] = left[i]
	}
	for ; j < lenRight; j, k = j+1, k+1 {
		tmp[k] = right[j]
	}

	copy(dest, tmp)
	return result
}
