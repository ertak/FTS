package Screen_Question

import (
	"fmt"
)

var (
	Billtype, BillAmount,Bill_opr_selected          int
	BillMonth, BillDate, BillDesc 					string
)

func Question() {
	f := fmt.Println
	f("Lütfen Yapmak İstediğiniz İşlemi Seçiniz?")
	//Bu işlemler daha sonradan artabilir artması durumunda aşağıdaki if bloğu değiştirilmeli!
	f("Fatura Ekleme (1)")
	f("Fatura Silme (2)")

	fmt.Scanf("%d\n", &Bill_opr_selected)

	for Bill_opr_selected !=1 && Bill_opr_selected !=2  {
		fmt.Println("Lütfen geçerli bir işlem numarası giriniz!")
		fmt.Scanf("%d\n", &Bill_opr_selected)
	}


	f("İşlem Yapmak İstediğiniz Fatura İşlemini Seçiniz?")
	f("Elektrik Faturası	: (1)")
	f("Doğalgaz Faturası	: (2)")
	f("Su Faturası	: (3)")
	f("Telefon Faturası	: (4)")
	f("Diğer Faturası	: (5)")

	fmt.Scanf("%d\n", &Billtype)

	if Bill_opr_selected == 1 {

		f("Fatura Dönemi:")
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
