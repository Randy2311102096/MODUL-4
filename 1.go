package main

import (
	"fmt"
)

// Fungsi faktorial
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Fungsi permutasi P(n, r)
func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

// Fungsi kombinasi C(n, r)
func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	// Masukkan 4 bilangan a, b, c, d
	var a, b, c, d int
	fmt.Println("Masukkan empat bilangan a, b, c, d (pisahkan dengan spasi):")
	fmt.Scan(&a, &b, &c, &d)

	// Permutasi dan kombinasi
	Pac := permutation(a, c)
	Cac := combination(a, c)
	Pbd := permutation(b, d)
	Cbd := combination(b, d)

	// Output hasil sesuai format gambar
	fmt.Printf("%d %d %d %d\n", a, b, c, d)
	fmt.Printf("%d %d\n", Pac, Cac)
	fmt.Printf("%d %d\n", Pbd, Cbd)
}
