!!!!NEVER PUBLICLY POPULATE THIS CODE WITH A PRIVATE KEY!!!!

!!!!ALWAYS USE ON AN AIR GAPPED MACHINE AND TRANSFER THE SIGNED TRANSACTION BY USB OR SIMILAR TO AN ONLINE MACHINE"

!!!NEVER REUSE AN ADDRESS - THINK PRIVACY!!!

# Manually create and broadcast a BSV transaction to the network

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
* utxo private key (do not send out or allow off machine. delete when tx broadcast. preferably create the tx on an air gapped machine)
