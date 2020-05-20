// _Slice_ adalah tipe data utama dalam Go, memberikan
// antarmuka yang lebih kuat untuk deretan jika dibandingkan
// dengan array.

package main

import "fmt"

func main() {

	// Beda dengan array, slice digolongkan hanya oleh
	// elemen yang menyertainya (bukan jumlah elemen).
	// Untuk membuat slice kosong dengan panjang tidak nol,
	// gunakan `make`. Di sini kita membuat sebuah slice
	// dari `string` dengan panjang `3` (nilai awal nol).
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// Kita dapat menentukan dan mengambil elemen seperti
	// array.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` mengembalikan panjang dari slice.
	fmt.Println("len:", len(s))

	// Sebagai tambahan untuk operasi-operasi dasar, slice
	// mendukung beberapa hal lain yang membuatnya lebih kaya
	// daripada array. Salah satunya adalah `append`, yang
	// mengembalikan sebuah slice yang mengandung satu atau
	// lebih nilai baru. Catatan: kita harus menerima nilai
	// kembalian dari `append` karena kita mungkin mendapatkan
	// sebuah slice baru.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slice juga bisa disalin (dengan `copy`). Di sini kita
	// membuat slice kosong `c` dengan panjang yang sama
	// dengan `s` dan menyalin `s` ke `c`.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slice mendukung operator "slice" dengan sintaks
	// `slice[kecil:besar]`. Sebagai contoh, ini mengambil
	// sebuah slice dengan elemen `s[2]`, `s[3]`, and `s[4]`.
	l := s[2:5]
	fmt.Println("sl1:", l)

	// Slice ini berisi hingga (tapi tidak termasuk) `s[5]`.
	l = s[:5]
	fmt.Println("sl2:", l)

	// Dan slice ini berisi dari (dan termasuk) `s[2]`.
	l = s[2:]
	fmt.Println("sl3:", l)

	// Kita bisa menyatakan dan meng-inisialisasi sebuah
	// variabel untuk slice dengan satu baris kode.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slice bisa dibentuk kedalam struktur data multi
	// dimemsi. Panjang dari slice dalam bisa tidak sama,
	// beda seperti array multi dimensi.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
