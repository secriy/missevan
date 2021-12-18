package missevan

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	URLPhoneLogin = "https://app.missevan.com/member/login"
)

func postRequest(_url, payload string) (res []byte, err error) {
	method := "POST"
	client := &http.Client{}
	req, err := http.NewRequest(method, _url, strings.NewReader(payload))

	if err != nil {
		return
	}

	equipID := "52e76e9a-7592-43af-ae72-8ec755094e41"
	date := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	// equipID := "52e76e9a-7592-43af-ae72-8ec755094e41"
	// date := "2021-12-17T13:05:07.290Z"
	req.Header.Add("Host", "app.missevan.com")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "")
	req.Header.Add("Accept-Language", "zh-Hans-CN;q=1, en-CN;q=0.9, ja-CN;q=0.8")
	req.Header.Add("User-Agent", "MissEvanApp/4.7.4 (iOS;15.1;iPhone11,6)")
	req.Header.Add("X-M-Date", date)
	req.Header.Add("X-M-Nonce", "247b7772-63e4-4185-b774-17d5ecf821e6")
	req.Header.Add("Cookie", BaseCookie())
	req.Header.Add("Cookie", "equip_id="+equipID)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	Authorization(req.Header, _url, method, payload, equipID)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return body, nil
}
