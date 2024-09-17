package gen

import (
	"crypto/ecdsa"
	"encoding/hex"
	"log"

	"github.com/amaury95/toolbox/src/util"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateEthereumKey(encryptPassword string, tags ...string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Failed to convert public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Println("Generated Ethereum Address:", address.Hex())

	privateKeyHex := hex.EncodeToString(crypto.FromECDSA(privateKey))
	log.Println("Generated Private Key:", privateKeyHex)

	if encryptPassword != "" {
		if err := util.CreateEncryptedZip(address.Hex(), privateKeyHex, encryptPassword, tags...); err != nil {
			log.Fatalf("Failed to zip private key with password: %v", err)
		}
	}
}
