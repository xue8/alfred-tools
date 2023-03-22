package url

import "net/url"

func Encode(s string) string {
	return url.QueryEscape(s)
}

func Decode(s string) (string, error) {
	return url.QueryUnescape(s)
}
