package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ekalinin/btxer/btcrpc"
)

var (
	rpcAddr  = flag.String("rpc-addr", "127.0.0.1:8332", "Bitcoin RPC server")
	rpcUser  = flag.String("rpc-user", "", "User for RPC server")
	rpcPass  = flag.String("rpc-psw", "", "Password for RPC server")
	txNumber = flag.Int("tx-number", 20, "Total number of transactions to show")
	txSkip   = flag.Int("tx-skip", 0, "The number of transactions to skip")
	btcAddr  = flag.String("addr", "", "Bitcoin address")
	debug    = flag.Bool("debug", false, "Show debug")
)

func main() {

	flag.Parse()

	if *btcAddr == "" {
		fmt.Println("Please, enter Bitcoin address")
		flag.PrintDefaults()
		os.Exit(1)
	}

	rpc := btcrpc.New(*rpcAddr, *rpcUser, *rpcPass, *debug)

	// is Bitcoin address is not watched then we need to start watching it:
	// call importaddress RPC
	isWatched, err := rpc.IsAddrWatched(*btcAddr)
	if err != nil {
		panic(err)
	}
	if !isWatched {
		if err = rpc.ImportAddress(*btcAddr, true); err != nil {
			panic(err)
		}
	}
	txs, err := rpc.ListTransactions(*btcAddr, *txNumber, *txSkip)
	if err != nil {
		panic(err)
	}
	fmtString := "%3d |%70s |%10s |%15f |%30s |\n"
	if len(txs) > 0 {
		fmt.Printf("%3s |%70s |%10s |%15s |%30v |\n", "#", "txid", "category", "amount", "time")
	}
	for i, tx := range txs {
		fmt.Printf(fmtString, i+1, tx.Txid, tx.Category, tx.Amount, time.Unix(tx.Blocktime, 0))
	}
	balance, err := rpc.GetBalance(*btcAddr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Balance: %f\n", balance)
}
