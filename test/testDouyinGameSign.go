package main

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

func main() {
	Signature := ""
	echostr := ""
	token := ""
	timestamp := ""
	nonce := ""
	msg := ""

	// 收到GET请求
	sortedString := make([]string, 0, 4)
	sortedString = append(sortedString, token)
	sortedString = append(sortedString, timestamp)
	sortedString = append(sortedString, nonce)
	sortedString = append(sortedString, msg)
	// 需要对这些参数按字符串自然大小进行排序
	sort.Strings(sortedString)
	// 使用SHA1算法
	h := sha1.New()
	h.Write([]byte(strings.Join(sortedString, "")))

	bs := h.Sum(nil)
	_signature := fmt.Sprintf("%x", bs)

	// signature 一致表示请求来源于 字节小程序服务端
	if _signature == Signature {
		println("signature verify success, return echostr: ", echostr)
	} else {
		println("signature verify failed")
	}
}
