package live

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestSafeMessage(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want string
	}{
		{"test1", "淘宝", "淘\u200b宝"},
		{"test2", "口袋妖怪", "口\u200b袋\u200b妖\u200b怪"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SafeMessage(tt.msg); got != tt.want {
				t.Errorf("SafeMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBrotliDecompress(t *testing.T) {
	bf := func(h string) []byte {
		b, _ := hex.DecodeString(h)
		return b
	}

	tests := []struct {
		name    string
		data    []byte
		want    []byte
		wantErr bool
	}{
		{"test1", bf("016700001B66000004BEBA7DEA2F6FC0DC70B61545D17EFAC5800D38601E788C749879B7C17148C75F31BAF3E5EB02E68402A22B15B091FFACB371F61CD922A642F83264E0965104FAC643FCE7C3F53D"), []byte(`{"type":"room","event":"statistics","room_id":868782449,"statistics":{"score":2840,"online":7,"vip":1}}`), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BrotliDecompress(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("BrotliDecompress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BrotliDecompress() got = %v, want %v", got, tt.want)
			}
		})
	}
}
