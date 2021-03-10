package root

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	_ "image/jpeg"
)

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}
func Decrypt(decrypted string, key []byte) string { //des解密
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return ""
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return ""
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return ""
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	return string(out)
}

