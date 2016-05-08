package Bill_Operation

import (
	"io"
	"net/http"
)

func Bill_Analysis(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Analiz yapım aşamasındadır.")
}
