package main

import "fmt"

// channel = channel bisa mengirim dan menerima data.
// channel <- data = mengirim data ke channel.
// data <- channel = menerima data dari channel.

func sum(s []int, chn chan int){
	sum := 0
	for _, v := range s{
		sum += v
	}
	chn <- sum // mengirim data sum ke dalam "chnl" atau channel.
}

func Try(){
	data := []int{7, 2, 8, -9, 4, 0}
	// fmt.Println(data[:len(data)/2])
	// fmt.Println(data[len(data)/2:])

	chnl := make(chan int) // untuk membuat chanel, di sini "chnl" dengan tipe datanya "int".

	go sum(data[:len(data)/2], chnl) //[0:4] => [7 2 8] slice data.
	go sum(data[len(data)/2:], chnl) //[4:0] => [-9, 4, 0] slice data.

	x := <-chnl
	y := <-chnl
	
	fmt.Println(x, y, x+y)
	defer close(chnl) // menutup channel mau berhasil atau error setelah di gunakan agar tidak memakan memory.
}