package entity

//定义用户结构体
type User struct {
	Id         int64  `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Password   string `db:"password" json:"-"` //json格式化的时候不返回
	RealName   string `db:"realName" json:"realName"`
	IdCard     string `db:"idCard" json:"idCard"`
	AuthStatus int    `db:"authStatus" json:"authStatus"`
	Mobile     string `db:"mobile" json:"mobile"`
	Email      string `db:"email" json:"email"`
	UpdateTime string `db:"updateTime" json:"updateTime"`
	CreateTime string `db:"createTime" json:"createTime"`
	Deleted    int    `db:"deleted" json:"-"`
}

type UserData struct {
	Token string   `json:"token"`
	User  UserResp `json:"user"`
}

type UserResp struct {
	LId     int64  `json:"lId"`
	StrName string `json:"strName"`
}
