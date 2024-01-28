package main

import "fmt"


// Struct
type bilangan struct {
	prima []int
	ganjil []int
	genap []int
	fibonacci int
}

// Method
func (temp_1 bilangan) deretanPrima() []int{
	return temp_1.prima
}
func (temp_1 bilangan) deretanGanjil() []int{
	return temp_1.ganjil
}
func (temp_1 bilangan) deretanGenap() []int{
	return temp_1.genap
}
// func (temp_1 bilangan) deretanFibonacci() []int{
// 	return temp_1.fibonacci
// }


func isPrime(n int) bool { // 5
    if n <= 1 {
        return false
    }
    for i := 2; i * i <= n; i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}


func Cetak(angka int){
	temp_prima := []int{}
	temp_ganjil := []int{}
	temp_genap := []int{}
	// temp_fibonacci := []int{}

	for i := 1; i <= angka; i++ { // 1,2,3,4,5,....40
		if i % 2 == 0 {
			// fmt.Println(i)
			temp_genap = append(temp_genap, i)
		} else if i % 2 == 1 {
			// fmt.Println(i)
			temp_ganjil = append(temp_ganjil, i)
		}

		if isPrime(i) {
            temp_prima = append(temp_prima, i)
        }
	}

	var data bilangan = bilangan{
		prima: temp_prima,
		ganjil: temp_ganjil,
		genap: temp_genap,
		fibonacci: 12,
	}

	fmt.Println("sudah masuk", data)
}