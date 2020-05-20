// Statemen _switch_ menyatakan kondisional dengan banyak
// percabangan.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Ini adalah sebuah `switch` sederhana.
	i := 2
	fmt.Print("Tulis ", i, " sebagai ")
	switch i {
	case 1:
		fmt.Println("satu")
	case 2:
		fmt.Println("dua")
	case 3:
		fmt.Println("tiga")
	}

	// Kamu bisa menggunakan koma untuk memisahkan beberapa
	// ekspresi dalam statemen `case` yang sama.
	// Kita juga menggunakan opsional `default` juga di contoh
	// ini.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Ini weekend")
	default:
		fmt.Println("Ini weekday")
	}

	// `switch` tanpa sebuah ekspresi adalah sebuah cara lain
	// untuk menyatakan logika if/else.
	// Di sini kita juga menunjukkan kalau ekspresi `case`
	// bisa dinyatakan tanpa konstanta.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Ini sebelum tengah hari")
	default:
		fmt.Println("Ini setelah tengah hari")
	}

	// Sebuah tipe `switch` membandingkan tipe, bukan nilai.
	// Kamu bisa menggunakan ini untuk mencaritahu tipe dari
	// sebuah nilai interface.
	// Di contoh ini, variabel `t` akan punya tipe sesuai dengan
	// kondisinya.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Aku sebuah bool")
		case int:
			fmt.Println("Aku sebuah int")
		default:
			fmt.Printf("Tidak tahu tipe %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
