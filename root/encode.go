package root

import "encoding/hex"

func Encode(str string) string {
	s, _ := hex.DecodeString(str)
	return string(s)
}
