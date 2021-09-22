package main

import (
	"fmt"
	"math"
	"unicode"
)

var penyebut float64
var pembilang float64

func row(i float64, j float64) float64 {

	if i <= 1 {
		fmt.Printf("%.1f", j)
		penyebut += j
		return j
	}

	k := row(i-1, j)
	fmt.Printf(" + %.1f", j*i*(math.Pow(10, i-1)))
	penyebut += j*i*(math.Pow(10, i-1))
	return k
}

func d(i float64) float64 {
	if i <= 1 {
		penyebut += 1
		pembilang += 2
		fmt.Print("1/2x")
		return i
	}
	y := d(i - 1)
	penyebut += i
	x := i*2
	pembilang += x
	fmt.Printf(" + %.0f/%.0fx", i, x)

	return y
}

func isNtCap(i int, word string) int {
	if i < 0 {
		return i
	}
	y := rune(word[i])
	if unicode.IsLower(y) {
		pembilang = pembilang + 1
	}

	return isNtCap(i-1, word)
}

func menu() {
	var i int
	var max float64
	var num float64
	var word string

	fmt.Println("\n--------------------")
	fmt.Println("1. Deret a")
	fmt.Println("2. Deret b")
	fmt.Println("3. Jumlah huruf non kapital")
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&i)

	switch i {
	case 1:
		fmt.Print("Bilangan Pertama: ")
		fmt.Scan(&num)
		fmt.Print("Batas : ")
		fmt.Scan(&max)
		row(max, num)
		fmt.Printf(" = %.2f", penyebut)
	case 2:
		fmt.Print("Batas : ")
		fmt.Scan(&max)

		d(max)
		fmt.Printf(" = %.0f/%.0fx", penyebut, pembilang)
		
	case 3:
		fmt.Print("Masukkan kata: ")
		fmt.Scan(&word)

		isNtCap(len(word)-1, word)
		fmt.Printf("Jumlah : %.0f", pembilang)
	default:
		fmt.Print("Menu tidak ada")
		menu()
	}
}

func main() {
	menu()
}