package logger

type SetLogger struct {
	IsJSON    bool   `json:"isJson"`
	WriteFile bool   `json:"writeFile"`
	Path      string `json:"path"`
	FileName  string `json:"fileName"`
}

type PatternLogger struct {
	Level           LogLevel
	ApplicationName string
	SetLogger       SetLogger
}

type LogAppMessageBean struct {
	Timestamp       string   `json:"@timestamp"`
	ApplicationName string   `json:"appName"`
	LogType         LogType  `json:"logType"`
	CorrelationID   string   `json:"correlationID"`
	Level           LogLevel `json:"level"`
	Message         string   `json:"message,omitempty"`
}

type LogMonMessageBean struct {
	LogAppMessageBean
	TargetURL    string `json:"targetURL,omitempty"`
	Action       string `json:"action"`
	ElapsedTime  int64  `json:"elapsedTime"`
	ResponseCode string `json:"responseCode"`
}

type logTypes string

func (logType logTypes) isType() logTypes {
	return logType
}

type LogType interface {
	isType() logTypes
}

const (
	Monitor   = logTypes("Monitor")
	AppLog    = logTypes("AppLog")
	ServerLog = logTypes("ServerLog")
)

type CallerInfoBean struct {
	ClassName  string `json:"className"`
	MethodName string `json:"methodName"`
	FileName   string `json:"fileName"`
}

type LogLevel string

const (
	LEVEL_ALL   = LogLevel("ALL")
	LEVEL_TRACE = LogLevel("TRACE")
	LEVEL_DEBUG = LogLevel("DEBUG")
	LEVEL_INFO  = LogLevel("INFO")
	LEVEL_WARN  = LogLevel("WARN")
	LEVEL_ERROR = LogLevel("ERROR")
	LEVEL_FATAL = LogLevel("FATAL")
	LEVEL_OFF   = LogLevel("OFF")
)

func (l LogLevel) Integer() int {
	var level int

	switch l {
	case LEVEL_ALL:
		level = 7
	case LEVEL_TRACE:
		level = 6
	case LEVEL_DEBUG:
		level = 5
	case LEVEL_INFO:
		level = 4
	case LEVEL_WARN:
		level = 3
	case LEVEL_ERROR:
		level = 2
	case LEVEL_FATAL:
		level = 1
	case LEVEL_OFF:
		level = 0
	default:
		level = 0
	}

	return level
}
