package main

import (
	"fmt"
	"strings"
	"unicode"
)

func verifyPassword(password string) error {
	var specialChar bool
	const minPass = 10
	const maxPass = 20
	var passLen int
	var errorStr string
	var upperCount = 0
	var lowerCount = 0
	var numCount = 0

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			numCount++
			passLen++
		case unicode.IsUpper(ch):
			upperCount++
			passLen++
		case unicode.IsLower(ch):
			lowerCount++
			passLen++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialChar = true
			passLen++
		case ch == ' ':
			passLen++
		}
	}
	appendError := func(err string) {
		if len(strings.TrimSpace(errorStr)) != 0 {
			errorStr += ", " + err
		} else {
			errorStr = err
		}
	}

	if upperCount < 2 {
		appendError("minimal 2 huruf besar")
	}
	if lowerCount < 5 {
		appendError("minimal 5 huruf kecil ")
	}
	if numCount < 2 {
		appendError("minimal 2 angka ")
	}
	if !specialChar {
		appendError("minimal special character (@#$%^&*()-+)")
	}
	if !(minPass <= passLen && passLen <= maxPass) {
		appendError(fmt.Sprintf("pajang password minimal %d maksimal %d", minPass, maxPass))
	}

	if len(errorStr) != 0 {
		return fmt.Errorf(errorStr)
	}
	return nil
}

func main() {
	password := "PP@swoyg55"
	err := verifyPassword(password)
	if err == nil {
		fmt.Println("password masuk kedalam kriteria", password)
	} else {
		fmt.Println(password, " ", err)
	}
}
