package root

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"github.com/auyer/steganography"
	"image"
	_ "io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

func Getimg(url string) string {  //解码img
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//http cookie接口
	cookieJar, _ := cookiejar.New(nil)
	c := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}
	resp, err := c.Get(url)
	if err != nil {
		fmt.Println("获取失败.")
	}
	defer resp.Body.Close()
	reader := bufio.NewReader(resp.Body)
	img, _, err := image.Decode(reader)
	sizeOfMessage := steganography.GetMessageSizeFromImage(img) // Uses the library to check the message size
	msg := steganography.Decode(sizeOfMessage, img)
	decodestr := string(msg)
	return decodestr
}
