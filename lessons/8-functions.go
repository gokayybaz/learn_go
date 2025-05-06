package lessons

import "fmt"

func Functions() {
	fmt.Println("---- Functions (Fonksiyonlar) ----")

	// * Fonksiyonlar belirli işleri yapan yapılardır. Fonksiyonlar programlamada tekrar eden işleri engellemek ve modülerliği sağlamak için sıkça kullanılır.
	// * Fonksiyonlar ile belirli işleri bir kez tanımlayıp, fonksiyonu istediğimiz kadar çağırarak ilgili işin çağırdığımız kadar yapılmasını sağlayabiliriz.

	// 2 - Fonksiyonları Çağırma (Kullanma)
	helloWorld()

	sayHello("Gökay")
	sayHello("Cahit")
	sayHello("Esat")

	area1 := calcArea(89.75)
	fmt.Println("Dairenin Alanı:", area1)
}

// 1 - Fonksiyon Tanımlama

// Geriye bir değer döndürmeyen sadece iş yapan fonksiyon;
// * Geriye değer döndürmeyen fonksiyonlar sadece içerisinde yazılan kodları çalıştırır yani yapması gereken işi yapar.
func helloWorld() {
	fmt.Println("Hello World!")
}

// Geriye bir değer döndürmeyen, aldığı parametreyi de kullanarak iş yapan fonksiyon;
// * Geriye değer döndürmeyen fonksiyonlar, yapması gereken işi yaparken dışarıdan da parametre alabilir. Aldığı parametreyi de kullanarak dinamik bir şekilde yapması gereken işi yapar.
// * GO dilinde fonksiyonun alacağı parametreyi belirlerken, parametre değerinin tipini de belirtmeliyiz. Aksi taktirde GO hata verecektir.
func sayHello(name string) {
	fmt.Println("Merhaba", name)
}

// Geriye bir değer döndüren, aldığı parametreyi de kullanarak iş yapan fonksiyon;
// * Geriye değer döndüren fonksiyonlar yapması gereken işi yaptıktan sonra oluşan değeri geri döndürür, dönen değeri biz istediğimiz gibi kullanabiliriz. (Değişkene atama, print etme vb.)
// * GO dilinde fonksiyon yazarken kod bloğundan önce dönecek değerin tipini belirtmeliyiz, Aksi taktirde GO hata verecektir.
func calcArea(r float64) float64 {
	return r * r * 3.14
	// return'den sonraki yazılan kodlar çalışmaz, bu yüzden return'ü fonksiyonun en sonuna almalıyız.
}
