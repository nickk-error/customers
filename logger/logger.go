package logger

func InitLogger(appName string) *PatternLogger {
	var logger = new(PatternLogger)
	logger.SetLogger.IsJSON = true
	logger.Level = LEVEL_INFO
	logger.ApplicationName = appName
	logger.SetLogger.WriteFile = false
	logger.SetLogger.Path = "fileLog"
	logger.SetLogger.FileName = "logApp"
	return logger
}

func InitUtilLogger(appName string) *PatternLogger {
	return InitLogger(appName)
}

// func InitInboundLogger(appName string, targetSys LogSystem) *PatternLogger {
// 	return InitLogger(appName, CrmInbound, targetSys)
// }

// func InitOutboundLogger(appName string, targetSys LogSystem) *PatternLogger {
// 	return InitLogger(appName, CrmOutbound, targetSys)
// }

// func InitValidationLogger(appName string, targetSys LogSystem) *PatternLogger {
// 	return InitLogger(appName, CrmValidation, targetSys)
// }

// func InitScheduleLogger(appName string, targetSys LogSystem) *PatternLogger {
// 	return InitLogger(appName, CrmSchedule, targetSys)
// }
