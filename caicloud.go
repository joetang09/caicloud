package caicloud

func solution(bigArray []int) int {

	var (
		l = len(bigArray)
		i = 0
	)

	if l > 0 {
		for i++; i < l; i++ {
			if bigArray[i] > bigArray[i-1] {
				break
			}
		}
	}

	return i - 1
}
