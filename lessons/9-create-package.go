package lessons

// GO DİLİNDE PAKET YAPISI VE MODÜLERLİK

// - Go dilinde dosyalar "paket" (package) mantığı ile birbirinden ayrılır.
// - Java gibi dillerde de benzer bir yapı bulunur ve bu, kodun modüler ve okunabilir olmasını sağlar.
// - Go'da modülerlik sayesinde büyük projelerde clean code (temiz kod) prensiplerini uygulamak kolaylaşır.

// ANA DOSYA (package main)
// - Paket yapısı kullanılmadığı durumda, tüm Go dosyaları genellikle "package main" altında toplanır.
// - Bu şekilde yazılan kodlar çalışabilir, fakat yapı karmaşıklaştıkça yönetilmesi zor hale gelir.
// - Aynı klasörde yer alan ve "package main" içeren tüm `.go` dosyaları, `go run .` komutu ile birlikte çalıştırılabilir.
// - Ancak bu yapı, modülerlik sağlamadığı için büyük projelerde önerilmez.

// PAKET YAPISI NASIL OLUŞTURULUR?
// 1. `go mod init learn_go` komutu ile proje başlatılır.
//    - Bu komut `go.mod` adında bir yapılandırma dosyası oluşturur (şimdilik içeriğiyle ilgilenmemize gerek yok).
// 2. Yeni bir paket oluşturmak için proje klasörü içinde yeni bir klasör açılır. Örnek: `lessons`
// 3. Bu klasör içerisinde `.go` uzantılı dosyalar oluşturulur.
// 4. Bu dosyaların en üstüne `package lessons` şeklinde ilgili paket adı yazılır.
// 5. Böylece bu dosyalar, `lessons` isimli pakete ait olur.

// ÖNEMLİ NOT:
// - Bir pakette tanımlanan fonksiyonlara ve değişkenlere **dışarıdan erişebilmek için** isimlerinin **ilk harfi büyük** olmalıdır.
//   - Örnek: `func SayHello()`, `var Version = "1.0"`
// - Küçük harfle başlayan öğeler sadece aynı paket içinden erişilebilir (private).

// PAKETİ KULLANMA (main.go):
// 1. `main.go` dosyasına gelip oluşturduğumuz paketi `import` ifadesiyle dahil ederiz:
//      import "learn_go/lessons"
// 2. Artık `lessons` paketindeki fonksiyonlara şu şekilde erişebiliriz:
//      lessons.SayHello()
//      fmt.Println(lessons.Version)

// SONUÇ:
// - Böylece Go dilinde modern ve modüler bir proje yapısı kurmuş oluruz.
// - Kodları fonksiyonel olarak ayırarak okunabilirliği ve yeniden kullanılabilirliği artırırız.
// - Bu yapı, temiz kod (Clean Code) yazımında önemli bir adımdır.
