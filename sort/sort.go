package main

import (
	"fmt"
	"sort"
)

// Struct User
type User struct {
	Name string
	Age  int
}

// Tipe slice untuk User dengan implementasi sort.Interface
type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age // Urutkan berdasarkan umur
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	// 1. Mengurutkan slice sederhana
	ints := []int{5, 3, 4, 1, 2}
	sort.Ints(ints)
	fmt.Println("Sorted ints:", ints)

	strings := []string{"zebra", "apple", "mango", "banana"}
	sort.Strings(strings)
	fmt.Println("Sorted strings:", strings)

	// 2. Mengurutkan slice custom menggunakan sort.Interface
	users := UserSlice{
		{"Muhammad", 30},
		{"Ayom", 25},
		{"Izzuddin", 35},
	}
	sort.Sort(users)
	fmt.Println("Sorted users by age:", users)

	// 3. Mengurutkan slice custom menggunakan sort.Slice
	users2 := []User{
		{"Muhammad", 30},
		{"Ayom", 25},
		{"Izzuddin", 35},
	}
	sort.Slice(users2, func(i, j int) bool {
		return users2[i].Age < users2[j].Age
	})
	fmt.Println("Sorted users2 by age (using sort.Slice):", users2)

	// 4. Mengurutkan slice dalam urutan terbalik
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println("Sorted ints descending:", ints)

	sort.Sort(sort.Reverse(sort.StringSlice(strings)))
	fmt.Println("Sorted strings descending:", strings)

	// 5. Mengurutkan dengan banyak kriteria (umur lalu nama)
	users3 := []User{
		{"Muhammad", 30},
		{"Ayom", 25},
		{"Izzuddin", 30},
	}
	sort.Slice(users3, func(i, j int) bool {
		if users3[i].Age == users3[j].Age {
			return users3[i].Name < users3[j].Name // Jika umur sama, urutkan berdasarkan nama
		}
		return users3[i].Age < users3[j].Age
	})
	fmt.Println("Sorted users3 by age and name:", users3)

	// 6. Mengurutkan keys dari map
	m := map[string]int{
		"zebra": 3,
		"apple": 1,
		"mango": 2,
	}
	// Ekstrak keys
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	// Urutkan keys
	sort.Strings(keys)
	// Cetak map berdasarkan keys yang terurut
	fmt.Println("Sorted map by keys:")
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, m[k])
	}
}
