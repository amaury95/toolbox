package gen

import (
	"crypto/ecdsa"
	"encoding/hex"
	"log"

	"github.com/amaury95/toolbox/src/util"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
)

func GenerateEthereumKey(output, encryptPassword, protocol string, tags ...string) {
	switch protocol {
	case "ethereum":
		generateEthereumKey(output, encryptPassword, tags...)
	default:
		log.Fatalf("Unsupported protocol: %s", protocol)
	}
}

func generateEthereumKey(output, encryptPassword string, tags ...string) {
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
		if output == "" {
			output = address.Hex()
		}
		if err := util.CreateEncryptedZip(output, privateKeyHex, encryptPassword, tags...); err != nil {
			log.Fatalf("Failed to zip private key with password: %v", err)
		}
	}
}

func GenerateMnemonic(name, encryptPassword string, tags ...string) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		log.Fatalf("Failed to generate entropy: %v", err)
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		log.Fatalf("Failed to generate mnemonic: %v", err)
	}

	log.Println("Generated Mnemonic:", mnemonic)

	if encryptPassword != "" {
		if name == "" {
			name = "mnemonic"
		}
		if err := util.CreateEncryptedZip(name, mnemonic, encryptPassword, tags...); err != nil {
			log.Fatalf("Failed to zip private key with password: %v", err)
		}
	}
}
