# Manually create and briadcast a BSV transaction to the network

First stab at building a BSV transaction and broadcasting to the network

Populate commented fields in the go code

USE: go run ./2026hash.go // to build the raw transaction from the CLI

Copy and paste the output into WhatsOnChain decode tool to check its encoded correctly

If good, paste into the WoC broadcast tool

Watch the tx on chain as it gets mined and confirmed for at least 6 tx's

Fields you need to gather and fill out in the go file
* inout txid
* input txid script pubkeyhash
* vin value
* receive address(to self) and vin(minus fee)
* utxo private key (don not send out or allow off machine. delete when tx broadcast. preferably create the tx on an air gapped machine)
