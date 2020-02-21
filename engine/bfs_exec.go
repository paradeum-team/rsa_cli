package engine

import (
	"bfs_cli_rsa/common/dict"
	"bfs_cli_rsa/common/e"
	"bfs_cli_rsa/common/plogger"
	"fmt"
	"os/exec"
	"path/filepath"
	"time"
)

func AFSExec4V2(command, cmdArg string, optionalTimeout int) EngineResponse { //(raw string, isTimeout bool, err error) {

	engineResponse := EngineResponse{}

	defer devTimeTrack(time.Now(), "AFSExec::"+cmdArg)
	engineResponse.IsTimeout = false

	ABSAFSDir:="/data"
	AFSProgram:="afs-x64"
	afsClient := filepath.Join(ABSAFSDir, AFSProgram)
	plogger.NewInstance().GetLogger().Infof("AFSExec:%s %s \n", afsClient, cmdArg)

	cmd := exec.Command(afsClient, cmdArg)
	cmd.Dir = ABSAFSDir

	// cmdChan is the channel for command execution output and timeout signal.
	cmdChan := make(chan v2CmdResponse)
	go func(cmd *exec.Cmd, cChan chan v2CmdResponse) {
		out, err := cmd.CombinedOutput()
		cChan <- v2CmdResponse{out, err, -1}
	}(cmd, cmdChan)
	afsTimeoutHandler := v2TimeoutFunc(afsClient, cmdArg, cmdChan)
	timer := time.AfterFunc(time.Duration(optionalTimeout)*time.Second, afsTimeoutHandler)

	result := <-cmdChan
	timer.Stop()

	engineResponse = EngineResponse{Command: command, CmdArgs: cmdArg, Raw: result.Data, Error: result.Err, SysCode: result.SysCode}

	if result.SysCode == e.ScodeAFSTimeout {
		engineResponse.IsTimeout = true
		return engineResponse

	}

	if result.Err != nil {
		plogger.NewInstance().GetLogger().Errorf(fmt.Sprintf("running AFS command:%s %s, %v \n ", afsClient, cmdArg, result.Err))
		if result.SysCode == e.ScodeAFSTimeout {
			engineResponse.IsTimeout = true
		}

		engineResponse.IsBadParams = false

		if len(result.Data) > 0 {
			return engineResponse
		}

		return engineResponse
	}
	/**
		if command == dict.CMDIDSetParam {
			// Insert set param record to a database.
			var db *leveldb.DB
			db, err := leveldb.OpenFile(filepath.Join(pldconf.AppConfig.Server.APIDataFolder, pldconf.AppConfig.Server.DBFolder), nil)
			if err != nil {
				plogger.NewInstance().GetLogger().Errorf("v2AFSExec failed to open db for set param records %v", err)
				engineResponse.Error = err
				engineResponse.SysCode=e.ScodeErr //设置一个服务器错误号
				return engineResponse
			}
			defer db.Close()

			// Key     string  SET:YYYYMMDD:AFID
			// Value   int64   unixtimestamp
			afid := ""
			if afidIdx := strings.Index(cmdArg, ";afid="); afidIdx > -1 {
				//这个位置出问题了
				afid = cmdArg[afidIdx+len(";afid=") : afidIdx+len(";afid=")+dict.AFIdLength]
			}

			if afid != "" {
				ymd := time.Now().Local().Format("20060102150405")[:8]
				timeNowb := make([]byte, 8)
				binary.LittleEndian.PutUint64(timeNowb, uint64(time.Now().Unix()))
				db.Put([]byte(fmt.Sprintf(dict.SetParamRecordPrefixPlaceholder, ymd)+afid), timeNowb, nil)
			} else {
				plogger.NewInstance().GetLogger().Warnf("v2AFSExec", fmt.Sprintf("cannot retrieve afid from set param %s", cmdArg))
			}
		}
	**/
	//如果上面没有 return ,这里返回，就应该重置为false。后面
	engineResponse.IsBadParams = false
	engineResponse.IsTimeout = false

	plogger.NewInstance().GetLogger().Infof("AFSExec gives %s", string(result.Data))
	plogger.NewInstance().GetLogger().Infof("engine common exec ==>AFSExec ::response.sysCode=[%v];raw=%v \n ",engineResponse.SysCode,string(engineResponse.Raw))
	return engineResponse
}

func devTimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	plogger.NewInstance().GetLogger().Infof("[DEV]TIMETRACK,command: %s took %s\n", name, elapsed)
}
func v2TimeoutFunc(afsClient, cmdArg string, cChan chan v2CmdResponse) func() {
	return func() {
		plogger.NewInstance().GetLogger().Infof("%s killed execution exceeding max:%d\n", fmt.Sprintf("%s %s", afsClient, cmdArg), dict.MaxExecTime)
		cChan <- v2CmdResponse{nil, fmt.Errorf("request:%s %s is timeout. %s", afsClient, cmdArg, e.GetMsg(e.ScodeAFSTimeout)), e.ScodeAFSTimeout}
	}
}