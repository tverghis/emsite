package handlers

import (
	"fmt"
	"net/http"

	"github.com/tverghis/emsite/util/files"
)

type Download struct{}

func NewDownload() *Download {
	return &Download{}
}

func (d *Download) GetDownload(w http.ResponseWriter, r *http.Request) {
	data, err := files.GetUploadsArchive()

	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		fmt.Printf("%s\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(data.Bytes())
}
