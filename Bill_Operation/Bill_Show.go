package Bill_Operation

import (
	"fmt"
	"html/template"
	"net/http"
)


func Bill_Show(w http.ResponseWriter, r *http.Request) {
	fmt.Println("girdim-Show1")
	w.Header().Set("Content-Type", "text/html;text/css")
	template.Must(template.ParseFiles("./web/showbill.html")).Execute(w, nil)

}
