package lessons

import "fmt"

func Arrays() {
	fmt.Println("---- Arrays (Diziler) ----")

	// Arrays
	students := [3]string{
		"Gökay",
		"Mustafa",
		"Ahmet",
	}
	// [3]: dizinin içerisinde ne kadar eleman olacağını söylüyoruz.
	// string: dizinin içerisinde bulunacak elemanların veri tipini belirtiyoruz.
	// {}: dizi içerisine veri eklemek için kod bloğu açıyoruz, içerisine belirttiğimiz veri tipinde elemanlar ekliyoruz.

	users := [3]string{}
	// İçerisine eleman eklemeden de bir dizi oluşturabiliyoruz.
	users[0] = "Gökay"
	users[1] = "Hakkı"
	users[2] = "Benan"
	// Oluşturduğumuz diziye sonradan eleman eklemesi yapabiliyoruz.
	fmt.Println(users)

	fmt.Println(students)    // diziyi direk konsola yazdırdık.
	fmt.Println(students[0]) // dizinin ilk elemanını konsola yazdırdık.
	// Not : Go programlama dilinde dizi eleman index'i 0'dan başlamaktadır.
	fmt.Println(students[1]) // dizinin ikinci elemanını konsola yazdırdık.
	fmt.Println(students[2]) // dizinin üçüncü elemanını konsola yazdırdık.

	// Aynı şekilde array'in elemanlarına erişip değerini değiştirebiliriz;
	fmt.Println("Değişmeden önceki değer:", students[1])
	students[1] = "Cengiz"
	fmt.Println("Değiştirdikten sonraki değer:", students[1])

	// Array içerisine sonradan eleman ekleyip çıkarma işlemi yapamayız!!!
	// Çünkü sabit bir uzunlukları var ve sonradan değiştirilemez!!!
}
