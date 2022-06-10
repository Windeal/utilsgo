package log

// Config : 日志配置
type Config struct {
	Outputs []OutputConfig // 日志输出
}

// OutputConfig : 日志输入端配置
type OutputConfig struct {
	// name string

	// Writer 日志输出端 (console, file)
	Writer string
	// WriterConfig 日志输出端配置
	WriterConfig WriterConfig

	// Formatter 日志输出格式 (console, json)
	Formatter    string
	FormatConfig FormatConfig

	// Level 控制日志级别 debug info error
	Level string

	// CallerSkip 控制log调用栈需要跳过的层数
	CallerSkip int `yaml:"caller_skip"`
}

type WriterConfig struct {
	// LogPath 日志路径
	LogPath string `yaml:"log_path"`

	// Filename 日志文件名
	Filename string `yaml:"filename"`

	// MaxSize 日志文件的最大大小. 单位M
	MaxSize int `json:"maxsize" yaml:"maxsize"`

	// MaxAge 旧日志文件能够存在的最长时间。单位: 天
	MaxAge int `json:"maxage" yaml:"maxage"`

	// MaxBackups 保存的最大历史文件数量
	MaxBackups int `json:"maxbackups" yaml:"maxbackups"`
}

type FormatConfig struct {
}
