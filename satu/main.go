package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	humanize "github.com/dustin/go-humanize"
)

var nama, durasi, name string
var arr1 []string
var arr [][]string
var id int
var biaya float64

// [
// 	["CODM","5.0"],["ML","4.4"]
// 	]

func addData() {
	fmt.Print("\nNama Customer : ")
	fmt.Scan(&nama)
	fmt.Print("Durasi sewa (jam) : ")
	fmt.Scan(&durasi)

	arr1 = []string{nama, durasi}
	arr = append(arr, arr1)

	sort.Slice(arr[:], func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})

}

func FormatRupiah(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(amount, 0)
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)
	return "Rp " + stringValue
}

func deleteData(ID int) {
	var before, after [][]string
	before = arr[:ID-1]
	after = arr[ID:]
	arr = append(after, before...)
}

func main() {
	var pilih int
	var x = true

	for x {
		fmt.Println("\n\t====Daftar Pilihan==== ")
		fmt.Println("1. Tampilkan Semua Data")
		fmt.Println("2. Tambah Data")
		fmt.Println("3. Hapus Data by ID")
		fmt.Println("4. Rata-rata durasi customer ")
		fmt.Println("5. Menampilkan 3 customer dengan durasi terpendek")
		fmt.Println("6. Menampilkan seluruh data dengan durasi dibawah rata-rata")
		fmt.Println("7. Keluar")
		fmt.Print("Masukkan Pilihan : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			fmt.Println("id \tNama \t\tDurasi \t\tBiaya")
			for i := 0; i < len(arr); i++ {
				var fltrating, err = strconv.ParseFloat(arr[i][1], 64)
				if err == nil {
					biaya = (fltrating * 1000) * 60
				}
				fmt.Printf("%d \t%s \t\t%s \t\t%s\n", i+1, arr[i][0], arr[i][1], FormatRupiah(biaya))
			}

		case 2:
			addData()
			fmt.Print(arr1)
		case 3:
			fmt.Print("\nMasukkan ID game yang akan dihapus: ")
			fmt.Scan(&id)

			deleteData(id)
			sort.Slice(arr[:], func(i, j int) bool {
				return arr[i][1] > arr[j][1]
			})

		case 4:
			var avg float64
			for i := 0; i < len(arr); i++ {
				var durasi, err = strconv.ParseFloat(arr[i][1], 64)
				if err == nil {
					avg += durasi
				}

			}

			var y float64 = float64(len(arr))

			var o float64 = avg / y
			fmt.Printf("\t%.1f jam\n", o)

		case 5:
			fmt.Println("3 Customer dengan durasi terpendek")
			fmt.Println("id \tNama \t\tDurasi \t\tBiaya")
			for i := 3; i >= 1; i-- {
				var durasi, err = strconv.ParseFloat(arr[i][1], 64)
				if err == nil {
					biaya = (durasi * 1000) * 60
				}
				fmt.Printf("%d \t%s \t\t%s \t\t%s\n", i+1, arr[i][0], arr[i][1], FormatRupiah(biaya))
			}

		case 6:
			var total float64
			fmt.Println("Customer dengan durasi kurang dari rata rata")
			fmt.Println("id \tNama \t\tDurasi \t\tBiaya")
			for i := 0; i < len(arr); i++ {

				var fltrating, err = strconv.ParseFloat(arr[i][1], 64)
				if err == nil {
					total += fltrating
					biaya = (fltrating * 1000) * 60
				}

				var y float64 = float64(len(arr))
				var o float64 = total / y
				var dur, er = strconv.ParseFloat(arr[i][1], 64)
				if er == nil {
					var x float64 = dur
					var z float64 = o
					if x < z {
						fmt.Printf("%d \t%s \t\t%s \t\t%s\n", i+1, arr[i][0], arr[i][1], FormatRupiah(biaya))
					}
				}
			}
		case 7:
			fmt.Print("Bye")
			x = false
		default:
			fmt.Print("Pilihan tidak tersedia")
		}
	}

}
