Simple bitcoin transactions tracker
===================================
**btxer** — **b**itcoin transaction (**tx**) track**er**.

Based on the original Bitcoin node with JSON-RPC API.

Main idea: to use [Watch-only Addresses](https://blog.blockchain.com/2016/05/31/how-to-use-watch-only-addresses/).

Table of Contents
-----------------

* [Simple bitcoin transactions tracker](#simple-bitcoin-transactions-tracker)
  * [Used RPC API](#used-rpc-api)
  * [Compile](#compile)
  * [Usage](#usage)
  * [Test bitcoind config](#test-bitcoind-config)
  * [Possible improvements](#possible-improvements)

Created by [gh-md-toc](https://github.com/ekalinin/github-markdown-toc.go)


Used RPC API
------------

* `importaddress` / GH PRs: [#2121](https://github.com/bitcoin/bitcoin/pull/2121), [#2861](https://github.com/bitcoin/bitcoin/pull/2861),  [#3383](https://github.com/bitcoin/bitcoin/pull/3383), [#4045](https://github.com/bitcoin/bitcoin/pull/4045)
* `listtransactions`
* `getbalance`
* `gettxoutsbyaddress` / GH PR [#5048](https://github.com/bitcoin/bitcoin/pull/5048)


Compile
-------

```bash
$ make build
```

Usage
-----

Let's take some address from testnet — [mmtrt5b6mhRNopMz6s8q7sBjUEV7ComK6J](https://www.blocktrail.com/tBTC/address/mmtrt5b6mhRNopMz6s8q7sBjUEV7ComK6J).

Here's a list of transactions for that address:
```bash
$ ./btxer -rpc-user "rpc" -rpc-psw "qwerty" -addr "mmtrt5b6mhRNopMz6s8q7sBjUEV7ComK6J"
  # |                                                                  txid |  category |         amount |                          time |
  1 |      37e0fa245375d39008cacd9fdaef97f8605e516c33b6ce0aa53300f737ba1d40 |   receive |       0.162500 | 2017-09-12 10:16:31 +0300 MSK |
  2 |      c5cb2d73eaa5ae124936f9bbc49fc941da63c1508dd6d95f6284e4e40bc07ed1 |   receive |       0.081250 | 2017-09-12 10:24:08 +0300 MSK |
  3 |      0919c3bfae64fe3c127b9a9d42bc3218a8652d3760e8814d210692126e6097a0 |   receive |       0.040625 | 2017-09-12 11:03:10 +0300 MSK |
  4 |      89a793d1d5ddc31bd8e5eff85270a1820970b591cc404eb40f2412fd9683db69 |   receive |       0.002539 | 2017-09-12 14:34:31 +0300 MSK |
  5 |      02a47797f928f1c02d460ee6d868a70e72c3921a7c1195bbf4ad6103c0ff9f88 |   receive |       0.001270 | 2017-09-12 14:37:03 +0300 MSK |
  6 |      c3c7a94192dc622b82ace9a4da7c1e563149f142a6ef19f4c73033f36e9ef6c1 |   receive |       0.000635 | 2017-09-12 14:39:00 +0300 MSK |
  7 |      054dffabdffd8051667138a746efd276f55448cc6ba3c3aabe157a1b70640a17 |   receive |       1.300000 | 2017-09-12 14:48:25 +0300 MSK |
  8 |      dca8ff7ad23dff05ea5585e36ff429bf07e11b31b41df56c66472ff84e1a4888 |   receive |       0.650000 | 2017-09-12 15:02:48 +0300 MSK |
  9 |      4c292e19f35e7672572e123d7e13c8b2dee3d83c9af50b06960e36f3c8796d3f |   receive |       0.162500 | 2017-09-12 15:15:10 +0300 MSK |
 10 |      1c2568c0e1a1b9f622cde42529df1a717a65e209315c1c7389dc00feb328d509 |   receive |       0.325000 | 2017-09-12 15:11:10 +0300 MSK |
 11 |      025f5bc2d61936b4c10c981d13dc001d4ecc0da3e590af2fb3c4d56f84555992 |   receive |       0.081250 | 2017-09-12 15:17:00 +0300 MSK |
Balance: 2.807568
```

*Warning*: if addess is not watched by bitcoin node, first call may be a bit
time consuming (~ 8—10 minutes).

Help:

```bash
➥ ./btxer --help
Usage of ./btxer:
  -addr string
    	Bitcoin address
  -debug
    	Show debug
  -rpc-addr string
    	Bitcoin RPC server (default "127.0.0.1:8332")
  -rpc-psw string
    	Password for RPC server
  -rpc-user string
    	User for RPC server
  -tx-number int
    	Total number of transactions to show (default 20)
  -tx-skip int
    	The number of transactions to skip
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


Possible improvements
---------------------

* To use [json-iterator](https://github.com/json-iterator/go) instead of `encoding/json`
* To use [decimal](https://github.com/shopspring/decimal) instead of `float`