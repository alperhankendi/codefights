package main

import (
	"fmt"
)

func main() {

	//	fmt.Println(checkPalindrome("aa3aa"))
	var i = []int{8,4,1,3}
	// fmt.Println(adjacentElementsProduct(i))
	//fmt.Println(shapeArea(4))
	makeArrayConsecutive2(i)
}


func makeArrayConsecutive2(statues []int) int {

	min :=statues[0]
	max :=statues[0]

	for _,value := range statues{

		if min>value{
			min=value
		}
		if max<value{
			max=value
		}
	}
	counter:=0
	for i:=min;i<max ;i++  {

		for _,value := range statues{

			if value==i{
				break;
			}
			counter++
		}


	}
	return counter
}


func shapeArea(n int) int {

	var r int=0
	for i:=1;i <n ;i++  {

		r += (i*2)*2
	}
	return  r+1
}



func adjacentElementsProduct(inputArray []int) int {

	var r int=-1000
	for i:=0;i< len(inputArray)-1 ;i++  {
		m := inputArray[i] * inputArray[i+1]
		if m>r{
			r =m
		}
	}

	return r
}


func checkPalindrome(inputString string) bool {

	var reverseStr string
	for i := len(inputString)-1; i >= 0; i-- {
		reverseStr += string(inputString[i])
	}
	fmt.Println(reverseStr)

	return inputString ==reverseStr

}



