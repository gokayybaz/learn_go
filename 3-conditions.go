package main

import "fmt"

func conditions() {

	fmt.Println("---- Conditions (Koşullar) ----")

	age := 24

	// Kontrol İşlemi
	fmt.Println(age > 10)  // age 10'dan büyük mü?
	fmt.Println(age >= 10) // age 10'dan büyük veya eşit mi?
	fmt.Println(age == 10) // age 10'a eşit mi?
	fmt.Println(age != 10) // age 10'a eşit değil?
	// Bir kontrol işlemi yaptığımız için bize mantıksal değer dönecek.
	// Yani: true/false dönecektir.

	// Koşul İşlemleri
	if age < 18 {
		fmt.Println("Kişi içeri giremez!")
	}
	if age < 18 {
		fmt.Println("Kişi reşit değil!")
	} else if age == 18 {
		fmt.Println("Kişi reşit oldu!")
	} else {
		fmt.Println("Kişi reşit!")
	}

	// if: eğer demektir.
	// else: değil ise demektir.
	// else if: değil ise eğer demektir.
	// Not: else if birden fazla koşulu kontrol etmek için kullanılır.

	dayIndex := 0

	switch dayIndex {
	case 0:
		fmt.Println("Bugün günlerden PAZARTESİ")
	case 1:
		fmt.Println("Bugün günlerden SALI")
	case 2:
		fmt.Println("Bugün günlerden ÇARŞAMBA")
	case 3:
		fmt.Println("Bugün günlerden PERŞEMBE")
	case 4:
		fmt.Println("Bugün günlerden CUMA")
	case 5:
		fmt.Println("Bugün günlerden CUMARTESİ")
	case 6:
		fmt.Println("Bugün günlerden PAZAR")
	default:
		fmt.Println("Girilen index'e eşleşen gün bulunmamaktadır!!!")
	}
}
