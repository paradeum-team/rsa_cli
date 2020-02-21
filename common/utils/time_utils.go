package utils

import (
	"bfs_cli_rsa/common/dict"
	"bfs_cli_rsa/common/plogger"
	"fmt"
	"log"
	"time"
)

/**
 * yyyyMMddHHmmss 转换成时间对象
 */
func ParseTime14Len(fmtTime string)(time.Time,  error){
	if fmtTime==""{
		return time.Now(),fmt.Errorf("time is invalid . Can not be nil ")
	}
	fmtTemplate:=dict.SysTimeFmt4compact
	return ParseTime(fmtTime,fmtTemplate)
}

func ParseTime(fmtTime string,timeFormatTemplate string )(time.Time,  error){
	stamp, err := time.ParseInLocation(timeFormatTemplate, fmtTime, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	//stamp, err := time.ParseInLocation("20060102150405", "20190108135030", time.Local)
	log.Println(stamp.Unix())  //输出：1546926630

	return stamp,err
}

func CurrentTime()string{
	//timeStr:=time.Now().Format("2006-01-02 15:04:05")  //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	timeStr:=time.Now().Format(dict.SysTimeFmt4compact)  //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	log.Println(timeStr)    //打印结果：2017-04-11 13:24:04
	return timeStr
}

func DevTimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	plogger.NewInstance().GetLogger().Infof("[DEV]TIMETRACK,command: %s took %s\n", name, elapsed)
}