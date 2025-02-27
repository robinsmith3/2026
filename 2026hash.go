package main

import (
	"context"
	"log"

	"github.com/libsv/go-bk/wif"
	"github.com/libsv/go-bt/v2"
	"github.com/libsv/go-bt/v2/unlocker"
)

func main() {
	tx := bt.NewTx()

	_ = tx.From(
		"5911daf19c139eb8aa6e700616616b885c273ae16529953d039da0f9080dbba1", // input txid
		0,
		"76a914857e6c152fb8fd0c98ffaadade597c6d72cea55688ac", // input txid script pubkeyhash
		344472, // vin value
	)

	_ = tx.PayToAddress("1KqGmXkvNbbJMSExrVN7V4UYeNG4Udiu7C", 344272) // receive address(to self) and vin - fee

	_ = tx.AddOpReturnOutput([]byte("44f5434966f56c4ab8afaca7b82607d418abfad31b1d664a0ed8b354acee0702-2026_Great_Recession_Roadmap.pdf")) // data :sha256 hash:filename:cloudlink

	decodedWif, _ := wif.DecodeWIF("DEADBEEF") // utxo private key

	err := tx.FillAllInputs(context.Background(), &unlocker.Getter{PrivateKey: decodedWif.PrivKey})
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("tx: ", tx.String())
}
