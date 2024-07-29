package main

import "fmt"

// Abstract menggunakan Interface
type Animal interface {
	GetName() string
	GetSpesies() string
	GetAge() int
	MakeSound() string
}

// Enacpulation dan Inheritance memakai Struct
type Hewan struct {
	name    string
	spesies string
	age     int
}

func (h *Hewan) GetName() string {
	return h.name
}

func (h *Hewan) GetSpesies() string {
	return h.spesies
}

func (h *Hewan) GetAge() int {
	return h.age
}

type Singa struct {
	Hewan
}

func (s *Singa) MakeSound() string {
	return "oh noo oh my god"
}

type Gajah struct {
	Hewan
}

func (g *Gajah) MakeSound() string {
	return "oh tydack sakitt"
}

type Monyet struct {
	Hewan
}

func (m *Monyet) MakeSound() string {
	return "hallo dek..."
}

func main() {

	// Polymorpyhsm
	hewanList := []Animal{
		&Singa{Hewan{name: "Widodo", spesies: "Singa", age: 69}},
		&Gajah{Hewan{name: "Unis", spesies: "Gajah", age: 46}},
		&Monyet{Hewan{name: "Abil", spesies: "Monye", age: 99}},
	}

	for _, animal := range hewanList {
		fmt.Printf("Nama: %s, Spesies: %s, Umur: %d, Suara: %s\n",
			animal.GetName(), animal.GetSpesies(), animal.GetAge(), animal.MakeSound())
	}

}

//class hewan terdiri dari nama, spesies, usia
// method suara
// subsclass untuk singa gajah monyet
// masing kelas mewarisi
// meng implemen suara hewan
// buat suatu fungsi yang memanggil method2 suara dari suatu object
