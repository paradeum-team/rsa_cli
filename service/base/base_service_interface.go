package base

import "bfs_cli_rsa/models"

/**
 * 忽略参数合法性验证，合法性验证放在 appService层做
 * 定义 基础服务类，是为了匹配BFS的基础命令行的变化。
 * 区别于 AppService 接口
 */
type BaseService interface {
	//获取BFS版本
	GetVersion(execTimeout int) (sCode int, response models.VersionResponse)

    CreateRsaKeyPeer(execTimeout int)(sCode int, response models.RsaKeyPeer)

}

