package main

import "fmt"

func main() {
	
	fmt.Println("Hello world") //string harus petik 2
	fmt.Println(len("Hello World")) //counter huruf
	fmt.Println("Telo goreng"[1]) //huruf ke-n, return ASCII
	
	fmt.Println("Satu:",1) // print angka
	fmt.Println("Tiga koma lima:",3.5)
	
	//tipe data boolean, keyword nya adalah bool
	fmt.Println("ini tipe data bool:",true)

	//1 variabel unik, gabisa dipakai lagi u/ tipe data berbeda
	var name = "Jeffry Gunawan"
	fmt.Println(name)
	
	//declare pengganti var= bisa aa :=
	country := "Indonesia"
	fmt.Println(country)
	
	//ganti tipe data int yang lain
	var age int64 = 38
	fmt.Println(age)

	//declare banyak variable
	var (
		firstname = "Jeffry"
		lastname = "Gunawan"
	)
	fmt.Println(firstname)
	fmt.Println(lastname)
}