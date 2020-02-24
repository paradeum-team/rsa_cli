package engine

import (
	"bfs_cli_rsa/common/e"
	"bfs_cli_rsa/common/plogger"
	"bfs_cli_rsa/models"
	"strings"
)

/**
 * version
 */
func HandlerRsaKeyPeer(enginResponse EngineResponse) (sCode int, response models.RsaKeyPeer) {
	raw := string(enginResponse.Raw)
	raw = strings.TrimSuffix(raw, "\n")
	sCode, note := handlerException(enginResponse)
	rsaPeer := models.RsaKeyPeer{Raw: raw, Note: note}
	if sCode != e.ScodeOK {
		return sCode, rsaPeer
	}

	sCode = e.ScodeAFSErr

	if strings.Contains(raw, "_r=true") {
		chunks := strings.Split(raw, ";")
		for _, chunk := range chunks {

			items:=strings.Split(chunk,"=")

			if len(items)==2{
				if "asymprivatekey"==items[0]{
					rsaPeer.PrivateKey=items[1]
				}else if "asympubkey"==items[0]{
					rsaPeer.PublicKey=items[1]
				}
			}
		}

		if rsaPeer.PublicKey != "" && rsaPeer.PrivateKey != "" {
			sCode = e.ScodeOK
		}
	}

	plogger.NewInstance().GetLogger().Infof("handler create rsa keys peer sCode=%d,rsaKeyPeers=%v \n", sCode, rsaPeer)

	return sCode, rsaPeer
}


/**
 * version
 */
func HandlerVersion(enginResponse EngineResponse) (sCode int, response models.VersionResponse) {
	raw := string(enginResponse.Raw)
	raw = strings.TrimSuffix(raw, "\n")
	sCode, note := handlerException(enginResponse)
	version := models.VersionResponse{Raw: raw, Note: note}
	if sCode != e.ScodeOK {
		return sCode, version
	}

	tempVer, tempCore := "", ""
	sCode = e.ScodeAFSErr

	if strings.Contains(raw, "_r=true") {
		chunks := strings.Split(raw, ";")
		for _, chunk := range chunks {
			if pos := strings.Index(chunk, "afs_v="); pos > -1 {
				tempVer = chunk[pos+len("afs_v="):]
			} else if pos := strings.Index(chunk, "core_v="); pos > -1 {
				tempCore = chunk[pos+len("core_v="):]
			}
		}

		if tempVer != "" && tempCore != "" {
			version.Core = tempCore
			version.Version = tempVer
			sCode = e.ScodeOK
		}
	}

	plogger.NewInstance().GetLogger().Infof("handler version sCode=%d,versionResp=%v \n", sCode, version)

	return sCode, version
}


/**
 * 统一处理异常，超时
 */
func handlerException(enginResponse EngineResponse) (sCode int, note string) {
	note = ""
	if enginResponse.IsTimeout {
		note = "connection AFS Server is timeout"
		return e.ScodeAFSTimeout, note
	} else if enginResponse.IsBadParams {
		note = "invalid params"
		return e.ScodeAPIErrParam, note

	}
	return e.ScodeOK, note
}