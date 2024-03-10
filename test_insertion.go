package main

import "fmt"

func insertionSort1(n int32, arr []int32) {
	arr1 := make([]int32, n)
    copy(arr1, arr)
	fmt.Println(arr1)
    for i :=n-2 ;i>=0;i--{
        if arr[n-1]<arr[i]{
            arr1[i+1] =arr[i]
			//fmt.Println(arr1[i+1],arr[i])
            
        }else {
            arr1[i+1]=arr[n-1]
        }
        fmt.Println(arr1)
    }
    // Write your code here

}
func main(){
	// var x = []int32{2,10,4,6,8,7,9,3}
	// insertionSort1(8,x)
	arr := []int{8, 7, 6, 5, 4, 3, 2, 1}
	start := 0
    end := 1
    shift := 1

    // Shift the specified range
    shiftedArr := make([]int, len(arr))
    copy(shiftedArr, arr)

    for i := start; i < end; i++ {

        shiftedArr[i+shift] = arr[i]
    }
	shiftedArr[start] = arr[end]
    fmt.Println("Original array:", arr)
    fmt.Println("Shifted array:", shiftedArr)
}