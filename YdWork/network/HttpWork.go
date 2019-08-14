package network

import (
	Base "../../YdStruct/base"
	"encoding/json"
	"net/http"
	"net/url"
)

type HttpClient struct {
	w http.ResponseWriter
	r *http.Request
}

func GetHttpClient(w http.ResponseWriter,r *http.Request) HttpClient {
	return HttpClient{w:w,r:r}
}

//获取请求参数
func (hc HttpClient) GetParam() url.Values{
	value := hc.r.Form
	return value
}

func (hc HttpClient) GetQueryParam(key string) string{
	value := hc.r.URL.Query().Get(key)
	return value
}

//返回消息
func (hc HttpClient)ReturnMsg(r Base.R)  {
	//tb.LogErr(json.NewEncoder(hc.w).Encode(r.OutLog()))
	json.NewEncoder(hc.w).Encode(r.OutLog())
	return
}
