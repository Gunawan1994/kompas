package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var first string
	fmt.Println("Input password: ")
	fmt.Scanln(&first)
	fmt.Println(checkStr(first))
}

func checkStr(pass string) string {
	var dc string = "benar"
	if len(pass) < 10 {
		dc = "Salah karena kurang dari 10 karakter"
		return dc
	}

	for i, x := range pass {
		res1 := strings.Count(pass, string(x))
		if res1 > 1 {
			res2 := string(pass[i])
			var idn string
			match, _ := regexp.MatchString("([a-z]+)", res2)
			if !match {
				idn = "angka"
			} else {
				idn = "huruf"
			}
			dc = "Salah karena " + idn + " " + res2 + " " + "berulang lebih dari" + " " + string(fmt.Sprint(res1)) + "x"
			return dc
		}
	}
	return dc
}
