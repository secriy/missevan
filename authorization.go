package missevan

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
)

const secret = "f8Yw1M3u0e5MV7z34PcbY9wgP56YwJ"

func Authorization(header http.Header, _url, method, body, equipID string) {
	str := strings.Builder{}
	// url = "https%3A//app.missevan.com/member/login"
	_url = strings.ReplaceAll(_url, ":", "%3A")
	str.WriteString(method + "\n")
	str.WriteString(_url + "\n")
	str.WriteString("" + "\n") // query string
	str.WriteString("equip_id:" + equipID + "\n")
	str.WriteString("x-m-date:" + header["X-M-Date"][0] + "\n")
	str.WriteString("x-m-nonce:" + header["X-M-Nonce"][0] + "\n")
	if method == "POST" {
		sum := sha256.Sum256([]byte(body))
		enc := base64.StdEncoding.EncodeToString(sum[:])
		str.WriteString(enc + "\n")
	}
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(str.String()))
	dst := h.Sum(nil)
	enc := base64.StdEncoding.EncodeToString(dst)
	header.Set("Authorization", "MissEvan "+enc)
}

func BaseCookie() string {
	_url := "https://fm.missevan.com/api/user/info"

	resp, err := http.Get(_url)
	if err != nil {
		return ""
	}
	cookie := strings.Builder{}
	for _, v := range resp.Header.Values("set-cookie") {
		cookie.WriteString(v)
	}
	return cookie.String()
}
