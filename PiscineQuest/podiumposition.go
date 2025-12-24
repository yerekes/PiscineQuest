package piscine

func PodiumPosition(podium [][]string) [][]string {
	left := 0
	right := len(podium) - 1

	for left < right {
		podium[left], podium[right] = podium[right], podium[left]
		left++
		right--
	}

	return podium
}
