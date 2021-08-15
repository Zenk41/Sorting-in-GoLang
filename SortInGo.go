package main

import (
	"fmt"
	"os"
)

func main() {
	var mInput int //input Menu pilihan

	fmt.Println("Program Bubble sort dan Insertion sort")
	fmt.Println("1. Bubble sort")
	fmt.Println("2. Insertion sort")
	fmt.Println("3. Exit")
	fmt.Print("Masukkan Pilihan anda : ")
	fmt.Scanln(&mInput)

	switch mInput {
	case 1:
		BubbleSort()
		main()
	case 2:
		insertionsort()
		main()
	case 3:
		os.Exit(0)
	default:
		fmt.Println("Anda salah memasukkan pilihan, silahkan pilih kembali")
		fmt.Println()
		main()
	}

}


func BubbleSort() []int {
	fmt.Println("~ Bubble Sort ~ ")
	//Input Jumlah Element
	fmt.Printf("Masukkan Jumlah Bilangan : ")
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Masukkan Bilangan ke-%d : ", (i + 1))
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println()
	fmt.Println("Bilangan yang akan di sortir dengan Bubble Sort:", arr)
	fmt.Println()

	iteration := 1
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				fmt.Printf("Iterasi ke - %d Membandingkan bilangan %d dengan %d : %d\n", iteration, (j + 1), (j + 2), arr)
			} else {
				fmt.Printf("Iterasi ke - %d Membandingkan bilangan %d dengan %d : %d\n", iteration, (j + 1), (j + 2), arr)
			}
		}
		iteration += 1
		fmt.Println()
	}
	fmt.Printf("Sesudah Bubble Sort : %d \n\n", arr)
	return arr

}

func insertionsort() []int {
	fmt.Println("~ Insertion Sort ~ ")
	//Input Jumlah Element
	fmt.Printf("Masukkan Jumlah Bilangan : ")
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Masukkan Bilangan ke-%d : ", (i + 1))
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println()
	fmt.Println("Bilangan yang akan di sortir dengan Insertion Sort: ", arr)
	fmt.Println()

	iteration := 0
	var n = len(arr)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j = j - 1
		}
		iteration += 1
		fmt.Printf("Iterasi ke -%d : %d \n", iteration, arr)
	}
	fmt.Printf("Sesudah Insertion Sort : %d \n\n", arr)
	return arr
}
