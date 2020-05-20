# Walaupun slice beda dengan array, mereka dicetak dengan
# tampilan yang sama oleh `fmt.Println`.
$ go run slices.go
emp: [  ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d:  [[0] [1 2] [2 3 4]]

# Cek [postingan blog keren](http://blog.golang.org/2011/01/go-slices-usage-and-internals.html) ini
# oleh tim Go untuk detil menyeluruh dari rancangan dan implementasi slice di Go.

# Sekarang kita sudah melihat array dan slice, kita akan
# mempelajari struktur data lain di Go: map.
