package controller

import (
	R "../../YdStruct/base"
	T "../../YdStruct/entity"
	"../../YdWork/network"
	Util "../../YdWork/util"
	"../service"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	hc := network.GetHttpClient(w, r)
	params := hc.GetParam()

	strName := params.Get("strName")
	strPwd := params.Get("strPwd")
	strMac := params.Get("strMac")

	if strName == "" {
		hc.ReturnMsg(R.ErrorMsg("请输入用户名"))
		return
	}
	if strPwd == "" {
		hc.ReturnMsg(R.ErrorMsg("请输入密码"))
		return
	}
	user, count := service.GetUserByName(strName)
	if count <= 0 {
		hc.ReturnMsg(R.ErrorMsg("用户不存在"))
		return
	}
	//校验密码
	pwd := Util.GetMd5(Util.DesEncode(strPwd))
	if pwd != user.StrPwd {
		hc.ReturnMsg(R.ErrorMsg("密码错误"))
		return
	}
	//生成token
	token := Util.CreateToken(string(user.LId), strMac)

	userData := T.UserData{}
	userData.User.LId = user.LId
	userData.User.StrName = user.StrName
	userData.Token = token
	hc.ReturnMsg(R.OK().SetData(userData))
}

func Register(w http.ResponseWriter, r *http.Request) {
	hc := network.GetHttpClient(w, r)
	params := hc.GetParam()

	strName := params.Get("strName")
	strPwd := params.Get("strPwd")
	strMac := params.Get("strMac")

	if strName == "" {
		hc.ReturnMsg(R.ErrorMsg("请输入用户名"))
		return
	}
	if strPwd == "" {
		hc.ReturnMsg(R.ErrorMsg("请输入密码"))
		return
	}
	_, count := service.GetUserByName(strName)
	if count > 0 {
		hc.ReturnMsg(R.ErrorMsg("用户已存在"))
		return
	}
	user := T.User{}
	user.StrName = strName
	user.StrPwd = Util.GetMd5(Util.DesEncode(strPwd))
	ret := service.SaveUser(user)
	if ret <= 0 {
		hc.ReturnMsg(R.ErrorMsg("注册失败，请重新注册"))
		return
	}
	//校验密码

	//生成token
	token := Util.CreateToken(strconv.FormatInt(user.LId, 10), strMac)

	userData := T.UserData{}
	userData.User.LId = user.LId
	userData.User.StrName = user.StrName
	userData.Token = token
	hc.ReturnMsg(R.OK().SetData(userData))
}
