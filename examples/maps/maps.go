// _Map_ adalah [tipe data asosiatif](http://en.wikipedia.org/wiki/Associative_array) di Go.
// (di bahasa pemrograman lain: _hash_ atau _dict_).

package main

import "fmt"

func main() {

	// Untuk membuat map kosong, gunakan `make`:
	// `make(map[tipe-kunci]tipe-nilai)`.
	m := make(map[string]int)

	// Atur pasangan kunci/nilai dengan sintaks umum
	// `nama[kunci] = nilai`
	m["k1"] = 7
	m["k2"] = 13

	// Mencetak sebuah map dengan `fmt.Println` akan
	// menunjukkan semua pasangan kunci/nilai.
	fmt.Println("map:", m)

	// Ambil nilai dari sebuah kunci dengan `nama[kunci]`.
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// `len` mengembalikan jumlah pasangan kunci/nilai
	// saat dipanggil pada sebuah map.
	fmt.Println("len:", len(m))

	// `delete` menghapus pasangan kunci/nilai dari
	// sebuah map.
	delete(m, "k2")
	fmt.Println("map:", m)

	// Nilai kembalian kedua yang bersifat opsional saat
	// mengambil nilai dari sebuah map, menandakan apabila
	// sebuah kunci ada pada map. Hal ini bisa digunakan
	// untuk menghilangkan ambiguitas antara kunci yang
	// tidak ada dan kunci dengan nilai nol (seperti `0`
	// atau `""`. Di sini kita tidak perlu nilai itu,
	// jadi kita menghiraukannya dengan _penanda kosong_
	// `_`.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Kamu juga bisa menyatakan dan memberi nilai awal
	// sebuah map baru dengan satu baris kode.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
