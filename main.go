package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/icodeface/hdkeyring"
	"os"
	"strings"
)

func main() {
	defer func() {
		fmt.Println("Press Enter to exit...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}()

	fmt.Print("Input your seed phrase:")
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	mnemonic := strings.TrimSpace(input)
	k, err := hdkeyring.NewFromMnemonic(mnemonic, hdkeyring.KeyTypeECDSA)
	if err != nil {
		panic(err)
	}

	var path hdkeyring.DerivationPath
	pk, err := k.DeriveECDSAPrivateKey(path)
	if err != nil {
		panic(err)
	}
	address.CurrentNetwork = address.Mainnet
	addr, err := address.NewSecp256k1Address(hdkeyring.ECDSAPublicKeyBytes(&pk.PublicKey))

	fmt.Println("Address:", addr.String())
	fmt.Println("Private Key:", hex.EncodeToString(pk.D.Bytes()))
}
