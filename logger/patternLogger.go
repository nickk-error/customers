package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	PermFileMode   os.FileMode = 0666
	TimeFormat     string      = "2006-01-02T15:04:05.000Z0700"
	MaxStackLength int         = 40
)

func (p *PatternLogger) AllowLogging(level LogLevel) bool {
	var isAllow bool

	if p.Level == LEVEL_ALL {
		isAllow = true
	} else if p.Level == LEVEL_OFF {
		isAllow = false
	} else {
		if p.Level.Integer() >= level.Integer() {
			isAllow = true
		}
	}

	return isAllow
}

func getCallersFrames(skipNoOfStack int) *runtime.Frames {
	stackBuf := make([]uintptr, MaxStackLength)
	length := runtime.Callers(skipNoOfStack, stackBuf[:])
	stack := stackBuf[:length]

	return runtime.CallersFrames(stack)
}

func getStackTraceError(err error) string {
	frames := getCallersFrames(5)
	trace := err.Error()

	for {
		frame, more := frames.Next()
		if !strings.Contains(frame.File, "runtime/") {
			trace = trace + fmt.Sprintf("\n File: %s, Line: %d, Func: %s", frame.File, frame.Line, frame.Function)
		}
		if !more {
			break
		}
	}

	return trace
}

func (p *PatternLogger) logApp(correlationID string, level LogLevel, message string, stackTrace string) {
	if !p.AllowLogging(level) {
		return
	}

	var messageBean = LogAppMessageBean{}
	messageBean.Timestamp = time.Now().Format(TimeFormat)
	messageBean.ApplicationName = p.ApplicationName
	messageBean.LogType = AppLog
	messageBean.CorrelationID = correlationID
	messageBean.Level = level
	messageBean.Message = message

	// jsonBinary, _ := json.Marshal(&messageBean)
	// p.writeLogToFile(string(jsonBinary))
	// fmt.Println(string(jsonBinary))

	if p.SetLogger.IsJSON {
		jsonBinary, _ := json.Marshal(&messageBean)

		if p.SetLogger.WriteFile {
			p.writeLogToFile(string(jsonBinary))
		} else {
			fmt.Println(string(jsonBinary))
		}
	}
}

func (p *PatternLogger) WriteRequestMsg(correlationID string, url string, httpMethod string, message interface{}) {
	if message == nil {
		p.logApp(correlationID, LEVEL_INFO, "Request URL: "+url+
			", HttpMethod: "+httpMethod+", Request Message: null", "")
	} else {
		msg, err := json.Marshal(message)
		if err != nil {
			p.Error(correlationID, err.Error(), err)
		}
		p.logApp(correlationID, LEVEL_INFO, "Request URL: "+url+
			", HttpMethod: "+httpMethod+", Request Message: "+string(msg), "")
	}
}

func (p *PatternLogger) WriteResponseMsg(correlationID string, message interface{}) {
	if message == nil {
		p.logApp(correlationID, LEVEL_INFO, "Response Message: null", "")
	} else {
		msg, err := json.Marshal(message)
		if err != nil {
			p.Error(correlationID, err.Error(), err)
		}
		p.logApp(correlationID, LEVEL_INFO, "Response Message: "+string(msg), "")
	}
}

func (p *PatternLogger) Info(correlationID string, message string, args ...interface{}) {
	if len(args) > 0 {
		msg, stackTrace := checkArguments(args)
		p.logApp(correlationID, LEVEL_INFO, message+msg, stackTrace)
	} else {
		p.logApp(correlationID, LEVEL_INFO, message, "")
	}
}

func (p *PatternLogger) Fatal(correlationID string, message string, args ...interface{}) {
	if len(args) > 0 {
		msg, stackTrace := checkArguments(args)
		p.logApp(correlationID, LEVEL_FATAL, message+msg, stackTrace)
	} else {
		p.logApp(correlationID, LEVEL_FATAL, message, "")
	}
}

func (p *PatternLogger) Error(correlationID string, message string, args ...interface{}) {
	if len(args) > 0 {
		msg, stackTrace := checkArguments(args)
		p.logApp(correlationID, LEVEL_ERROR, message+msg, stackTrace)
	} else {
		p.logApp(correlationID, LEVEL_ERROR, message, "")
	}
}

func (p *PatternLogger) Warn(correlationID string, message string, args ...interface{}) {
	if len(args) > 0 {
		msg, stackTrace := checkArguments(args)
		p.logApp(correlationID, LEVEL_WARN, message+msg, stackTrace)
	} else {
		p.logApp(correlationID, LEVEL_WARN, message, "")
	}
}

func (p *PatternLogger) Debug(correlationID string, message string, args ...interface{}) {
	if len(args) > 0 {
		msg, stackTrace := checkArguments(args)
		p.logApp(correlationID, LEVEL_DEBUG, message+msg, stackTrace)
	} else {
		p.logApp(correlationID, LEVEL_DEBUG, message, "")
	}
}

func checkArguments(args []interface{}) (message string, staceTrace string) {
	for i := 0; i < len(args); i++ {
		switch v := args[i].(type) {
		case float32, float64, complex64, complex128:
			message = message + " " + fmt.Sprintf("%g", v)
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			message = message + " " + fmt.Sprintf("%d", v)
		case bool:
			message = message + " " + strconv.FormatBool(v)
		case string:
			message = message + " " + v
		case error:
			staceTrace = staceTrace + " " + getStackTraceError(v)
		case *strconv.NumError:
			staceTrace = staceTrace + " " + getStackTraceError(v)
		default:
			str, _ := json.Marshal(v)
			message = message + " " + string(str)
		}
	}
	return
}

func (p *PatternLogger) writeLogToFile(message string) {
	fileName := p.SetLogger.Path + "/" + p.SetLogger.FileName + "_" + time.Now().Format("2006-01-02") + ".log"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, PermFileMode)

	if err != nil {
		fmt.Printf("Opening log file error: %v", err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.SetFlags(log.LstdFlags)
	log.Println(message)
}
