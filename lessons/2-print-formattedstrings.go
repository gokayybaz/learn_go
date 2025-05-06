package lessons

import "fmt"

func PrintAndFormattedStrings() {
	fmt.Println("---- Konsola Yazdırma & String Formatlama ----")

	name := "Gökay"
	age := 21
	var name2 string = "Mustafa"
	age2 := 56

	// Print: Satır atlaması yapmadan konsola yazı yazdırmak için kullanılır.
	//        Manuel olarak \n ile satır atlaması yaptırtabiliriz.
	fmt.Print("Hello ")
	fmt.Print("World \n")

	// Println: Satır atlaması yaparak konsola yazı yazdırmak için kullanılır.
	fmt.Println("Merhaba", name2, "Yaşınız:", age2)

	// Printf: String'i formatlayarak konsola yazı yazdırmak için kullanılır.
	fmt.Printf("Merhaba: %v , Yaşınız: %v \n", name, age)
	// fmt zaten format demektir.
	// Yani bu sayede string ifadelerin içerisine doğrudan değişkenleri enjekte edebiliriz.

	// "%v" : bütün tiplerde çalışır.

	// "%q" : sadece string tiplerinde çalışır ve string'leri çift tırnak içerisinde gösterir.

	// "%T" : değişkenin tipini görmek için kullanılır.

	// "%f" : Float yani ondalıklı sayılarda çalışır.
	//        Virgülden sonra belirli sayıda basamak bastırmak için: %0.1f şeklinde yazmalıyız.
	//        0.1: virgülden sonra sadece bir basamak göster/konsola yazdır.

	// Formatlanmış değeri değişkene atama işlemi
	price := 895.89
	var nameFormatted = fmt.Sprintf("Bu ürünün fiyatı: %0.2f₺", price)
	fmt.Println(nameFormatted)
}
