package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1. strings.HasPrefix: Mengecek apakah string diawali oleh substring tertentu.
	result1 := strings.HasPrefix("Muhammad Ayom Izzuddin", "Muhammad")
	fmt.Println("HasPrefix 'Muhammad':", result1) // Output: true

	// 2. strings.HasSuffix: Mengecek apakah string diakhiri oleh substring tertentu.
	result2 := strings.HasSuffix("Muhammad Ayom Izzuddin", "Izzuddin")
	fmt.Println("HasSuffix 'Izzuddin':", result2) // Output: true

	// 3. strings.Count: Menghitung jumlah kemunculan substring dalam string.
	result3 := strings.Count("Muhammad Ayom Izzuddin Ayom", "Ayom")
	fmt.Println("Count 'Ayom':", result3) // Output: 2

	// 4. strings.Index: Mendapatkan posisi indeks pertama dari substring dalam string.
	result4 := strings.Index("Muhammad Ayom Izzuddin", "Ayom")
	fmt.Println("Index of 'Ayom':", result4) // Output: 9

	// 5. strings.LastIndex: Mendapatkan posisi indeks terakhir dari substring dalam string.
	result5 := strings.LastIndex("Muhammad Ayom Izzuddin Ayom", "Ayom")
	fmt.Println("LastIndex of 'Ayom':", result5) // Output: 24

	// 6. strings.Repeat: Mengulang string beberapa kali.
	result6 := strings.Repeat("Go ", 3)
	fmt.Println("Repeated string:", result6) // Output: Go Go Go

	// 7. strings.Join: Menggabungkan elemen slice menjadi satu string dengan delimiter.
	words := []string{"Go", "is", "awesome"}
	result7 := strings.Join(words, " ")
	fmt.Println("Joined string:", result7) // Output: Go is awesome

	// 8. strings.Fields: Memecah string menjadi slice berdasarkan spasi (mengabaikan spasi kosong).
	result8 := strings.Fields("  Go is   fun  ")
	fmt.Println("Fields:", result8) // Output: [Go is fun]

	// 9. strings.TrimSpace: Menghapus semua spasi di awal dan akhir string (lebih spesifik daripada strings.Trim).
	result9 := strings.TrimSpace("   Hello, Go!   ")
	fmt.Println("Trimmed with TrimSpace:", result9) // Output: Hello, Go!

	// 10. strings.Compare: Membandingkan dua string secara leksikografis.
	result10 := strings.Compare("Go", "go")
	fmt.Println("Compare 'Go' and 'go':", result10) // Output: -1 (Go < go secara leksikografis)
}
