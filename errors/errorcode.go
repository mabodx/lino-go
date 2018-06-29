package errors

// Code types.
const (
	CodeOK CodeType = iota // 0
	CodeQueryFail
	CodeFailedToBroadcast
	CodeCheckTxFail
	CodeDeliverTxFail
	CodeFialedToGetPubKeyFromHex
	CodeFailedToGetPrivKeyFromHex
	CodeInvalidArg
	CodeEmptyResponse
)