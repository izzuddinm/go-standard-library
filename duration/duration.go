package main

import (
	"fmt"
	"time"
)

func main() {
	// Membuat durasi secara manual
	duration := time.Duration(5 * time.Hour) // 5 jam
	fmt.Println("Durasi manual:", duration)

	// Durasi dari string
	durationFromString, err := time.ParseDuration("2h45m30s")
	if err != nil {
		fmt.Println("Error parsing durasi:", err)
	} else {
		fmt.Println("Durasi dari string:", durationFromString)
	}

	// Operasi dengan durasi
	startTime := time.Now()
	fmt.Println("Waktu mulai:", startTime)

	// Menambah durasi ke waktu
	endTime := startTime.Add(2*time.Hour + 30*time.Minute)
	fmt.Println("Waktu setelah 2 jam 30 menit:", endTime)

	// Mengurangi durasi dari waktu
	previousTime := startTime.Add(-1 * time.Hour)
	fmt.Println("Waktu 1 jam sebelumnya:", previousTime)

	// Durasi antara dua waktu
	time1 := time.Date(2023, 12, 1, 14, 0, 0, 0, time.UTC)
	time2 := time.Date(2024, 12, 1, 14, 0, 0, 0, time.UTC)
	durationBetween := time2.Sub(time1)
	fmt.Println("Durasi antara dua waktu:", durationBetween)
	fmt.Printf("Dalam hari: %.0f hari\n", durationBetween.Hours()/24)

	// Konversi durasi
	fmt.Println("Durasi dalam detik:", durationBetween.Seconds())
	fmt.Println("Durasi dalam menit:", durationBetween.Minutes())
	fmt.Println("Durasi dalam jam:", durationBetween.Hours())

	// Durasi sebagai string
	fmt.Println("Durasi dalam string:", durationBetween.String())

	// Sleep menggunakan durasi
	fmt.Println("Menunggu 3 detik...")
	time.Sleep(3 * time.Second) // Tidur selama 3 detik
	fmt.Println("Selesai menunggu.")

	// Loop dengan interval durasi
	fmt.Println("Loop dengan interval 1 detik:")
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for i := 0; i < 3; i++ { // Loop 3 kali
		fmt.Println("Tick pada:", <-ticker.C)
	}
}
