package main

import (
	"fmt"
	"sync"
)

func fibonacci(angka int, chn chan<- int){
	// fmt.Println("sudah masuk params",angka) // 10
	x, y := 0, 1
	for i := 0; i < angka; i++ {
		chn <- x
		x, y = y, x + y
	}
	defer close(chn)
}

func GanjilGenap(chn <-chan int){
	var wg sync.WaitGroup
	// temp_ganjil := []int{}
	// temp_genap := []int{}
	for data := range chn {
		// fmt.Println(data)
		wg.Add(1)
		data := data
		go func() {
			wg.Done()
			if data % 2 == 1 {
				// temp_ganjil = append(temp_ganjil, i)
				fmt.Println("Ganjil", data)
			} else if data % 2 == 0 {
				// temp_genap = append(temp_genap, i)
				fmt.Println("Genap", data)
			}
		}()
		wg.Wait()
	}
	// fmt.Println("Ganjil", temp_ganjil)
	// fmt.Println("Genap", temp_genap)
}

func TukarData(num int){

	channel := make(chan int, num)

	fibonacci(cap(channel), channel) // channel masuk dan di olah function ini pertama, dan mengembalikan isian yg telah di olah.
	GanjilGenap(channel) // mengambil isi dari channel yg telah di olah dari function atas.
}