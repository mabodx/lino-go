package transport

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/lino-network/lino-go/model"
	crypto "github.com/tendermint/go-crypto"
)

const (
	msgTypeRegister          = 0x1
	msgTypeFollow            = 0x2
	msgTypeUnfollow          = 0x3
	msgTypeTransfer          = 0x4
	msgTypePost              = 0x5
	msgTypeLike              = 0x6
	msgTypeDonate            = 0x7
	msgTypeValidatorDeposit  = 0x8
	msgTypeValidatorWithdraw = 0x9
	msgTypeValidatorRevoke   = 0x10
	msgTypeClaim             = 0x11
	msgTypeVoterDeposit      = 0x12
	msgTypeVoterRevoke       = 0x13
	msgTypeVoterWithdraw     = 0x14
	msgTypeDelegate          = 0x15
	msgTypeDelegatorWithdraw = 0x16
	msgTypeRevokeDelegation  = 0x17
	msgTypeVote              = 0x18
	msgTypeCreateProposal    = 0x19
	msgTypeDeveloperRegister = 0x20
	msgTypeDeveloperRevoke   = 0x21
	msgTypeProviderReport    = 0x22

	pubKeyTypeEd25519   = 0x1
	pubKeyTypeSecp256k1 = 0x2

	signatureTypeEd25519   = 0x1
	signatureTypeSecp256k1 = 0x2
)

type Transaction struct {
	Msg  []interface{} `json:"msg"`
	Sigs []interface{} `json:"signatures"`
}

type Signature struct {
	PubKey   []interface{} `json:"pub_key"`
	Sig      []interface{} `json:"signature"`
	Sequence int64         `json:"sequence"`
}

type SignMsg struct {
	ChainID   string  `json:"chain_id"`
	Sequences []int64 `json:"sequences"`
	MsgBytes  []byte  `json:"msg_bytes"`
}

func EncodeTx(msg interface{}, pubKey crypto.PubKey, sig crypto.Signature, seq int64) ([]byte, error) {
	typeMsg, err := GetMsgType(msg)
	if err != nil {
		return nil, err
	}

	typeKey, hexKey, err := GetPubKeyTypeAndHex(pubKey)
	if err != nil {
		return nil, err
	}

	typeSig, hexSig, err := GetSignatureTypeAndHex(sig)
	if err != nil {
		return nil, err
	}

	stdSig := Signature{
		PubKey:   []interface{}{typeKey, hexKey},
		Sig:      []interface{}{typeSig, hexSig},
		Sequence: seq,
	}

	stdTx := Transaction{
		Msg:  []interface{}{typeMsg, msg},
		Sigs: []interface{}{stdSig},
	}
	return json.Marshal(stdTx)
}

func EncodeMsg(msg interface{}) ([]byte, error) {
	typeMsg, err := GetMsgType(msg)
	if err != nil {
		return nil, err
	}
	stdMsg := []interface{}{typeMsg, msg}
	return json.Marshal(stdMsg)
}

func EncodeSignMsg(msgBytes []byte, chainId string, seq int64) ([]byte, error) {
	stdSignMsg := SignMsg{
		ChainID:   chainId,
		MsgBytes:  msgBytes,
		Sequences: []int64{seq},
	}
	return json.Marshal(stdSignMsg)
}

func GetMsgType(msg interface{}) (byte, error) {
	switch msg := msg.(type) {
	case model.TransferToAddress:
		return msgTypeTransfer, nil
	default:
		fmt.Println("Cannot find this message", msg)
		return 0, nil
	}
}

func GetPubKeyTypeAndHex(pubKey crypto.PubKey) (byte, string, error) {
	keyEd25519, ok := pubKey.(crypto.PubKeyEd25519)
	if ok {
		rawBytes := [32]byte(keyEd25519)
		return pubKeyTypeEd25519, hex.EncodeToString(rawBytes[:]), nil
	}

	keySecp256k1, ok := pubKey.(crypto.PubKeySecp256k1)
	if ok {
		rawBytes := [33]byte(keySecp256k1)
		return pubKeyTypeSecp256k1, hex.EncodeToString(rawBytes[:]), nil
	}
	return 0, "", nil
}

func GetSignatureTypeAndHex(sig crypto.Signature) (byte, string, error) {
	sigEd25519, ok := sig.(crypto.SignatureEd25519)
	if ok {
		rawBytes := [64]byte(sigEd25519)
		return signatureTypeEd25519, hex.EncodeToString(rawBytes[:]), nil
	}

	sigSecp256k1, ok := sig.(crypto.SignatureSecp256k1)
	if ok {
		return signatureTypeSecp256k1, hex.EncodeToString(sigSecp256k1[:]), nil
	}
	return 0, "", nil

}
