# VWAP streamer from Coinbase
This application will connect to coinbase websocket and retrieve
## Build

```sh
go build 
```

## Run

```sh
./vwap
```

## Options

```sh
Usage of ./vwap:
  -product_ids string
        coma separated list of product ids (default "BTC-USD,ETH-USD,ETH-BTC")
  -size int
        window size (default 200)
```

## Test

```sh
go test ./...
```