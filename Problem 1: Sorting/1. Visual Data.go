package main

import "fmt"

func visual(arr []int) {
	var arr2 string
	temp := 0
	fmt.Println("\n")
	
	for j := 0 ; j < len(arr) ; j++{ //Mengetahui Nilai Teringgi sudatu array
			if arr[j] > temp{
				temp = arr[j]
			}
		} 
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
  
 func main(){
	fmt.Print("\nFor Visual Data only: ")
 	arr  := []int{1,4,5,3,8,2} 
	visual(arr)
}