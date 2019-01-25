package testInterface

import (
	"common"
	"strconv"
)

// 任务中心
func TaskList() {
	requestUrl := "/api/machine/task/list"
	session := MachineLogin()
	parm := map[string]string{
		"patientId":   "2216",
		"pageNo":      "1",
		"pageSize":    "10",
		"tokenId":     session.Data.TokenId,
		"instituteId": strconv.FormatInt(session.Data.InstituteId, 10),
	}
	common.HttpPost(requestUrl, parm)
}

// 检测数据上传
func UploadMedicalData(){
	requestUrl := "/api/machine/medicalReport/uploadMedicalData"
	session := MachineLogin()
	parm := map[string]string{
	"appKey":"888888",
		"versionNo":"225",
		"lym":"12.7",
		"timestamp":"1547712291929",
		"patientId":"2249",
		"sysType":"2",
		"deviceToken":"M0000",
		"appVersion":"2.2.5",
		"appType":"1",
		"eos":"2.2",
		"tokenId":session.Data.TokenId,
		"assbxb":"4.0",
		"neu":"6.2",
	     "mou":"5.6",
		"dataSource":"0",
		"deviceNo":"5161c7da8a773da5",
		"osVersion":"5.1",
		"phoneModel":"f201_3565U",
		"mon":"11.8",
		"bas":"4.0",
		"dataId":"",
		"deviceCode":"M0000",
		"hscrp":"3.2",
}
	common.HttpPost(requestUrl, parm)
}