package main

import "fmt"

func main() {
	var (
		ay  string
		gun int
	)
	koc := "Burcunuz KOÇ...\nGünlük düzenlemeniz gereken alanlarda hızlı hareket etmemeye çalışın.\nFarkında olmadığınız iş kazaları yaşanabilir. Bu tür olaylara maruz kalmamak için yavaş hareket etmeli, heyecanınızı yönetmeyi öğrenmelisiniz."
	boga := "Burcunuz BOĞA...\nYalnızlık hissi sizleri köşeye sıkıştırıyor olabilir. Adım atmanız için manevi desteğe ihtiyaç duyuyorsunuz. Bu manevi desteği hiç beklemediğimiz bir kişiden görebilirsiniz.\n Hayatın sürprizlerine karşı hazır olun ve evrenin sizin için gönderdiği desteği kabul edin."
	ikizler := "Burcunuz İKİZLER...\nRuhsal gerginliğinizi atmak için kendinizi alışverişe yönlendirebilirsiniz. Duygularınızla hareket edeceksiniz. Duygularınız size biraz pahalıya patlayabilir. Lütfen böyle zamanlarda dışarıya değil içinize yönelin."
	yengec := "Burcunuz YENGEÇ...\nPartnerinizle duygusal konuşmalar yapacaksınız. İkili ilişkilerinizde geleceğinizi değerlendiriyorsunuz. Bu değerlendirme de karşılıklı düşüncelere önem vermeli ve ortak kararlar noktasından hareket etmelisiniz. Eğer kendi düşüncelerinizde ısrarcı olursanız ileri de o düşüncelerin altında kalabilirsiniz."
	aslan := "Burcunuz ASLAN...\nAşk hayatınız renkleniyor. Size karşı duygularını ifade etmek isteyenler olabilir. Duygusal olarak kendinizi tatminkar hissedebilir ve gelen istekleri olumlu değerlendirebilirsiniz. Anı yaşayın."
	basak := "Burcunuz BAŞAK...\nÖfkenizi kontrol etmekte zorlanacaksınız. İletişiminizde sivri bir dil takınarak yaralayıcı konuşmalar da bulunabilirsiniz. İçsel gerginliğinizi başkalarına göstermek zayıflığınızı vurgulayabilir. İçinizdeki problemleri kendinizle çözmeli karşı tarafa bunu aksettirmemelisiniz."
	terazi := "Burcunuz TERAZİ...\nPartnerinizle kavga ediyorsunuz. Ayrılma noktasına gelen konuları büyütmemelisiniz. Bu durum sizi uzun vadeli pişmanlıklara gebe bırakabilir. Lütfen kavga ederken ölçünüzü bilin. Kendinizi sonradan pişman olacağınız durumlar içine sokmayın."
	akrep := "Burcunuz AKREP...\nGüne yeni bir haberle uyanıyorsunuz. Harekete geçmenizi gerektirecek beklenmedik konu planlarınızda değişikliğe gitmenize yol açabilir. Sürpriz gelişmelere karşı hazırlıklı olun."
	yay := "Burcunuz YAY...\nÇalışma ortamınızda ekip arkadaşlarınıza karşı öfkeli bir iletişim dili kullanabilirsiniz. Kırıcı olmamaya çalışın. İlerleyen zamanlarda yaşadığınız bu sıkıntı size büyüyerek geri dönebilir."
	oglak := "Burcunuz OĞLAK...\nYeni tanıştığınız kişilerin referansı ile yeni iş girişimlerinde bulunabilirsiniz. Bu girişim sizi yoğun bir tempo içine "
	kova := "Burcunuz KOVA...\nMiras konularınızla ilgileniyorsunuz. Almak istediğiniz gayrimenkul için yüksek fiyat istenebilir. Bu doğrultuda bazı kararlar alabilir ve büyük borçlar altına girebilirsiniz. Doğru yatırım yapmak için enine boyuna iyice düşünüp harekete geçin."
	balik := "Burcunuz BALIK...\nÇevrenize karşı tahammül seviyeniz oldukça düşecektir. Sabrınızın son noktasında öfke patlaması yaşayabilirsiniz. Bu noktaya gelmemek için insanlarla görüşmeye biraz ara vermeniz iyi olabilir."
	fmt.Println("Hangi ayda doğdunuz?")
	fmt.Scan(&ay)
	fmt.Println("Ayın kaçında doğdunuz?")
	fmt.Scan(&gun)

	if ay == "mart" && gun > 21 || ay == "nisan" && gun < 20 {
		fmt.Println(koc)
	} else if ay == "nisan" && gun > 20 || ay == "mayıs" && gun < 21 {
		fmt.Println(boga)
	} else if ay == "mayıs" && gun > 21 || ay == "haziran" && gun < 23 {
		fmt.Println(ikizler)
	} else if ay == "haziran" && gun > 22 || ay == "temmuz" && gun < 23 {
		fmt.Println(yengec)
	} else if ay == "temmuz" && gun > 22 || ay == "ağustos" && gun < 23 {
		fmt.Println(aslan)
	} else if ay == "ağustos" && gun > 22 || ay == "eylül" && gun < 23 {
		fmt.Println(basak)
	} else if ay == "eylül" && gun > 22 || ay == "ekim" && gun < 23 {
		fmt.Println(terazi)
	} else if ay == "ekim" && gun > 22 || ay == "kasım" && gun < 22 {
		fmt.Println(akrep)
	} else if ay == "kasım" && gun > 22 || ay == "aralık" && gun < 22 {
		fmt.Println(yay)
	} else if ay == "aralık" && gun > 21 || ay == "ocak" && gun < 22 {
		fmt.Println(oglak)
	} else if ay == "ocak" && gun > 21 || ay == "şubat" && gun < 20 {
		fmt.Println(kova)
	} else if ay == "şubat" && gun > 19 || ay == "mart" && gun < 21 {
		fmt.Println(balik)
	}

}
