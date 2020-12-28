package types

import (
	"crypto/ed25519"
	"encoding/hex"

	"github.com/portto/solana-go-sdk/common"
)

type Account struct {
	PublicKey  common.PublicKey
	PrivateKey ed25519.PrivateKey
}

func AccountFromPrivateKeyBytes(privateKey []byte) Account {
	sk := ed25519.PrivateKey(privateKey)
	return Account{
		PublicKey:  common.PublicKeyFromHex(hex.EncodeToString(sk.Public().(ed25519.PublicKey))),
		PrivateKey: sk,
	}
}

type AccountMeta struct {
	PubKey     common.PublicKey
	IsSigner   bool
	IsWritable bool
}