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
Path string
Name string

}

func (f *NotExistFile) Error() string {
if err != nil {
log.Fatalf("FilePath Error Description -> %v ,FileName: %v , FilePath: %v",f.ErrorDescription,f.Name,f.Path)
}
return nil
}

//Sistem durması!
func systemExit(err error) string  {
if err != nil {
log.Panicf("%v \n %v","Program beklenmedik şekilde durdu!",time.Now())
}
return nil
}

//Fatura oluşturma
func DoesNotCreateBill(err error) string  {
if err != nil {
fmt.Printf("%v","Fatura oluşturulurken hata oluştu!")
}
return nil
}

//Veritabanına bağlantı
func DbConnectionError(err error) string  {
if err != nil {
fmt.Printf("%v", "Veritabanına bağlanılamadı!")
}
return nil
}

//Fatura Silinmesi
func DeleteBillError(err error) string  {
if err != nil {
log.Panicf("%v", "Fatura silinirken bir hata oluştu!")
}
return nil
}










