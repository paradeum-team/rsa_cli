package e



var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERR_NO_AUTH:					 "没有权限",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",


	//服务器代码
	ScodeOK:                      "",
	ScodeErr:                     "server error",
	ScodeAFSErrNull:              "AFS null output",
	ScodeAFSErr:                  "AFS error",
	ScodeAPIErrParam:             "invalid parameter",
	ScodeAPIErrFile:              "invalid file",
	ScodeInvalidRenewal:          "cannot submit a file again on its creation day or renew expiration on the last day",
	ScodeAFSErrExec:              "AFS core execution error",
	ScodeAFSTimeout:              "AFS core timed out",
	ScodeNotEnoughFund:           "not enough fund",
	ScodeErrOS:                   "server os error",
	ScodeErrIO:                   "server io error",
	ScodeOpOnNotExist:            "operation on non-existing file",
	ScodeInvalidContract:         "invalid signed contract",
	ScodeRNodesInfoErr:           "failed to fetch rnodes info",
	ScodeNotFoundFNDat:           "cannot locate fnode-daily data",
	ScodeNotReadyFNDat:           "fnode-daily data not ready",
	ScodeNoFNDat:                 "no fnode data available",
	ScodeCriticalErr:             "server preparing broadcast request failed after payment",
	ScodeCriticalAFSTimeout:      "AFS core timed out after payment",
	ScodeCriticalAFCErr:          "broadcast request to AFC failed after payment",
	ScodeCriticalTrackErr:        "failed to query rnode info from track node for extension synchronization",
	ScodeCriticalTrackTimeout:    "timed out querying track node for extension synchonization",
	ScodeCriticalRnodeSyncExt:    "failed to synchronize extension on rnode",
	ScodeAFCErr:                  "request to AFC failed",
	ScodeAFCTimeout:              "request to AFC timed out",
	ScodeContractVerificationErr: "contract verification error",
	ScodeReusingContract:         "contract already processed and consumed before",
	ScodeTrackErr:                "failed to request track node api",
	ScodeConfigErr:               "current server configuration does not support this request",



}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

func GetCodeByMsg(value string) int {
	for k, x := range MsgFlags {
		if x == value {
			return k
		}
	}
	return -1
}