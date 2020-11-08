package main

import "fmt"

func insertion(arr []int) {
	var arr2 string
	temp := 0
	fmt.Print("\n")
	for j := 0 ; j < len(arr) ; j++{ //Mengetahui Nilai Teringgi sudatu array
			if arr[j] > temp{
				temp = arr[j]
			}
		} 
	for i := 0; i < len(arr); i++ { //InsertionSort
	tmp := arr[i]
	j := i
			for j > 0 && arr[j-1] > tmp {
			  arr[j] = arr[j-1]
			  j--
			}
		arr[j] = tmp
		for i := temp ; i >  0 ; i -- 	{
			for j := 0 ; j < len(arr) ; j++{
				if arr[j] <  i{
					arr2 = " "
					fmt.Print(" ",arr2)
				}else {
					arr2 = "|"
					fmt.Print(" ",arr2)
				}
			}
		fmt.Print("\n")
	}
		for j := 0 ; j < len(arr) ; j++{
			fmt.Print(" ",arr[j])
		}
		fmt.Println("\n")
	}
}
  
 func main(){
	fmt.Print("\nFor Increment: \n")
 	arr  := []int{1,4,5,3,8,2} 
	insertion(arr)
}