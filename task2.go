package main

import (
	"fmt"
	"sync"
)

func fibonacci(angka int, chn chan<- int){
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
	for v := range chn {
		wg.Add(1)
		go func() {
			wg.Done()
			if v % 2 == 1 {
				// temp_ganjil = append(temp_ganjil, i)
				fmt.Println("Ganjil", v)
			} else if v % 2 == 0 {
				// temp_genap = append(temp_genap, i)
				fmt.Println("Genap", v)
			}
		}()
		wg.Wait()
	}
	// fmt.Println("Ganjil", temp_ganjil)
	// fmt.Println("Genap", temp_genap)
}

func TukarData(num int){

	channel := make(chan int, num)

	fibonacci(cap(channel), channel)
	GanjilGenap(channel)
}