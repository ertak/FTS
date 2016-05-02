--Web Service nedir? 

*Request alýr kendi içinde deðerlendirir ve geriye cevap döner genel yapýsý ve iþleyiþi budur.
*HTTP üzerinden xml veya json ile data alýþveriþi saðlar.
*Go'da web servis kullanýmý için net/http paketi kullanýlýr.
*Baðýmsýz platformlarda çalýþabilir.
*Son zamanlarda daha çok kullanýlan yapýlar REST mimarisi ve SOAP protokolüdür.

--REST?(Representational State Transfer)

*HTTP tabanlýdýr.
*Bütün REST requestleri birer HTTP requesttir aslýnda ve serverlar bu istekleri farklý metodlarla handle etmektedir.
*HTTP gönderilir cevap olarak XML,Json hatta text alýnýr.
*Kullanýmýd daha rahat flexible'dýr.

--SOAP?(Simple Object Access Protocol)

*Çapraz network bilgi alýþveriþi için kullanýlýr.
*XML gönderilir response olarak XML alýnýr.
*Yapýsý karmaþýktýr bu eksi bir özelliktir.
*Go dilinde program yazarken REST kullanýlmalýdýr bunda temel sebep basit ve kullanýþlý olmasý Go kendi
yapýsý itibariyle simple bir yapýya sahiptir bu yüzden SOAP protokolünü kullanmak doðru bir karar deðildir.
*Microsoft tarafýndan geliþtirilmiþtir.
*Daha kurallý bir yapýya sahiptir.

--Socket: Serverlar HTTP nin düþük performansýnýn üstesinden gelebilmek için soketleri kullanmaktadýrlar. 

--Gorilla: Bir web toolkitidir ve net/http paketinden extend edilmiþtir ve performans açýsýndan iyidir.Gorilla gibi daha
birçok yapý bulunmaktadýr bunlardan biri tercih edilebilinir hepsinin çeþitli özellikleri ve eklemeleri bulunmaktadýr.
