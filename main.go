package main

import (
	"fmt"
	"strings"
)

func main() {

	//	fmt.Println(checkPalindrome("aa3aa"))
	//var i = []int{1, 2, 3, 4, 99, 5, 6}
	// fmt.Println(adjacentElementsProduct(i))
	//fmt.Println(shapeArea(4))
	//fmt.Println(makeArrayConsecutive2(i))
	//fmt.Println(almostIncreasingSequence(i))
	/*
	var s = []string{
		"abc",
		"eeee",
		"abcd",
		"dcd",
	}
	allLongestStrings(s)

	var matrix = [][]int {
		{0,1,1,2},
		{0,5,0,0},
		{2,0,3,3},
	}
	fmt.Println(matrixElementsSum(matrix))
	*/
	commonCharacterCount("abca","xyzbac")

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
	return (max-min +1) - len(statues)
}

func almostIncreasingSequence(sequence []int) bool {

	var count=0
	for i:=1; i<len(sequence) ;i++  {

		if  sequence[i] <= sequence[i-1]{
			count++

			if count>1 {
				return false
			}
			if (i>1 && i<(len(sequence)-1) && sequence[i] <= sequence[i-2] && sequence[i+1] <= sequence[i-1]){
				return false
			}

		}
	}
	return true
}

func commonCharacterCount(s1 string, s2 string) int {

	var count  int=0

	for i:=0;i<len(s1);i++{
		for j:=0;j<len(s2);j++ {
			if s1[i]==s2[j]{
				count++
				s1=strings.Replace(s1,string(s1[i])," ",1)
				s2=strings.Replace(s2,string(s2[j])," ",1)
				break
			}
		}
	}

	return count
}


func allLongestStrings(inputArray []string) []string {

	var l =0
	var result []string
	for _,s := range inputArray{

		if len(s)>l{
			l=len(s)
		}
	}
	for _,s := range inputArray{

		if len(s)==l{
			result = append(result,s)
		}
	}
	return result
}

func matrixElementsSum(matrix [][]int) int {

	var sum=0
	for i:=0; i< len(matrix) ;i++  {
		for j:=0;j<len(matrix[0]) ;j++  {

			if matrix[i][j]==0{

				for k:=i;k< len(matrix); k++{
					matrix[k][j]=0
				}
			}else{
				sum += matrix[i][j]
			}
			fmt.Println(matrix)
			fmt.Println("sum:",sum,"\n")
		}
	}

	return sum

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



