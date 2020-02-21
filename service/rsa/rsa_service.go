package rsa

import (
	"bfs_cli_rsa/common/dict"
	"bfs_cli_rsa/common/plogger"
	"bfs_cli_rsa/common/utils"
	"bfs_cli_rsa/models"
	"bfs_cli_rsa/service"
	"bfs_cli_rsa/service/base"
	"fmt"
	"time"
)
var rsaServiceInstace *rsaService

type rsaService struct {
}

func NewRsaServiceInstace() service.RsaService {
	if rsaServiceInstace == nil {
		rsaServiceInstace = &rsaService{}
	}
	return rsaServiceInstace
}

/**
 * 查询版本号
 */
func (*rsaService)GetVersion( ) (sCode int, response models.VersionResponse){
	defer utils.DevTimeTrack(time.Now(), "GetVersion")
	plogger.NewInstance().GetLogger().Info("AFS Get version .")
	execTimeout:=dict.MaxExecTime
	sCode, resp := base.NewBaseServiceInstace().GetVersion(execTimeout)

	fmt.Println("获取版本：")
	fmt.Printf("scode=%d \n",sCode)
	fmt.Printf("resp=%v \n",resp)

	return sCode,resp
}

/**
 * 创建rsa 公私钥 对
 */
func (*rsaService)CreateRsaKeyPeer()(sCode int, response models.RsaKeyPeer){
	defer utils.DevTimeTrack(time.Now(), "GetVersion")
	plogger.NewInstance().GetLogger().Info("AFS Get version .")

	execTimeout:=dict.MaxExecTime
	sCode, resp := base.NewBaseServiceInstace().CreateRsaKeyPeer(execTimeout)

	fmt.Println("公私钥密码对：")
	fmt.Printf("scode=%d \n",sCode)
	fmt.Printf("resp=%v \n",resp)
	return sCode,resp
}
