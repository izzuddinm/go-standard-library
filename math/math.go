package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. math.Abs: Mengembalikan nilai absolut dari sebuah angka.
	fmt.Println("=== Abs ===")
	fmt.Println("math.Abs(-5.5):", math.Abs(-5.5)) // Output: 5.5

	// 2. math.Ceil: Membulatkan angka ke atas (ke bilangan bulat terdekat).
	fmt.Println("\n=== Ceil ===")
	fmt.Println("math.Ceil(2.3):", math.Ceil(2.3)) // Output: 3

	// 3. math.Floor: Membulatkan angka ke bawah (ke bilangan bulat terdekat).
	fmt.Println("\n=== Floor ===")
	fmt.Println("math.Floor(2.7):", math.Floor(2.7)) // Output: 2

	// 4. math.Round: Membulatkan angka ke bilangan bulat terdekat.
	fmt.Println("\n=== Round ===")
	fmt.Println("math.Round(2.5):", math.Round(2.5)) // Output: 3
	fmt.Println("math.Round(2.4):", math.Round(2.4)) // Output: 2

	// 5. math.Sqrt: Menghitung akar kuadrat.
	fmt.Println("\n=== Sqrt ===")
	fmt.Println("math.Sqrt(16):", math.Sqrt(16)) // Output: 4

	// 6. math.Pow: Menghitung pangkat.
	fmt.Println("\n=== Pow ===")
	fmt.Println("math.Pow(2, 3):", math.Pow(2, 3)) // Output: 8

	// 7. math.Max: Mengembalikan nilai maksimum dari dua angka.
	fmt.Println("\n=== Max ===")
	fmt.Println("math.Max(10, 20):", math.Max(10, 20)) // Output: 20

	// 8. math.Min: Mengembalikan nilai minimum dari dua angka.
	fmt.Println("\n=== Min ===")
	fmt.Println("math.Min(10, 20):", math.Min(10, 20)) // Output: 10

	// 9. math.Mod: Menghitung sisa pembagian dua angka.
	fmt.Println("\n=== Mod ===")
	fmt.Println("math.Mod(10, 3):", math.Mod(10, 3)) // Output: 1

	// 10. math.Sin, math.Cos, math.Tan: Menghitung nilai trigonometri (radian).
	fmt.Println("\n=== Trigonometri ===")
	angle := math.Pi / 4                           // 45 derajat dalam radian
	fmt.Println("math.Sin(45°):", math.Sin(angle)) // Output: ~0.707
	fmt.Println("math.Cos(45°):", math.Cos(angle)) // Output: ~0.707
	fmt.Println("math.Tan(45°):", math.Tan(angle)) // Output: ~1

	// 11. math.Log: Menghitung logaritma natural (basis e).
	fmt.Println("\n=== Log ===")
	fmt.Println("math.Log(10):", math.Log(10)) // Output: ~2.302

	// 12. math.Exp: Menghitung eksponensial (e^x).
	fmt.Println("\n=== Exp ===")
	fmt.Println("math.Exp(1):", math.Exp(1)) // Output: ~2.718

	// 13. math.Hypot: Menghitung panjang sisi miring segitiga siku-siku.
	fmt.Println("\n=== Hypot ===")
	fmt.Println("math.Hypot(3, 4):", math.Hypot(3, 4)) // Output: 5

	// 14. math.Cbrt: Menghitung akar kubik dari sebuah angka.
	fmt.Println("\n=== Cbrt ===")
	fmt.Println("math.Cbrt(27):", math.Cbrt(27)) // Output: 3
}
