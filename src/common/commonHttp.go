package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"
)

// 请求地址前缀
const baseUrl string = "http://192.168.18.152:8282"

func HttpPost(requestUrl string, parm map[string]string) string {
	start := time.Now()
	contentType := "application/json;charset=utf-8"
	// 地址
	urlPath := baseUrl + requestUrl

	// 公共参数
	basePar := map[string]string{
		"appKey":      "888888",
		"appType":     "2",
		"appVersion":  "1.5.2",
		"deviceNo":    "andrid_device",
		"deviceToken": "token_999",
		"osVersion":   "6.0",
		"phoneModel":  "Android SDK built for x86",
		"sysType":     "2",
		"timestamp":   "1543976627469",
		"versionNo":   "1",
	}

	// 追加业务请求参数
	for k, v := range parm {
		basePar[k] = v
	}

	var keys []string
	for k := range basePar {
		keys = append(keys, k)
	}
	//排序key并且拼接sign字符串
	sort.Strings(keys)
	var signStr string
	for _, name := range keys {
		signStr += name + "=" + basePar[name]
	}

	b, err := json.Marshal(basePar)
	if err != nil {
		log.Println("json format error:", err)
		return err.Error()
	}

	body := bytes.NewBuffer(b)
	log.Println("接口请求参数", body)

	// 签名
	sign := CryptographyMd5(signStr + secret)
	// 请求接口
	resp, err := http.Post(urlPath+"?sign="+sign, contentType, body)

	if err != nil {
		log.Println("Post failed:", err)
		return err.Error()
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read failed:", err)
		return err.Error()
	}
	// 输出接口响应时间
	end := time.Now()
	delta := end.Sub(start)
	log.Printf(urlPath+" this amount of time: %s\n", delta)
	log.Println("请求接口", urlPath, " 返回content:", string(content))
	return string(content)
}

