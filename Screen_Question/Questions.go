package Screen_Question

import "fmt"

var (
	Billtype, BillAmount          int
	BillMonth, BillDate, BillDesc string
)

func Question() {
	f := fmt.Println
	f("Lütfen Yapmak İstediğiniz İşlemi Seçiniz?")
	//Bu işlemler daha sonradan artabilir artması durumunda aşağıdaki if bloğu değiştirilmeli!
	f("Fatura Ekleme (1)")
	f("Fatura Silme (2)")

	var bill_opr_selected int
	fmt.Scanf("%d\n", &bill_opr_selected)

	f("İşlem Yapmak İstediğiniz Fatura İşlemini Seçiniz?")
	f("Elektrik Faturası	: (1)")
	f("Doğalgaz Faturası	: (2)")
	f("Su Faturası	: (3)")
	f("Telefon Faturası	: (4)")
	f("Diğer Faturası	: (5)")

	fmt.Scanf("%d\n", &Billtype)

	if bill_opr_selected == 1 {

		fmt.Println("Fatura Ayı:")
		fmt.Scanf("%s\n", &BillMonth)

		f("Tutar:")
		fmt.Scanf("%d\n", &BillAmount)

		f("Son Ödeme Tarihi:")
		fmt.Scanf("%s\n", &BillDate)

		f("Açıklama:")
		fmt.Scanf("%s\n", &BillDesc)

	} else {
		//Sildirme durumunda tüm bilgiler alınmayacak..
		f("Fatura Ayı:")
		fmt.Scanf("%s\n", &BillMonth)
	}
}
