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
		fmt.Println(err)
		return
	}

	mnemonic := strings.TrimSpace(input)
	k, err := hdkeyring.NewFromMnemonic(mnemonic, hdkeyring.KeyTypeECDSA)
	if err != nil {
		fmt.Println(err)
		return
	}

	var path hdkeyring.DerivationPath
	pk, err := k.DeriveECDSAPrivateKey(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	address.CurrentNetwork = address.Mainnet
	addr, err := address.NewSecp256k1Address(hdkeyring.ECDSAPublicKeyBytes(&pk.PublicKey))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Address:", addr.String())
	fmt.Println("Private Key:", hex.EncodeToString(pk.D.Bytes()))
}
