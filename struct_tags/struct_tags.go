package main

import (
	"fmt"
	"reflect"
)

// Struct dengan tag
type Person struct {
	Name    string `json:"name" xml:"person_name"`
	Age     int    `json:"age" xml:"person_age"`
	Address string `json:"address" xml:"person_address"`
}

// Fungsi untuk membaca tag dalam struct
func readStructTags(value any) {
	valueType := reflect.TypeOf(value)

	// Loop melalui setiap field dalam struct
	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)

		// Mengakses tag dari field
		jsonTag := field.Tag.Get("json")
		xmlTag := field.Tag.Get("xml")

		// Menampilkan nama field dan tag-nya
		fmt.Printf("Field: %s, JSON Tag: %s, XML Tag: %s\n", field.Name, jsonTag, xmlTag)
	}
}

// Fungsi untuk membaca nilai dan tag dalam struct
func readStructValuesAndTags(value any) {
	valueType := reflect.TypeOf(value)
	valueValue := reflect.ValueOf(value)

	// Loop melalui setiap field dalam struct
	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		fieldValue := valueValue.Field(i)

		// Mengakses tag dari field
		jsonTag := field.Tag.Get("json")
		xmlTag := field.Tag.Get("xml")

		// Menampilkan nama field, nilai field, dan tag-nya
		fmt.Printf("Field: %s, Value: %v, JSON Tag: %s, XML Tag: %s\n", field.Name, fieldValue, jsonTag, xmlTag)
	}
}

// Struct dengan tag dan embedded struct
type Employee struct {
	Person          // Embedded struct
	Position string `json:"position" xml:"person_position"`
}

func main() {
	// Membuat objek Person
	p := Person{Name: "Alice", Age: 30, Address: "123 Main St"}

	// Membaca tags dalam struct Person
	fmt.Println("Reading Struct Tags for Person:")
	readStructTags(p)

	// Membaca nilai dan tags dalam struct Person
	fmt.Println("\nReading Struct Values and Tags for Person:")
	readStructValuesAndTags(p)

	// Membuat objek Employee (dengan embedded struct Person)
	e := Employee{Person: Person{Name: "Bob", Age: 40, Address: "456 Side St"}, Position: "Manager"}

	// Membaca tags dalam struct Employee
	fmt.Println("\nReading Struct Tags for Employee:")
	readStructTags(e)

	// Membaca nilai dan tags dalam struct Employee
	fmt.Println("\nReading Struct Values and Tags for Employee:")
	readStructValuesAndTags(e)
}
