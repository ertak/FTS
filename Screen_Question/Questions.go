package Screen_Question
import "fmt"

var Billtype int
var BillMonth string
var BillAmount int
var BillDate string
var BillDesc string


func Question()  {
	fmt.Println("Lütfen Yapmak İstediğiniz İşlemi Seçiniz?")
	//Bu işlemler daha sonradan artabilir artması durumunda aşağıdaki if bloğu değiştirilmeli!
	fmt.Println("Fatura Ekleme (1)")
	fmt.Println("Fatura Silme (2)")

	var bill_opr_selected int
	fmt.Scanf("%d\n",&bill_opr_selected)


	fmt.Println("İşlem Yapmak İstediğiniz Fatura İşlemini Seçiniz?")
	fmt.Println("Elektrik Faturası	: (1)")
	fmt.Println("Doğalgaz Faturası	: (2)")
	fmt.Println("Su Faturası	: (3)")
	fmt.Println("Telefon Faturası	: (4)")
	fmt.Println("Diğer Faturası	: (5)")

	fmt.Scanf("%d\n",&Billtype)

	if bill_opr_selected == 1 {

		fmt.Println("Fatura Ayı:")
		fmt.Scanf("%s\n",&BillMonth)

		fmt.Println("Tutar:")
		fmt.Scanf("%d\n",&BillAmount)

		fmt.Println("Son Ödeme Tarihi:")
		fmt.Scanf("%s\n",&BillDate)

		fmt.Println("Açıklama:")
		fmt.Scanf("%s\n",&BillDesc)

	}else  {
		//Sildirme durumunda tüm bilgiler alınmayacak..
		fmt.Println("Fatura Ayı:")
		fmt.Scanf("%s\n",&BillMonth)
	}
}
