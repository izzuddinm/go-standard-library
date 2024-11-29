package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Membuat linked list baru
	data := list.New()

	// 1. PushBack: Menambahkan elemen di belakang
	data.PushBack("Muhammad")
	data.PushBack("Ayom")
	data.PushBack("Izzuddin")
	fmt.Println("=== Setelah PushBack ===")
	printList(data)

	// 2. PushFront: Menambahkan elemen di depan
	data.PushFront("Dr.")
	fmt.Println("\n=== Setelah PushFront ===")
	printList(data)

	// 3. InsertBefore: Menambahkan elemen sebelum elemen tertentu
	first := data.Front() // Elemen pertama
	data.InsertBefore("Mr.", first)
	fmt.Println("\n=== Setelah InsertBefore ===")
	printList(data)

	// 4. InsertAfter: Menambahkan elemen setelah elemen tertentu
	last := data.Back() // Elemen terakhir
	data.InsertAfter("Ph.D", last)
	fmt.Println("\n=== Setelah InsertAfter ===")
	printList(data)

	// 5. Remove: Menghapus elemen tertentu
	data.Remove(first)
	fmt.Println("\n=== Setelah Remove ===")
	printList(data)

	// 6. Len: Mendapatkan panjang linked list
	fmt.Printf("\n=== Panjang Linked List ===\nLength: %d\n", data.Len())

	// 7. Iterasi dari belakang ke depan
	fmt.Println("\n=== Iterasi dari Belakang ke Depan ===")
	for el := data.Back(); el != nil; el = el.Prev() {
		fmt.Println(el.Value)
	}
}

// Fungsi untuk mencetak semua elemen dalam linked list
func printList(data *list.List) {
	for el := data.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value)
	}
}
