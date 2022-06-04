package missevan

import (
	"fmt"
	"net/url"
)

type LoginParam struct {
	Phone    string
	Password string
	Region   string
}

func NewLoginParam(phone string, password string, region string) *LoginParam {
	if region == "" {
		region = "CN"
	}
	password = url.QueryEscape(password)
	return &LoginParam{Phone: phone, Password: password, Region: region}
}

func (lp *LoginParam) Login() (res []byte, err error) {
	_url := URLPhoneLogin
	payload := fmt.Sprintf("account=%s&password=%s&region=%s", lp.Phone, lp.Password, lp.Region)
	res, err = postRequest(_url, payload)
	return
}
