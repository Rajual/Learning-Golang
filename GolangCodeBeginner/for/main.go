package main

import "fmt"

func main() {
	//For clasico
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//For continuo (while)ç
	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}
	/*
		//For forever
		k := 0
		for {
			fmt.Println(k)
			k++
		}*/
	//For range (Para map, array,string, etc...)
	nums := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, value := range nums {
		fmt.Println("Hola Mundo")
		value *= 2
		nums[i] *= 2
	}
	fmt.Println(nums)

}
