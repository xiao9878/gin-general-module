package err

// 系统级别的相应代码在此返回，10000开头
const (
	// 系统相关
	CodeSuccess             = 0
	DataBaseConnectionError = 10000 + iota
	DataBaseOperationError
	UserInfoNotExist
)
