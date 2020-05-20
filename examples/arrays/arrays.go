// Di Go, sebuah _array_ adalah urutan dengan angka dari
// elemen-elemen dengan panjang tertentu.

package main

import "fmt"

func main() {

	// Di sini kita membuat sebuah array `a` yang akan
	// menampung tepat 5 `int`. Tipe dari elemen-elemen
	// dan panjangnya merupakan bagian dari tipe array.
	// Standarnya, sebuah array itu bernilai nol (zero
	// valued), untuk `int` artinya angka `0`.
	var a [5]int
	fmt.Println("emp:", a)

	// Kita dapat mengatur nilai pada sebuah index dengan
	// menggunakan `array[index] = nilai`, dan mengambil
	// nilainya dengan `array[index]`.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// Fungsi terpasang `len` mengembalikan panjang dari
	// sebuah array.
	fmt.Println("len:", len(a))

	// Gunakan sintaks ini untuk menyatakan dan memberi
	// nilai awal array dalam satu baris.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Tipe array adalah satu dimensi, tapi kamu bisa
	// membuat tipe struktur data multi dimensi.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
