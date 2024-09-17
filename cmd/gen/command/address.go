package command

import (
	"crypto/ecdsa"
	"encoding/hex"
	"log"

	"github.com/amaury95/tools/src/util"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

var _ethereum_tags []string
var _ethereum_encrypt_password string

var ethereumCmd = &cobra.Command{
	Use:   "address",
	Short: "Generate Ethereum private key and address",
	Run: func(cmd *cobra.Command, args []string) {
		generateEthereumKey(_ethereum_encrypt_password, _ethereum_tags...)
	},
}

func init() {
	rootCmd.AddCommand(ethereumCmd)
	ethereumCmd.Flags().StringSliceVarP(&_ethereum_tags, "tag", "t", []string{}, "Tags for the password")
	ethereumCmd.Flags().StringVarP(&_ethereum_encrypt_password, "encrypt-password", "e", "", "Password for the private key")
}

func generateEthereumKey(encryptPassword string, tags ...string) {
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
