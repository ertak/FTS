package Error

import (
	"fmt"
	"log"
	"time"
)

//Gerekli Error tanımlamaları buraya eklenecektir.

//Belirtilen dosya bulunamadı! (Daha kompleks yapıya sahiptir.)
type NotExistFile struct {
	ErrorDescription string
	Path             string
	Name             string
}

func (f *NotExistFile) Error(err error) error {
	if err != nil {
		log.Fatalf("FilePath Error Description -> %v ,FileName: %v , FilePath: %v", f.ErrorDescription, f.Name, f.Path)
	}
	return nil
}

//Sistem durması!
func SystemExit(err error) error {
	if err != nil {
		log.Panicf("%v \n %v", "Program beklenmedik şekilde durdu!", time.Now())
	}
	return nil
}

//Fatura oluşturma
func DoesNotCreateBill(err error) error {
	if err != nil {
		fmt.Printf("%v", "Fatura oluşturulurken hata oluştu!")
	}
	return nil
}

//DB bağlantı hatası
func DbConnectionError(err error) error {
	if err != nil {
		fmt.Printf("%v", "Veritabanına bağlanılamadı!")
	}
	return nil
}

//Fatura Silinmesi
func DeleteBillError(err error) error {
	if err != nil {
		log.Panicf("%v", "Fatura silinirken bir hata oluştu!")
	}
	return nil
}

func WriteJsonFile(err error) error {
	if err != nil {
		log.Panicf("%v", "Json dosyasına yazılırken bir hata oluştu!")
	}
	return nil
}
