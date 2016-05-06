package Bill_Operation
import (
	"net/http"
	"fmt"
)

var show_billtype string
var show_billMonth string

func Bill_Show(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("girdim-Show1")
	f := fmt.Println
	show_billtype = r.FormValue("ShowBillType")
	show_billMonth = r.FormValue("ShowMonth")

	f("Fatura Tipi:",show_billtype)
	f("Fatura Kesim Ayı:",show_billMonth)







}























