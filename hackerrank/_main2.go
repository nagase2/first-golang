package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	numbers := []int{4, 1, 2, 5, 3}
	fmt.Printf("ソート前: %d\n", numbers) // ソート前: [4 1 2 5 3]
	sort.Ints(numbers)                // sort.Sort(sort.IntSlice(numbers))
	fmt.Printf("ソート後: %d\n", numbers) // ソート後: ソート後: [1 2 3 4 5]

	result := getMinCost([]int{1, 3, 5}, []int{3, 5, 7})

	fmt.Println("R1:", result)

	result = getMinCost([]int{5, 3, 1, 4, 6}, []int{9, 8, 3, 15, 1})

	fmt.Println("R2:", result)

	result = getMinCost([]int{5, 1, 4, 2}, []int{4, 7, 9, 10})

	fmt.Println("R3:", result)

	byteArray:=divideNumberToBytes(84344)
	fmt.Println("😸", byteArray)
	digit,_ := strconv.Atoi(string(byteArray[0]))
	fmt.Println("digit:",digit)
}

// func findSum(numbers []int32, queries [][]int32) []int64 {
// }

func getMinCost(crew_id []int, job_id []int) int {
	sort.Ints(crew_id)
	sort.Ints(job_id)

	//fmt.Print("job_id:", job_id)

	totalCost := 0
	for i := 0; i < len(crew_id); i++ {
		if crew_id[i] <= job_id[i] {
			totalCost += job_id[i] - crew_id[i]
		} else if crew_id[i] > job_id[i] {
			totalCost += crew_id[i] - job_id[i]
		}
	}
	return totalCost
}

// 渡された数値をByte配列に変換する
func divideNumberToBytes(number int) []byte {
	numberString := strconv.Itoa(number)
	// convert to [] bytes
	bytesArray := []byte(numberString)
	return bytesArray
}

