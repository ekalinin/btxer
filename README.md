Simple bitcoin transactions tracker
===================================
btxer — **b**itcoin transaction (**tx**) track**er**.

Based on the original Bitcoin node with JSON-RPC API.
Main idea — to use [Watch-only Addresses](https://blog.blockchain.com/2016/05/31/how-to-use-watch-only-addresses/).

Used RPC API
------------

* `importaddress` / GH PRs: [2121](https://github.com/bitcoin/bitcoin/pull/2121), [2861](https://github.com/bitcoin/bitcoin/pull/2861),  [3383](https://github.com/bitcoin/bitcoin/pull/3383), [4045](https://github.com/bitcoin/bitcoin/pull/4045)
* `listtransactions`
* `getbalance`
* `gettxoutsbyaddress` / [GH PR](https://github.com/bitcoin/bitcoin/pull/5048)


Compile
-------

```bash
$ make build
```

Usage
-----

```bash
$ ./btxer -rpc-user "rpc" -rpc-psw "qwerty" -addr "n2CJu4WM5EdWmUBakvWdYRJmhdmaq9BHwy"
```


Test bitcoind config
--------------------

```bash
$ cat ~/.bitcoin/bitcoin.conf 
# Run on the test network instead of the real bitcoin network.
testnet=1

# server=1 tells Bitcoin-Qt and bitcoind to accept JSON-RPC commands
server=1

# RPC settings
rpcuser=rpc

rpcpassword=qwerty

# Rescan of the wallet may take a considerable amount of time
rpctimeout=120

rpcport=8332
```