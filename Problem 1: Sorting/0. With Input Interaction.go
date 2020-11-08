package main

import "fmt"

func insertion(ops int,num int,arr []int) {
	var arr2 string
	temp := 0
	opt := ops
	fmt.Print("\n")
	for j := 0 ; j < len(arr) ; j++{ //mengetahui nilai teringgi sudatu array
			 if arr[j] > temp{
				 temp = arr[j]
			}
		} 
	for i := 0; i < num; i++ { //InsertionSort
	tmp := arr[i]
	j := i
		if opt == 0 {
			for j > 0 && arr[j-1] < tmp {
			  arr[j] = arr[j-1]
			  j--
			}
		}else{
			for j > 0 && arr[j-1] > tmp {
			  arr[j] = arr[j-1]
			  j--
			}
		}
		arr[j] = tmp
		for i := temp ; i >  0 ; i -- 	{
			for j := 0 ; j < num ; j++{
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
		for j := 0 ; j < num ; j++{
			fmt.Print(" ",arr[j])
		}
		fmt.Println("\n")
		}
  }
  
 func main(){
	var opt int
	var num int
 	arr := []int{0,0,0,0,0,0,0,0,0,0}
	fmt.Print("Masukkan jumlah array (max:10) : ")
	fmt.Scan(&num)
	
	fmt.Print("Masukkan Nilai ",num," array : \n")
	for i:=0;i<num;i++{
		fmt.Scan(&arr[i])
	}
	fmt.Print("\nKetik:\n------\n0. Untuk Decrement\n1. Untuk Increment : ")
	fmt.Scan(&opt)
	insertion(opt,num,arr)
}
