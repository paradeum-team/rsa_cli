package base

import (
	"bfs_cli_rsa/common/dict"
	"bfs_cli_rsa/common/plogger"
	"bfs_cli_rsa/engine"
	"bfs_cli_rsa/models"
)

/**
 * 忽略参数合法性验证，合法性验证放在 appService层做
 */
var baseServiceInstace *baseService

type baseService struct {
}

func NewBaseServiceInstace() BaseService {
	if baseServiceInstace == nil {
		baseServiceInstace = &baseService{}
	}
	return baseServiceInstace
}


//获取BFS版本
func (c *baseService) GetVersion(execTimeout int) (sCode int, response models.VersionResponse) {
	response = models.VersionResponse{}

	cmdArgs := dict.CombineCMD4Version
	plogger.NewInstance().GetLogger().Info("baseService:GetVersion .The cmdArgs=%s \n ", cmdArgs)
	engineResponse := engine.AFSExec4V2(dict.CMDIDVer, cmdArgs, execTimeout)
	sCode, response = engine.HandlerVersion(engineResponse)
	return sCode, response
}

//创建 rsakey peer
func (c *baseService) CreateRsaKeyPeer(execTimeout int) (sCode int, response models.RsaKeyPeer) {
	response = models.RsaKeyPeer{}

	cmdArgs := dict.CreateRsaKeyPeer
	plogger.NewInstance().GetLogger().Info("baseService:GetVersion .The cmdArgs=%s \n ", cmdArgs)
	engineResponse := engine.AFSExec4V2(dict.CMDIDVer, cmdArgs, execTimeout)
	sCode, response = engine.HandlerRsaKeyPeer(engineResponse)
	return sCode, response
}





func (c *baseService) DeRsaWithPrivateKey(privateKeyHex string,cipherText string , execTimeout int)(sCode int, response models.DeRsa){
	response = models.DeRsa{}

	cmdArgs := dict.CreateRsaKeyPeer
	plogger.NewInstance().GetLogger().Info("baseService:GetVersion .The cmdArgs=%s \n ", cmdArgs)
	engineResponse := engine.AFSExec4V2(dict.CMDIDVer, cmdArgs, execTimeout)
	sCode, response = engine.HandlerDeRsa(engineResponse)
	return sCode, response
}