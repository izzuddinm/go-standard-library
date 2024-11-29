package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1. strconv.ParseBool: Mengonversi string menjadi boolean.
	booleans := []string{"true", "True", "TRUE", "false", "False", "FALSEs"}
	fmt.Println("=== ParseBool ===")
	for _, boolean := range booleans {
		result, err := strconv.ParseBool(boolean)
		handleError(result, err)
	}

	// 2. strconv.Atoi: Mengonversi string menjadi integer.
	integerStrings := []string{"1000", "23000", "False"}
	fmt.Println("\n=== Atoi ===")
	for _, integerString := range integerStrings {
		result, err := strconv.Atoi(integerString)
		handleError(result, err)
	}

	// 3. strconv.Itoa: Mengonversi integer menjadi string.
	fmt.Println("\n=== Itoa ===")
	integers := []int{42, -100, 5000}
	for _, integer := range integers {
		result := strconv.Itoa(integer)
		fmt.Println("Converted integer to string:", result)
	}

	// 4. strconv.ParseFloat: Mengonversi string menjadi float64.
	fmt.Println("\n=== ParseFloat ===")
	floatStrings := []string{"3.14", "2.71", "NaN"}
	for _, floatString := range floatStrings {
		result, err := strconv.ParseFloat(floatString, 64)
		handleError(result, err)
	}

	// 5. strconv.FormatFloat: Mengonversi float64 menjadi string.
	fmt.Println("\n=== FormatFloat ===")
	floats := []float64{3.14159, 2.71828, -1.41421}
	for _, floatVal := range floats {
		// 'f' adalah format fixed-point, 2 adalah jumlah angka desimal.
		result := strconv.FormatFloat(floatVal, 'f', 2, 64)
		fmt.Println("Converted float to string:", result)
	}

	// 6. strconv.ParseInt: Mengonversi string menjadi integer (basis yang ditentukan).
	fmt.Println("\n=== ParseInt ===")
	intStrings := []string{"1010", "FF", "123"}
	for _, intString := range intStrings {
		// Basis 10, atau 16 untuk hexadecimal.
		result, err := strconv.ParseInt(intString, 10, 64)
		handleError(result, err)
	}

	// 7. strconv.FormatInt: Mengonversi integer menjadi string (basis yang ditentukan).
	fmt.Println("\n=== FormatInt ===")
	integers64 := []int64{10, 255, -42}
	for _, intVal := range integers64 {
		result := strconv.FormatInt(intVal, 10) // Basis 10
		fmt.Println("Converted int64 to string:", result)
	}

	// 8. strconv.FormatUint: Mengonversi uint menjadi string.
	fmt.Println("\n=== FormatUint ===")
	uintValues := []uint64{10, 255, 42}
	for _, uintVal := range uintValues {
		result := strconv.FormatUint(uintVal, 10) // Basis 10
		fmt.Println("Converted uint64 to string:", result)
	}
}

// Fungsi untuk menangani error dan menampilkan hasil.
func handleError(data any, err error) {
	if err != nil {
		fmt.Println("Error:", err, "| Data:", data)
		return
	}
	fmt.Println("Success:", data)
}
