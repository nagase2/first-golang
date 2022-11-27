package main

import "fmt"

func main() {

	// numbers := []int32{5, 10, 10}
	// queries := [][]int32{{1, 2, 5}}
	// result := findSum(numbers, queries)
	// fmt.Println(result)

	// numbers := []int32{-5, 0, 2}
	// queries := [][]int32{{2, 2, 20}, {1, 2, 10}}
	// result := findSum(numbers, queries)
	// fmt.Println(result)

	sides := [][]int64{{4, 8}, {15, 30}, {25, 50}}
	fmt.Println("result1:", nearlySimilarRectangles(sides))

	sides2 := [][]int64{{2, 1}, {10, 7}, {9, 6}, {6, 9}, {7, 3}}
	fmt.Println("result2:", nearlySimilarRectangles(sides2))
}

func findSum(numbers []int32, queries [][]int32) []int64 {
	// Write your code here
	resultArray := []int64{}

	fmt.Print(numbers, " ", queries)
	for i := 0; i < len(queries); i++ {
		fmt.Println("query ", i)
		start := queries[i][0] - 1
		end := queries[i][1] - 1
		zeroValue := queries[i][2]

		result := int64(0)
		for k := start; k <= end; k++ {
			targetNum := numbers[k]
			if targetNum == 0 {
				result += int64(zeroValue)
			} else {
				result += int64(targetNum)
			}
		}
		resultArray = append(resultArray, result)
	}

	return resultArray

}

func nearlySimilarRectangles(sides [][]int64) int64 {
	count := int64(0)
	for i := 0; i < len(sides)-1; i++ {
		for k := i + 1; k < len(sides); k++ {
			if float32(sides[i][0])/float32(sides[k][0]) == float32(sides[i][1])/float32(sides[k][1]) {
				count++
			}
		}
	}
	return int64(count)
}
