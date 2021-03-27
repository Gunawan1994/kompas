package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("artikel.txt")

	if err != nil {
		log.Fatal(err)
	}
	var countVar = 0
	for char := 'a'; char <= 'z'; char++ {
		countVar = 0
		for _, y := range string(content) {
			if string(char) == strings.ToLower(string(y)) {
				countVar++
			}
		}
		fmt.Printf("\n%s = %d", strings.ToUpper(string(char)), countVar)
	}
}

// tampilkan huruf a - z dengan perulngan, lakukan perulangan di dalamnya untuk setiap string/huruf di dalam artikel.txt
// jika dalam perulangan terdapat huruf yg sama (lowercase) maka variable countVar kan ditambah sebanyak huruf
// yg dia temukan di dlam perulangan
