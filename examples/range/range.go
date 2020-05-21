// _range_ melakukan iterasi pada elemen-elemen pada
// berbagai struktur data. Mari kita lihat bagaimana
// cara menggunakan `range` untuk beberapa struktur
// data yang telah kita pelajari.

package main

import "fmt"

func main() {

	// Di sini kita menggunakan `range` untuk menjumlahkan
	// angka-angka pada sebuah slice.
	// Array juga bisa.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// `range` pada array dan slice menyediakan indeks
	// dan nilai dari tiap entri. Di atas kita tidak
	// butuh index, jadi kita menghiraukannya dengan
	// identifier kosong `_`. Tapi kadang kita ingin
	// informasi indeksnya.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// `range` pada map mengiterasi pasangan kunci/nilai.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` juga bisa mengiterasi hanya kunci dari map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// `range` pada string mengiterasi pada titik kode
	// Unicode. Nilai pertama adalah awalan index byte
	// dari `rune` dan yang kedua adalah `rune` itu
	// sendiri.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
