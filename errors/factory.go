package errors

import "fmt"

// CodeToDefaultMsg returns the default message based on different error code.
func CodeToDefaultMsg(code CodeType) string {
	switch code {
	case CodeQueryFail:
		return "Failed To Query"
	case CodeFailedToBroadcast:
		return "Failed To Broadcast"
	case CodeCheckTxFail:
		return "Check Tx Error"
	case CodeDeliverTxFail:
		return "Deliver Tx Error"
	case CodeFailedToGetPubKeyFromHex:
		return "Failed To Get Pub Key From Hex"
	case CodeFailedToGetPrivKeyFromHex:
		return "Failed To Get Priv Key From Hex"
	case CodeInvalidArg:
		return "Invalid argument"
	case CodeInvalidNodeURL:
		return "Invalid node url"
	case CodeInvalidSequenceNumber:
		return "Invalid sequence number"
	case CodeEmptyResponse:
		return "Empty Response"
	default:
		return fmt.Sprintf("Unknown code %d", code)
	}
}

//CodeQueryFail creates an error with CodeQueryFail
func QueryFail(msg string) Error {
	return newError(CodeQueryFail, msg)
}

//QueryFailf creates an error with CodeQueryFail and formatted message
func QueryFailf(format string, args ...interface{}) Error {
	return newError(CodeQueryFail, fmt.Sprintf(format, args...))
}

//FailedToBroadcast creates an error with CodeFailedToBroadcast
func FailedToBroadcast(msg string) Error {
	return newError(CodeFailedToBroadcast, msg)
}

//FailedToBroadcastf creates an error with CodeFailedToBroadcast and formatted message
func FailedToBroadcastf(format string, args ...interface{}) Error {
	return newError(CodeFailedToBroadcast, fmt.Sprintf(format, args...))
}

//CheckTxFail creates an error with CodeCheckTxFail
func CheckTxFail(msg string) Error {
	return newError(CodeCheckTxFail, msg)
}

//CheckTxFailf creates an error with CodeCheckTxFail and formatted message
func CheckTxFailf(format string, args ...interface{}) Error {
	return newError(CodeCheckTxFail, fmt.Sprintf(format, args...))
}

//DeliverTxFail creates an error with CodeDeliverTxFail
func DeliverTxFail(msg string) Error {
	return newError(CodeDeliverTxFail, msg)
}

//DeliverTxFailf creates an error with CodeDeliverTxFail and formatted message
func DeliverTxFailf(format string, args ...interface{}) Error {
	return newError(CodeDeliverTxFail, fmt.Sprintf(format, args...))
}

//FailedToGetPubKeyFromHex creates an error with CodeFialedToGetPubKeyFromHex
func FailedToGetPubKeyFromHex(msg string) Error {
	return newError(CodeFailedToGetPubKeyFromHex, msg)
}

//FailedToGetPubKeyFromHex creates an error with CodeDeliverTxFail and formatted message
func FailedToGetPubKeyFromHexf(format string, args ...interface{}) Error {
	return newError(CodeFailedToGetPubKeyFromHex, fmt.Sprintf(format, args...))
}

//FailedToGetPrivKeyFromHex creates an error with CodeFailedToGetPrivKeyFromHex
func FailedToGetPrivKeyFromHex(msg string) Error {
	return newError(CodeFailedToGetPrivKeyFromHex, msg)
}

//FailedToGetPrivKeyFromHexf creates an error with CodeFailedToGetPrivKeyFromHex and formatted message
func FailedToGetPrivKeyFromHexf(format string, args ...interface{}) Error {
	return newError(CodeFailedToGetPrivKeyFromHex, fmt.Sprintf(format, args...))
}

//InvalidArg creates an error with CodeInvalidArg
func InvalidArg(msg string) Error {
	return newError(CodeInvalidArg, msg)
}

//InvalidArgf creates an error with CodeInvalidArg and formatted message
func InvalidArgf(format string, args ...interface{}) Error {
	return newError(CodeInvalidArg, fmt.Sprintf(format, args...))
}

//InvalidNodeURL creates an error with CodeInvalidNodeURL
func InvalidNodeURL(msg string) Error {
	return newError(CodeInvalidNodeURL, msg)
}

//InvalidNodeURLf creates an error with CodeInvalidNodeURL and formatted message
func InvalidNodeURLf(format string, args ...interface{}) Error {
	return newError(CodeInvalidNodeURL, fmt.Sprintf(format, args...))
}

//InvalidSequenceNumber creates an error with CodeInvalidSequenceNumber
func InvalidSequenceNumber(msg string) Error {
	return newError(CodeInvalidSequenceNumber, msg)
}

//InvalidSequenceNumberf creates an error with CodeInvalidSequenceNumber and formatted message
func InvalidSequenceNumberf(format string, args ...interface{}) Error {
	return newError(CodeInvalidSequenceNumber, fmt.Sprintf(format, args...))
}

//EmptyResponse creates an error with CodeEmptyResponse
func EmptyResponse(msg string) Error {
	return newError(CodeEmptyResponse, msg)
}

//EmptyResponsef creates an error with CodeEmptyResponse and formatted message
func EmptyResponsef(format string, args ...interface{}) Error {
	return newError(CodeEmptyResponse, fmt.Sprintf(format, args...))
}
