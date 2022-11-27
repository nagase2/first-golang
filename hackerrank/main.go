package main

import "fmt"

func main() {

// str := fmt.Sprintf("%02d",12)
// fmt.Println(str)
fmt.Println(dayOfProgrammer(2012))
}


func dayOfProgrammer(year int) string {
    // Write your code here

	// check if the year is leap year 
	feb:=28
	// 28 or 26



    // 30 August
    aug30 := 31 + feb + 31 + 30 + 31 + 30 + 31 + 31 

	fmt.Println(256 - aug30)

	// 256th is programmer day


	// 12.09.2016
	//str := fmt.Sprintf("%02d",12)

	return fmt.Sprintf("%02d.09.%d",12,year)
}
