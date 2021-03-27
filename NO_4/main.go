package main

import (
	"fmt"
	"math"
)

func main() {
	var numCount = 0
	var jumCount = 0
	var genCount = 0
	var ulngCount = 0
	var data = []int{1, 2, 3, 4, 5, 1, 4, 6, 8, 10}
	for _, y := range data {
		numCount += y
		jumCount++
	}

	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range data {
		if keys[entry] == true {
			ulngCount++
		}
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	for _, y := range data {
		if y%2 == 0 {
			genCount++
		}
	}

	var fibo = []int{}
	for i, _ := range data {
		if fb(5*data[i]*data[i]+4) || fb(5*data[i]*data[i]-4) {
			fibo = append(fibo, data[i])
		}

	}

	swp := true
	x := len(data)
	for swp {
		swp = false
		for i := 1; i < x; i++ {
			if data[i-1] < data[i] {
				data[i], data[i-1] = data[i-1], data[i]
				swp = true
			}
		}
	}

	fmt.Println("1.	Jumlah (Sum) angka di dalam array  = ", numCount)
	fmt.Println("2.	Berapa Jumlah angka berulang di dalam array tersebut = ", jumCount)
	fmt.Println("3.	Hapus angka berulang di array = ", list)
	fmt.Println("4.	Berapa jumlah bilangan genap  = ", genCount)
	fmt.Println("6.	Jumlah angka berulang di array  = ", ulngCount)
	fmt.Println("7.	Urutkan angka di dalam array   = ", data)
	fmt.Println("5.	Berapa jumlah bilangan fibonaci    = ", fibo, "jumlah =", len(fibo))
}

func fb(no int) bool {
	int_root := int(math.Sqrt(float64(no)))
	return int_root*int_root == no
}
