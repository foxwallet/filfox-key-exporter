# Filfox Key Exporter
Filfox Web Wallet(https://wallet.filfox.info) is no longer maintained, but their private key derivation method is incompatible with most mainstream wallets, including FoxWallet.
  
Here is a tool to help you export private keys.

## Usage
- Download this tool from https://github.com/foxwallet/filfox-key-exporter/releases
  - `filfox-key-exporter_darwin_arm64` for arm mac 
  - `filfox-key-exporter_win64.exe` for windows
- Double-click the executable file to run it
- Input your seed phrase created in wallet.filfox.info, and export private key
- Import the private key to [FoxWallet](https://foxwallet.com/download) and transfer

## Build Yourself
- `git clone https://github.com/foxwallet/filfox-key-exporter`
- `cd filfox-key-exporter`
- `make all`
