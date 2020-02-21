package engine


type EngineResponse struct {
	Command     string
	CmdArgs     string
	Raw         []byte // 执行命令，返回内容
	Error       error
	SysCode     int
	IsTimeout   bool //是否超时
	IsBadParams bool //是否参数错误
	LocalFilePath string
	FileName string
}


type v2CmdResponse struct {
	Data    []byte
	Err     error
	SysCode int
}