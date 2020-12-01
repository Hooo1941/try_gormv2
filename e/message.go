package e

var MsgFlags = map[int]string{
	SUCCESS:       "成功",
	UNKNOWN_ERROR: "未知错误",

	CREATE_ERROR: "创建失败",
	DELETE_ERROR: "删除失败",
	UPDATE_ERROR: "更新失败",
	SELECT_ERROR: "查询失败",

	USERNAME_OR_PASSWORD_ERROR:   "用户名或密码错误",
	USER_PASSWORD_NOT_CONSISTENT: "两次输入的密码不匹配",
	USERNAME_ALREADY_EXISTED:     "用户名已存在",
	USER_SET_PASSWORD_ERROR:      "设置密码失败",
	REGISTER_ERROR:               "注册失败",
	USERNAME_TOO_SHORT:           "用户名过短",
	USERNAME_TOO_LONG:            "用户名过长",
	PASSWORD_TOO_SHORT:           "密码过短",
	PASSWORD_TOO_LONG:            "密码过长",
	INPUT_EMPTY:                  "输入为空",
	CONTENT_TOO_SHORT:            "内容过短",
	ALREADY_STAR:                 "收藏过了",
	NO_STAR:                      "没有收藏",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[UNKNOWN_ERROR]
}
