# twse - TWSE API in Go
[![GoDoc](https://godoc.org/github.com/z-Wind/twse?status.png)](http://godoc.org/github.com/z-Wind/twse)

## Table of Contents

* [Installation](#installation)
* [Examples](#examples)
* [Reference](#reference)

## Installation

    $ go get github.com/z-Wind/twse

## Examples

### Client
```go
client := GetClient()
twse, err := New(client)
```

### Quotes
```go
call := twse.Quotes.GetStockInfoTWSE("0050")
stockInfo, err := call.Do()
```


## Reference
- [https://github.com/toomore/gogrs](https://github.com/toomore/gogrs)