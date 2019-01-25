/**
	测试居民家庭关系
 */
package testInterface

import (
	"common"
	"strconv"
)

// 获取居民家庭列表
func GetFamilyMember() {
	requestUrl := "/machine/patient/familyMember/getFamilyMember"
	session := MachineLogin()
	parm := map[string]string{"patientId": "792137", "tokenId": session.Data.TokenId}
	common.HttpPost(requestUrl, parm)
}

// 绑定家庭关系-添加平台存在且未存在家庭关系的居民
func BindingPatientRelationInfo() {
	requestUrl := "/machine/patient/familyMember/bindingPatientRelationInfo"
	session := MachineLogin()
	parm := map[string]string{
		"patientId":      "2216",
		"tokenId":        session.Data.TokenId,
		"relation":       "兄弟",
		"familyMemberId": "2215",
	}
	common.HttpPost(requestUrl, parm)
}

// 解除家庭关系
func RelieveFamilyMember() {
	requestUrl := "/machine/patient/familyMember/relieveFamilyMember"
	session := MachineLogin()
	parm := map[string]string{
		"patientId": "2216",
		"tokenId":   session.Data.TokenId,
	}
	common.HttpPost(requestUrl, parm)
}

func AddFamilyMember() {
	requestUrl := "/machine/patient/familyMember/relieveFamilyMember"
	session := MachineLogin()
	parm := map[string]string{
		"patientId":     "2216",
		"institutionId": strconv.FormatInt(session.Data.InstituteId, 10),
		"relation":      "兄弟",
		"tokenId":       session.Data.TokenId,
	}
	common.HttpPost(requestUrl, parm)
}
