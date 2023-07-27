package main

import (
	"net/http"
	"os"
)

// transcode handlers the file transcoding API.
func (v VideoHandler) transcode(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := os.Setenv("FILE_TO_TRANSCODE", filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = v.ffmpeg.Transcode() // calls ffmpeg natively using the os env FILE_TO_TRANSCODE as filename to be used.
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
