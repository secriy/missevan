package missevan

import (
	"testing"
)

func TestLoginParam_Login(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		lp := &LoginParam{
			Phone:    "",
			Password: "",
			Region:   "CN",
		}
		gotRes, err := lp.Login()
		if err != nil {
			t.Errorf("Login() error = %v, resp %v", err, gotRes)
			return
		}
		t.Logf("Login() success, resp %v", string(gotRes))
	})
}
