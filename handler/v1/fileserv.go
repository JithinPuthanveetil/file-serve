package v1

import (
	"net/http"
	"path"
	"strconv"

	"github.com/file-serve/handler/config"
	"github.com/file-serve/handler/helper"
	"github.com/file-serve/model"
	"github.com/file-serve/usecase/fileserv"
	"github.com/go-chi/chi"
)

// Route traffic
func Route() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/*", showFile)
	return r
}

func showFile(w http.ResponseWriter, r *http.Request) {
	var (
		isDownload bool
		res        *model.Server
		err        error
	)

	reqParams := r.URL.Query()
	if reqParams.Get("download") != "" {
		isDownload, err = strconv.ParseBool(reqParams.Get("download"))
		if err != nil {
			helper.Fail(w, http.StatusBadRequest, err)
			return
		}
	}

	req := fileserv.ShowFileRequest{
		Root: config.Root + path.Clean(r.URL.Path),
	}

	if isDownload {
		res, err = req.DownloadFile()
		if err != nil {
			helper.Fail(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		res, err = req.ShowFile()
		if err != nil {
			helper.Fail(w, http.StatusInternalServerError, err)
			return
		}
	}

	helper.Send(w, http.StatusOK, res)
}
