package err

type ResCode int

func (r ResCode) GetMsg() string {
	return codeMsgMap[r]
}

var codeMsgMap = map[ResCode]string{
	CodeSuccess: "success",
}
