package main

import "fmt"

func slice() {
	fmt.Println("---- Slice (Uzunluğu Değiştirilebilir Diziler) ----")

	// Slice ile Array'ler birbiri ile aynı işlemi görürler fakat Slice'larda dizinin boyutunu belirtmemize gerek yoktur, dizi boyutu değişebilir.

	users := []string{"Gökay", "Hakkı", "Benan"}
	fmt.Println(users)
	fmt.Println(len(users))

	// Slice'a eleman ekleme;
	users = append(users, "Zeynep")
	// append: ekleme yapmak için kullanılacak fonksiyondur.
	// parametre olarak hangi slice'a ekleme yapacaksak onu ve ekleyeceğimiz değeri/değerleri giriyoruz.
	// Not: append ile birden fazla slice'ı da birleştirebiliriz.
	fmt.Println(users)
	fmt.Println(len(users))

	// Slice'dan belirli veriyi veya veri aralıklarını seçme;
	//fmt.Println(users[0])   // Slice'da bulunan ilk index'e karşılık gelen değeri getirir.
	//fmt.Println(users[0:2]) // Slice'da bulunan index'i 0 ila 2 arasında olan değerleri getirir. (2. index'i dahil etmez!!)

	// Slice'dan veri çıkarmak için de aralık yönetimini kullanırız!
	users = users[0:2]
	fmt.Println(users)

	// Slice ve Array'lerin uzunluklarını öğrenme;
	fmt.Println(len(users))
	// len(): length yani uzunluktan gelmektedir.
	// len fonksiyonun içerisine parametre olarak uzunluğunu öğrenmek istediğimiz: string, array, slice gibi dizi olan veri türlerini gönderebiliriz.
	// Not: Array'lerde çok kullanılmaz çünkü zaten tanımlanırken uzunluğu belirlidir.
}
