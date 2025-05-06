package lessons

import "fmt"

func Variables() {
	fmt.Println("---- Variables (Değişkenler) ----")

	var name string = "Gökay"
	var surName = "Baz"
	age := 21
	var isActive bool = true
	length := 1.67

	fmt.Println("Kullanıcı Adı:", name)
	fmt.Println("Kullanıcı Soyadı:", surName)
	fmt.Println("Kullanıcı Yaşı:", age)
	fmt.Println("Kullanıcı Aktif mi?:", isActive)
	fmt.Println("Kullanıcı Boyu:", length)

	number1 := 58
	number2 := 77
	number3 := 89.876

	fmt.Println("Toplama:", number1+number2)
	fmt.Println("Çıkarma:", number2-number1)
	fmt.Println("Çarpma:", number1*number2)
	fmt.Println("Bölme:", number2/number1)

	// Farklı veri türlerine sahip değerler ile işlem yapma
	fmt.Println("Integer ile Float Toplama:", number1+int(number3))
	// Sonuç integer yani tam sayı olacak.

	fmt.Println("Float ile Integer Toplama:", float64(number1)+number3)
	// Sonuç float yani ondalıklı sayı olacak.
	// Burada sonucu hangi tipte görmek istiyorsak, dönüşümü ona göre yapmalıyız.

	// ---------------------------------------------------------------------------
}
