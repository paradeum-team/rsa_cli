package plogger

import (
	"bfs_cli_rsa/common/dict"
	"github.com/kataras/golog"
	"log"
	"os"
	"time"
)

type PldLogger struct {
	logger      *golog.Logger
	currentDate string //当前时间
}

var pldLoggerInstance *PldLogger
func getCurrentDate()(dateLen8 string){
	currentDate :=time.Now().Format(dict.SysTimeFmt4compact)[:8]
	return currentDate
}

func NewInstance() *PldLogger {
	currentDate := getCurrentDate() //当前的8位长度的日期
	pldLoggerInstance = &PldLogger{
		logger: golog.Default,
		currentDate: currentDate,
	}
	pldLoggerInstance.logger.TimeFormat=dict.SysTimeFmt
	//level :=pldconf.AppConfig.Server.LogLevel
	level:="info"
	pldLoggerInstance.logger.SetLevel(level)//设置日志输出级别
	/**
	暂停日志输出到：落地磁盘

	baseLogPath := filepath.Join(pldconf.AppConfig.Server.APIDataFolder, dict.RN_LOG_FOLDER)
	//check base log dir .if not exits then create .
	createAfsLogDir(baseLogPath)

	logFileName := "rn_" + currentDate + ".log"
	logFilePath := filepath.Join(baseLogPath, logFileName)
	f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {

		log.Printf("ERROR: %s\n", fmt.Sprintf("%s append|create failed:%v", logFilePath, err))
		return nil
	}
	**/

	//设置output
	pldLoggerInstance.logger.SetOutput(os.Stdout)

	/**
	暂停日志输出到：落地磁盘
	 pldLoggerInstance.logger.AddOutput(f)

	 */
	return pldLoggerInstance
}



func (lf *PldLogger) GetLogger() *golog.Logger {
	if pldLoggerInstance == nil {
		NewInstance()
	}else{
		if lf.currentDate==getCurrentDate(){
			//同一天，说明日志不用切换文件，否则就新打开一个文件
		}else{
			NewInstance()
		}
	}
	return lf.logger
}

func createAfsLogDir(baseLogPath string) {

	err := os.MkdirAll(baseLogPath, os.ModePerm) //创建多级目录，如果path已经是一个目录，MkdirAll什么也不做，并返回nil。

	if err != nil {
		log.Println("ERROR: init log dir is something wrong ...")
		os.Exit(1) //日志文件目录创建不成功，则失败
	}

}
