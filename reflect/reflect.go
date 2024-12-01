package main

import (
	"fmt"
	"reflect"
)

// Struct Example
type Sample struct {
	Name string
	Age  int
}

// Membaca Field dalam Struct
func readStruct(value any) {
	valueType := reflect.TypeOf(value)
	valueValue := reflect.ValueOf(value)

	// Menampilkan nama tipe
	fmt.Println("Type Name:", valueType.Name())

	// Loop untuk membaca setiap field
	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		fieldValue := valueValue.Field(i)
		fmt.Printf("Field: %s, Value: %v\n", field.Name, fieldValue)
	}
}

// Pointer Example
func readPointer(value any) {
	valueType := reflect.TypeOf(value)
	valueValue := reflect.ValueOf(value)

	// Memeriksa apakah nilai adalah pointer
	if valueType.Kind() == reflect.Ptr {
		// Mengambil nilai yang ditunjukkan pointer
		valueValue = valueValue.Elem()
		valueType = valueValue.Type()
	}

	// Menampilkan tipe dan nilai
	fmt.Println("Type Name:", valueType.Name())
	fmt.Println("Value:", valueValue)
}

// Map Example
func readMap(value any) {
	valueType := reflect.TypeOf(value)
	valueValue := reflect.ValueOf(value)

	// Memeriksa apakah value adalah map
	if valueType.Kind() == reflect.Map {
		// Loop untuk membaca elemen dalam map
		for _, key := range valueValue.MapKeys() {
			val := valueValue.MapIndex(key)
			fmt.Printf("Key: %v, Value: %v\n", key, val)
		}
	}
}

// Slice Example
func readSlice(value any) {
	valueType := reflect.TypeOf(value)
	valueValue := reflect.ValueOf(value)

	// Memeriksa apakah value adalah slice
	if valueType.Kind() == reflect.Slice {
		// Loop untuk membaca elemen dalam slice
		for i := 0; i < valueValue.Len(); i++ {
			elem := valueValue.Index(i)
			fmt.Printf("Element %d: %v\n", i, elem)
		}
	}
}

// Mengubah Nilai dengan Reflect
func modifyValue(value any) {
	valueValue := reflect.ValueOf(value)

	// Memeriksa apakah value adalah pointer dan dapat dimodifikasi
	if valueValue.Kind() == reflect.Ptr && valueValue.Elem().CanSet() {
		// Mengubah nilai field dalam struct
		valueValue.Elem().FieldByName("Name").SetString("Updated Name")
	}
}

// Mengetahui Tipe Variabel Secara Dinamis
func printType(value any) {
	// Menampilkan tipe data secara dinamis
	fmt.Println("Tipe variabel:", reflect.TypeOf(value))
}

// Memeriksa Field Tersembunyi (Unexported)
func readUnexportedField(value any) {
	valueValue := reflect.ValueOf(value)
	field := valueValue.FieldByName("unExportedField")

	// Memeriksa dan memodifikasi field yang tidak diekspor
	if field.IsValid() && field.CanSet() {
		field.SetString("Modified")
	}
}

func main() {
	// Contoh Struct
	sample := Sample{Name: "John Doe", Age: 30}
	fmt.Println("Read Struct:")
	readStruct(sample)

	// Contoh Pointer
	fmt.Println("\nRead Pointer:")
	readPointer(&sample)

	// Contoh Map
	studentGrades := map[string]int{
		"John":  85,
		"Alice": 92,
		"Bob":   78,
	}
	fmt.Println("\nRead Map:")
	readMap(studentGrades)

	// Contoh Slice
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("\nRead Slice:")
	readSlice(numbers)

	// Mengubah Nilai dengan Reflect
	fmt.Println("\nBefore Modify:", sample)
	modifyValue(&sample)
	fmt.Println("After Modify:", sample)

	// Mengetahui Tipe Variabel
	fmt.Println("\nTipe Variabel:")
	printType(10)           // int
	printType("Hello, Go!") // string
	printType(3.14)         // float64
	printType(true)         // bool

	// Memeriksa Field Tersembunyi
	fmt.Println("\nBefore Modify Unexported Field:")
	myStruct := struct {
		ExportedField   string
		unExportedField string
	}{ExportedField: "Visible", unExportedField: "Hidden"}

	readUnexportedField(&myStruct)
	fmt.Println("After Modify Unexported Field:", myStruct)
}
