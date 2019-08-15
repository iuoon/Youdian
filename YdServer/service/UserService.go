package service

import (
	T "../../YdStruct/entity"
)

func GetUserByMobile(mobile string) (T.User, int) {
	user := T.User{}
	_ := "SELECT * FROM tb_user WHERE nDeleted=0 AND strMobile=? ORDER BY dtCreateTime DESC"
	return user, 1
}

func GetUserById(lId int64) (T.User, int) {
	user := T.User{}
	_ := "SELECT * FROM tb_user WHERE nDeleted=0 AND lId=? ORDER BY dtCreateTime DESC"

	return user, 1
}

func GetUserByName(name string) (T.User, int) {
	user := T.User{}
	_ := "SELECT * FROM tb_user WHERE nDeleted=0 AND strName=? ORDER BY dtCreateTime DESC"
	return user, 1
}

func SaveUser(user T.User) int {
	_ := "INSERT INTO tb_user (strName,strPwd,dtUpdateTime,dtCreateTime) VALUES (?,?,?,?)"

	return int(1)
}
