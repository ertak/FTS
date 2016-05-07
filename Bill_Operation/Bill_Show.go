package Bill_Operation

import (
	"fmt"
	"html/template"
	"net/http"
)

var show_billtype string
var show_billMonth string

func Bill_Show(w http.ResponseWriter, r *http.Request) {
	fmt.Println("girdim-Show1")
	w.Header().Set("Content-Type", "text/html;text/css")
	//t, _ := template.ParseFiles("./web/showbill.html")
	template.Must(template.ParseFiles("./web/showbill.html")).Execute(w, nil)
	f := fmt.Println
	f(r.FormValue("buton"))
	/*show_billtype = r.FormValue("ShowBillType")
	show_billMonth = r.FormValue("ShowMonth")

	f("Fatura Tipi:", show_billtype)
	f("Fatura Kesim Ayı:", show_billMonth)
	*/
	//t.Execute(w, nil)
}
