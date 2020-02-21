package e

/**
 * 自定内内部错误好，统一异常信息处理
 */
const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
	ERR_NO_AUTH    = 401

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	//服务器代码
	ScodeOK              = 0
	ScodeErr             = 30001
	ScodeAPIErrParam     = 30004
	ScodeAPIErrFile      = 30005
	ScodeInvalidContract = 30013
	ScodeAFSErrNull      = 30002
	ScodeAFSErr          = 30003
	ScodeAFSErrExec      = 30006
	ScodeAFSTimeout      = 30007
	ScodeNotEnoughFund   = 30008
	ScodeInvalidRenewal  = 50009
	ScodeOpOnNotExist    = 30010

	ScodeErrOS = 30011
	ScodeErrIO = 30012

	ScodeNotFoundFNDat = 30051
	ScodeNotReadyFNDat = 30052
	ScodeNoFNDat       = 30053

	ScodeAFCErr                  = 30061
	ScodeAFCTimeout              = 30063
	ScodeContractVerificationErr = 30064
	ScodeReusingContract         = 30065

	ScodeTrackErr      = 30087
	ScodeRNodesInfoErr = 30088

	ScodeConfigErr = 30099

	ScodeCriticalErr        = 39021
	ScodeCriticalAFSErr     = 39022
	ScodeCriticalAFSTimeout = 39023

	ScodeCriticalAFCErr       = 39031
	ScodeCriticalRnodeSyncExt = 39036

	ScodeCriticalTrackErr     = 39054
	ScodeCriticalTrackTimeout = 39055
)
