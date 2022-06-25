package live

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io/ioutil"
	"strings"

	"github.com/andybalholm/brotli"
	"github.com/google/uuid"
)

// MessageID used to generate a unique message ID for each message
// that is platform-specific, only used in MissEvan.
func MessageID() string {
	u := uuid.NewString()
	return "3" + u[1:]
}

// SafeMessage is used to add zero-width space between each character
// that avoid being blocked.
func SafeMessage(msg string) string {
	return strings.Join(strings.Split(msg, ""), "\u200B")
}

// BrotliDecompress decompress the brotli data that received from the WebSocket.
func BrotliDecompress(data []byte) ([]byte, error) {
	if len(data) < 4 {
		return nil, errors.New("data is too short")
	}
	if data[0] == 0x01 {
		r := brotli.NewReader(bytes.NewReader(data[4:]))
		b, err := ioutil.ReadAll(r)
		if err != nil {
			return nil, err
		}
		// Calculate the length of the data.
		length := binary.LittleEndian.Uint32(append(data[1:4], 0))
		if len(b) != int(length) {
			return nil, errors.New("data length is incorrect")
		}
		return b, nil
	} else if data[0] == 0x7B || data[0] == 0x5B {
		// If data[0] is 0x7B ('{') or 0x5B ('['), it is a JSON object.
		return data, nil
	}
	return nil, errors.New("unknown data")
}
