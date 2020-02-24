package rsa

import (
	"bfs_cli_rsa/common/dict"
	"bfs_cli_rsa/common/plogger"
	"bfs_cli_rsa/common/utils"
	"bfs_cli_rsa/models"
	"bfs_cli_rsa/service"
	"bfs_cli_rsa/service/base"
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

	plogger.NewInstance().GetLogger().Infof("获取版本：\n")
	plogger.NewInstance().GetLogger().Infof("scode=%d \n",sCode)
	plogger.NewInstance().GetLogger().Infof("version=%s \n",resp.Version)
	plogger.NewInstance().GetLogger().Infof("Core=%s \n",resp.Core)
	plogger.NewInstance().GetLogger().Infof("Raw=%v \n",resp.Raw)
	plogger.NewInstance().GetLogger().Infof("Note=%v \n",resp.Note)

	return sCode,resp
}

/**
 * 创建rsa 公私钥 对
 */
func (*rsaService)CreateRsaKeyPeer()(sCode int, response models.RsaKeyPeer){
	defer utils.DevTimeTrack(time.Now(), "CreateRsaKeyPeer")
	plogger.NewInstance().GetLogger().Info("AFS create rsa key peers  .")

	execTimeout:=dict.MaxExecTime
	sCode, resp := base.NewBaseServiceInstace().CreateRsaKeyPeer(execTimeout)

	plogger.NewInstance().GetLogger().Infof("公私钥密码对：\n")
	plogger.NewInstance().GetLogger().Infof("scode=%d \n",sCode)
	plogger.NewInstance().GetLogger().Infof("privateKey=%s \n",resp.PrivateKey)
	plogger.NewInstance().GetLogger().Infof("publicKey=%s \n",resp.PublicKey)
	plogger.NewInstance().GetLogger().Infof("Raw=%v \n",resp.Raw)
	plogger.NewInstance().GetLogger().Infof("Note=%v \n",resp.Note)
	return sCode,resp
}

func (*rsaService)GetPassWithPrivateKey(privateKeyHex string,sypassenHex string )(sCode int, response models.DeRsa){
	defer utils.DevTimeTrack(time.Now(), "deRsa with private key ")
	plogger.NewInstance().GetLogger().Info("AFS decrypte rsa using private key   .")

	execTimeout:=dict.MaxExecTime
	sCode, resp := base.NewBaseServiceInstace().DeRsaWithPrivateKey(privateKeyHex,sypassenHex,execTimeout)

	plogger.NewInstance().GetLogger().Infof("使用私钥解密密文获得明文：\n")
	plogger.NewInstance().GetLogger().Infof("scode=%d \n",sCode)
	plogger.NewInstance().GetLogger().Infof("mypass=%s \n",resp.Plaintext)
	plogger.NewInstance().GetLogger().Infof("Raw=%v \n",resp.Raw)
	plogger.NewInstance().GetLogger().Infof("Note=%v \n",resp.Note)
	return sCode,resp
}
