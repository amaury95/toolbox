/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/amaury95/toolbox/src/util"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

// ethereumCmd represents the ethereum command
var ethereumCmd = &cobra.Command{
	Use:   "ethereum <text>",
	Short: "Sign content with an Ethereum private key (EIP-191 personal_sign)",
	Long:  `Signs the given text, prompts for the private key, and outputs the hex-encoded signature (0x-prefixed).`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		content := []byte(args[0])
		keyHex := strings.TrimSpace(util.PromptPrivateKey())
		keyHex = strings.TrimPrefix(keyHex, "0x")
		keyBytes, err := hex.DecodeString(keyHex)
		if err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), "Error: invalid private key (hex):", err)
			return
		}
		privateKey, err := crypto.ToECDSA(keyBytes)
		if err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), "Error: invalid private key:", err)
			return
		}
		// EIP-191 personal_sign: "\x19Ethereum Signed Message:\n" + len(message) + message
		msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(content), content)
		hash := crypto.Keccak256Hash([]byte(msg))
		signature, err := crypto.Sign(hash.Bytes(), privateKey)
		if err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), "Error signing:", err)
			return
		}
		// crypto.Sign returns recovery id 0/1; Ethereum uses v = 27 + recovery_id
		signature[64] += 27
		fmt.Println("0x" + hex.EncodeToString(signature))
	},
}

func init() {
	signCmd.AddCommand(ethereumCmd)
}
