package entity

//定义用户结构体
type User struct {
	LId          int64  `db:"lId" json:"lId"`
	StrName      string `db:"strName" json:"strName"`
	StrPwd       string `db:"strPwd" json:"-"` //json格式化的时候不返回
	StrRealName  string `db:"strRealName" json:"strRealName"`
	StrIdCardNo  string `db:"strIdCardNo" json:"strIdCardNo"`
	NAuthStatus  int    `db:"nAuthStatus" json:"nAuthStatus"`
	StrMobile    string `db:"strMobile" json:"strMobile"`
	StrEmail     string `db:"strEmail" json:"strEmail"`
	DtUpdateTime string `db:"dtUpdateTime" json:"dtUpdateTime"`
	DtCreateTime string `db:"dtCreateTime" json:"dtCreateTime"`
	NDeleted     int    `db:"nDeleted" json:"-"`
}

type UserData struct {
	Token string   `json:"token"`
	User  UserResp `json:"user"`
}

type UserResp struct {
	LId     int64  `json:"lId"`
	StrName string `json:"strName"`
}
