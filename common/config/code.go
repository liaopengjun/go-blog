package config

type ResCode int64

const (
	SUCCESS        ResCode = 200
	ERROR          ResCode = 500
	INVALID_PARAMS ResCode = 400

	ERROR_EXIST_TAG         ResCode = 10001
	ERROR_NOT_EXIST_TAG     ResCode = 10002
	ERROR_NOT_EXIST_ARTICLE ResCode = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL    ResCode = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT ResCode = 20002
	ERROR_AUTH_TOKEN               ResCode = 20003
	ERROR_AUTH                     ResCode = 20004
	NOT_TOKEN                      ResCode = 20005
)

var MsgFlags = map[ResCode]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_NOT_EXIST_TAG:            "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	NOT_TOKEN:                      "请先登录",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}

func (c ResCode) Msg() string {
	msg, ok := MsgFlags[c]
	if !ok {
		return MsgFlags[ERROR]
	}
	return msg
}
