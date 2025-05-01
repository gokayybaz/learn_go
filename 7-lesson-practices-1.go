package main

import "fmt"

func lessonRepracticesOne() {
	fmt.Println("Hello World!")
    // Değişkenler
	var name string = "Gökay"
	var surName = "Baz"
	age := 21
	userName := "gokayybaz"
	salary := 24578.67
	workExperience := 24
    
	// Print (Konsola Yazdırma) ve Çeşitleri
	fmt.Print("Kullanıcı: ", name, " ")
	fmt.Print(surName, "\n")
    
	fmt.Println("Kullanıcı Yaşı:", age)

	age = age + 1
	
	// Tip Dönüşümü Yapma
	fmt.Println("Toplam Aldığı Para:", salary * float64(workExperience))

	fmt.Println("* Güncellenmiş * Kullanıcı Yaşı:", age)
    // String Formatlama
	userNameText := fmt.Sprintf("Kullanıcı Adı: %v ", userName)
	fmt.Println(userNameText)
	fmt.Printf("Kullanıcı Adı Veri Tipi: %T \n", userName)

	fmt.Println()

    // Kontrol Yapıları
	fmt.Println("* Kişi 18 yaşından küçük mü?" , age < 18)
	fmt.Println("* Kişi 18 yaşında mı?", age == 18)
	fmt.Println("* Kişi 18 yaşından büyük mü?", age > 18)
    
	// Koşul Yapıları
	if age < 18 {
		fmt.Println("- Bu kişi iş yerinde çalıştırılamaz!")
	}else if age == 18 {
		fmt.Println("- Bu kişi bundan sonra iş yerinde çalıştırılabilir!")
	}else {
		fmt.Println("- Bu kişi iş yerinde çalıştırılabilir!!")
	}

	fmt.Println()

	switch workExperience {
	case 0:
		fmt.Println("Daha yeni işe başladı")
	case 12:
		fmt.Println("Junior seviyede")
	case 24:
		fmt.Println("Junior to Middle seviyede")
	case 36:
		fmt.Println("Middle seviyede")
	case 48:
		fmt.Println("Middle to Senior seviyede")
	case 60:
		fmt.Println("Senior seviyede")
	default:
		fmt.Println("Seviye bilinmiyor!!!")
	}
    
	// Diziler
    var personalProjects [4]string = [4]string{
      "C# ile MVC kullanarak Blog Sitesi",
	  "C# ile WEB API kullanarak api",
	  "Go ile TCP Handler",
	  "React ile Food App",
	}

	fmt.Println(personalProjects)
    
	// Dizi Elemanlarına Erişme ve Manipüle Etme
	fmt.Println(personalProjects[0])
	personalProjects[0] = "Go ile MVC kullanarak Blog Sitesi"
	fmt.Println(personalProjects[0])
    
	// Slice
	var personalAbilities = []string {
		"C#",
		"Go",
		"ReactJS",
		"React Native",
		"Rest API",
	}

	fmt.Println(personalAbilities)

	// Slice Elemanlarına Ulaşma
	fmt.Println(personalAbilities[0])
	// Slice'a Eleman Ekleme
	personalAbilities = append(personalAbilities, "GraphQL API")
	fmt.Println(personalAbilities)
	// Slice'dan aralık dilim alarak slice'dan eleman silme
	personalAbilities = personalAbilities[0:4]
	fmt.Println(personalAbilities)

	fmt.Println()
    
	// String, Array ve Slice Uzunluklarını Öğrenme
	fmt.Println("Kullanıcı adı kelime uzunluğu:",len(userName))
	fmt.Println("Kullanıcı projeleri dizi uzunluğu:",len(personalProjects))
	fmt.Println("Kullanıcı yetenekleri slice uzunluğu:",len(personalAbilities))

}