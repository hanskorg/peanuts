package tools

import (
	qrcode "github.com/skip2/go-qrcode"
	"net/http"
)

func Qrcode(resp http.ResponseWriter, req *http.Request) {
	var (
		info       string
		qBuff []byte
		err        error
	)
	info = req.URL.Query().Get("i")
	if info == "" {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("bad"))
		return
	}
	resp.Header().Set("Content-Type", "image/png")
    qBuff, err = qrcode.Encode(info, qrcode.Medium, 256)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("bad qrcode"))
		return
	}
	_, _ = resp.Write(qBuff)

	return

}
