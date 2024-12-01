package main

import (
	"fmt"
	"time"
)

func main() {
	// Mendapatkan waktu saat ini
	now := time.Now()

	// Format waktu kustom: dd-MM-yyyy hh:mm:ss
	customFormat := "02-01-2006 03:04:05 PM"
	formattedCustomTime := now.Format(customFormat)
	fmt.Println("Waktu dengan format kustom:", formattedCustomTime)

	// Format waktu bawaan: RFC3339
	rfc3339Format := now.Format(time.RFC3339)
	fmt.Println("Waktu dengan format RFC3339:", rfc3339Format)

	// Format waktu bawaan: RFC1123
	rfc1123Format := now.Format(time.RFC1123)
	fmt.Println("Waktu dengan format RFC1123:", rfc1123Format)

	// Parsing waktu dengan RFC3339
	rfc3339TimeString := "2003-11-27T12:30:45Z"
	parsedRFC3339Time, err := time.Parse(time.RFC3339, rfc3339TimeString)
	if err != nil {
		fmt.Println("Error parsing waktu dengan RFC3339:", err)
	} else {
		fmt.Println("Waktu hasil parsing RFC3339:", parsedRFC3339Time)
	}

	// Membuat waktu dengan time.Date
	birthDate := time.Date(2003, 11, 27, 12, 30, 45, 0, time.UTC)
	fmt.Println("Tanggal lahir dengan format kustom:", birthDate.Format(customFormat))
	fmt.Println("Tanggal lahir dengan format RFC3339:", birthDate.Format(time.RFC3339))
	fmt.Println("Tanggal lahir dengan format RFC1123:", birthDate.Format(time.RFC1123))

	// Manipulasi waktu (menambah/mengurangi waktu)
	nextWeek := now.AddDate(0, 0, 7)
	fmt.Println("Satu minggu dari sekarang (kustom):", nextWeek.Format(customFormat))
	fmt.Println("Satu minggu dari sekarang (RFC3339):", nextWeek.Format(time.RFC3339))

	lastMonth := now.AddDate(0, -1, 0)
	fmt.Println("Satu bulan yang lalu (kustom):", lastMonth.Format(customFormat))
	fmt.Println("Satu bulan yang lalu (RFC1123):", lastMonth.Format(time.RFC1123))

	// Konversi zona waktu
	localTime := now.Local()
	fmt.Println("Waktu lokal (kustom):", localTime.Format(customFormat))
	fmt.Println("Waktu lokal (RFC3339):", localTime.Format(time.RFC3339))

	utcTime := now.UTC()
	fmt.Println("Waktu UTC (kustom):", utcTime.Format(customFormat))
	fmt.Println("Waktu UTC (RFC1123):", utcTime.Format(time.RFC1123))
}
