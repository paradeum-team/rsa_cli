package service

import "bfs_cli_rsa/models"

type RsaService interface {
	//获取BFS版本
	GetVersion() (sCode int, response models.VersionResponse)

	CreateRsaKeyPeer()(sCode int, response models.RsaKeyPeer)

	GetPassWithPrivateKey(privateKeyHex string,sypassenHex string )(sCode int, response models.DeRsa)

}