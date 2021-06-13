package common

// ReqParam 可选择传入的参数
type ReqParam struct {
	UserExtList []string	// 用户确认的视频后缀名支持列表
	SaveMultiSub bool		// 存储每个网站 Top1 的字幕
	DebugMode bool			// 调试标志位

	FoundExistSubFileThanSkip bool	// 如果视频的目录下面有字幕文件了，就跳过
	UseUnderDocker	bool	// 是否在 docker 下使用


	HttpProxy string		// HttpClient 相关
	UserAgent string		// HttpClient 相关
	Referer   string		// HttpClient 相关
	MediaType string		// HttpClient 相关
	Charset   string		// HttpClient 相关
	Topic	  int			// 搜索结果的时候，返回 Topic N 以内的
}
