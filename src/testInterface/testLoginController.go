// 登录测试模块
package testInterface

import (
	"common"
	"encoding/json"
	"log"
)

// 登录session信息
type Session struct {
	Data         Data
	Code         string
	ErrorMessage string
}

type Data struct {
	TokenId     string
	DoctorCode  string
	Status      string
	InstituteId int64
}

const account string = "13926565962"
const password string = "123456"

// 一体机登录
func MachineLogin() Session {
	requestUrl := "/api/doctor/userAccount/login"
	password := common.CryptographyMd5(password)
	parm := map[string]string{
		"account":  account,
		"password": password,
	}
	content := common.HttpPost(requestUrl, parm)
	// 解析登录信息
	var session Session
	err := json.Unmarshal([]byte(content), &
		session)
	if err != nil {
		log.Println("登录接口结果json反序列化异常", err)
	}
	return session
}
