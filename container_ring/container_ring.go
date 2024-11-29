package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Membuat ring dengan kapasitas 5
	r := ring.New(5)

	// Mengisi ring dengan data
	for i := 0; i < r.Len(); i++ {
		r.Value = i + 1
		r = r.Next()
	}

	fmt.Println("=== Isi Ring ===")
	printRing(r)

	// 1. Mengakses elemen berikutnya
	fmt.Println("\n=== Elemen Berikutnya (Next) ===")
	fmt.Println(r.Value)

	// 2. Mengakses elemen sebelumnya
	fmt.Println("\n=== Elemen Sebelumnya (Prev) ===")
	fmt.Println(r.Prev().Value)

	// 3. Melakukan iterasi pada ring
	fmt.Println("\n=== Iterasi pada Ring ===")
	r.Do(func(value any) {
		fmt.Println(value)
	})

	// 4. Menambahkan elemen baru (mengganti nilai pada node tertentu)
	r.Value = "Replaced"
	fmt.Println("\n=== Setelah Mengganti Elemen Saat Ini ===")
	printRing(r)

	// 5. Menambahkan ring ke ring lain
	newRing := ring.New(2)
	newRing.Value = "New1"
	newRing.Next().Value = "New2"

	// Hubungkan kedua ring
	r.Link(newRing)

	fmt.Println("\n=== Setelah Menggabungkan Ring ===")
	printRing(r)

	// 6. Menghapus elemen
	r.Unlink(2) // Menghapus 2 node berikutnya
	fmt.Println("\n=== Setelah Menghapus Elemen ===")
	printRing(r)
}

// Fungsi untuk mencetak semua elemen dalam ring
func printRing(r *ring.Ring) {
	r.Do(func(value any) {
		fmt.Println(value)
	})
}
