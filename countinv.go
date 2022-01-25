package countinv

// SortCount parses an integer array and returns the array sorted and the number of inversions in the array.
func SortCount(arr []int) ([]int, int64) {
	arrLen := len(arr)

	if arrLen == 1 {

		// base case
		return arr, 0

	} else {

		// recursive calls
		m := arrLen / 2
		leftArr, leftInvCount := SortCount(arr[:m])
		rightArr, rightInvCount := SortCount(arr[m:])
		invCountSum := leftInvCount + rightInvCount

		// merge sort and count inversion(s)
		var (
			sortedArr []int
			i, j      = 0, 0
		)
		for i < len(leftArr) && j < len(rightArr) {
			if leftArr[i] <= rightArr[j] {
				sortedArr = append(sortedArr, leftArr[i])
				i++
			} else {
				sortedArr = append(sortedArr, rightArr[j])
				j++
				invCountSum += int64(len(leftArr) - i) // increment inversion count
			}
		}

		// merge leftover(s)
		for ; i < len(leftArr); i++ {
			sortedArr = append(sortedArr, leftArr[i])
		}
		for ; j < len(rightArr); j++ {
			sortedArr = append(sortedArr, rightArr[j])
		}

		return sortedArr, invCountSum
	}
}
