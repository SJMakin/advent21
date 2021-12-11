package main

func sum(array [3]int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
