package base64

import "encoding/base64"

func Decode(s string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
