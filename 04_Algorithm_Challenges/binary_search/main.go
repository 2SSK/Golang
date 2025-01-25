package main

func binarySearch(arr []int, target int) int {

	for l, r := 0, len(arr)-1; l <= r; {
		m := l + (r-l)/2
		if arr[m] == target {
			return m
		} else if arr[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}

func main() {}
