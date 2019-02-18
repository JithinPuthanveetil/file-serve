package helper

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/file-serve/handler/config"
	"github.com/file-serve/model"
)

// Send writes success response
func Send(w http.ResponseWriter, status int, result *model.Server) {
	if result.Data == nil || len(result.Data) < 1 {
		t, err := template.New("show-file").Parse(config.Template)
		if err != nil {
			http.Error(w, "Error in parsing template", http.StatusInternalServerError)
			return
		}

		err = t.Execute(w, result)
		if err != nil {
			http.Error(w, "Error in executing template", http.StatusInternalServerError)
			return
		}
	}
	if result.IsDownload {
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", result.Root))
	}

	w.Write(result.Data)
}

// Fail writes failure response
func Fail(w http.ResponseWriter, status int, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
