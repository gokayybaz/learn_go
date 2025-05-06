package lessons

import "fmt"

func Loops() {
	fmt.Println("---- Loops (Döngüler) ----")
	// Döngüler tekrarlanan işlemleri programlamada kolayca yapmak için kullanılan yapılardır.

	// * temel for döngüsü tanımlama:
	x := 0
	for x <= 5 {
		fmt.Println("X'in Değeri:", x)
		x++
	}

	// * sonsuz for döngüsü tanımlama:
	/*
		for x > 0 {
			fmt.Println("Sonsuz Döngü:", x)
			x++
		}
	*/

	// * tek satırda for döngüsü tanımlama:
	for i := 0; i < 50; i++ {
		fmt.Println("i'nin değeri:", i)
	}

	// * for döngüsü ile slice array'leri beraber kullanma:
	// -> diziler birden fazla veriyi içerisinde barındırdığından dolayı içerisindeki verilere tek tek erişmemiz gerekebilir.
	// -> burada da tekrarlanan bir işlem olacağı için biz bu işi döngüler ile yapıp, dizi içerisindeki elemanlara tek tek erişip üzerinde işlem yaptırabiliyoruz.
	students := []string{
		"Gökay",
		"Mustafa",
		"Ali",
		"Zeynep",
	}

	for y := 0; y < len(students); y++ {
		fmt.Println("Öğrenci:", students[y])
	}

	// -> for: for döngüsü tanımlıyoruz.
	// -> y:=0: for döngüsünün başlangıç değerini belirtiyoruz.
	// -> y < len(students): for döngüsünün koşulunu belirtiyoruz; başlangıç değeri, dizimizin uzunluğundan küçük olduğu sürece çalışacaktır.
	// -> y++: her döngü işleminde başlangıç değerini arttırıyoruz.

	// Böylece dizi içerisindeki elemanlara tek tek ulaşabiliyor, elemanlar üzerinde istediğimiz işlemleri yapabiliyoruz.

	// * diziler için daha kolay ve daha kullanışlı bir for döngü çeşidi daha var:
	for index, value := range students {
		fmt.Printf("%v - %v \n", index, value)
	}
	// -> range ile bir diziyi kolayca dönebiliyoruz, bu da bize dizi elemanlarına çok daha kolay ulaşmamızı sağlıyor.
	// -> range bize index ve value değerlerini veriyor, burada index dizi elemanının indexi, value ise o anki index'e sahip gezilen elemanı veriyor.
	// -> range'in bize verdiği index'e ihtiyacımız yoksa; "_" katarak Go'ya biz bunu kullanmayacağız vermene gerek yok diyoruz, böylece Go bize kızmıyor.
	// -> range yapısı diğer programlama dillerinde bulunan forEach yapısına oldukça benzemektedir.

	// Döngülerde tanımlanan değişkenler (i, index, value gibi.) sadece döngü kod bloğunun içerisinde erişilebilir. Başka bir yerde erişilemez!
}
