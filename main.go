package main

import (
	"fmt"
	data "github.com/rizkaasta/project_module/v2"
)

type DataMember struct {
	nama string
}

func (member DataMember) Ucapan() {
	fmt.Println("Selamat datang kembali di Perpustakaan Daerah, ", member.nama)
}


func main() {
	fmt.Println("Ini adalah Perpustakaan Daerah")

	//Anynomous Function
	tolak := func(nama string) bool {
		return nama != "Rizka"
	}

	//Function sebagai parameter
	data.Anggota("Doni", tolak)

	//Anynomous Struct
	memberBaru := struct {
		nama string
		noAnggota int
		} { nama : "Dita",
		noAnggota: 2342,
		}
	fmt.Println("Selamat bergabung di Perpustakaan Daerah, ", memberBaru.nama)
	
	//struct method
	member1 := DataMember{nama: "Rizka"}
	member1.Ucapan()

	//struct
	data.TampilData()

	//Multiple Return Value
	var bayarDenda, bayarRusak = data.Denda(1,3,1)
	fmt.Printf("Anda telah melewati batas pengembalian buku dan melakukan perusakan buku dengan denda masing-masing sebesar %d dan %d\n", bayarDenda, bayarRusak)
	
	//Return Value
	var totalDenda = data.TotalDenda(bayarDenda, bayarRusak)
	fmt.Println("Segera kembalikan buku dan lakukan pembayaran denda Anda sebesar ", totalDenda)
}
	
		
	
	
