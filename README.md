# banking 
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/kappapay/banking) 
[![Build Status](https://cloud.drone.io/api/badges/jbub/banking/status.svg)](https://cloud.drone.io/jbub/banking)
[![Go Report Card](https://goreportcard.com/badge/github.com/kappapay/banking)](https://goreportcard.com/report/github.com/kappapay/banking)

Banking library for Go.
This is a fork of the Banking library for Go. It adds IBAN validation support for [partial IBAN countries](https://www.iban.com/structure).

## Install

```bash
go get github.com/kappapay/banking
```

## Docs

http://godoc.org/github.com/kappapay/banking

## Iban

```go
package main

import (
    "fmt"
    "log"

    "github.com/kappapay/banking/iban"
)

var (
    testIban = iban.MustParse("BE68539007547034")
)

func main() {
    ibn, err := iban.Parse("BE68539007547034")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(ibn.BankCode())
    fmt.Println(ibn.AccountNumber())

    err = iban.Validate("BE68539007547034")
    if err != nil {
        log.Fatal(err)
    }
}
```

## Swift

```go
package main

import (
    "fmt"
    "log"

    "github.com/kappapay/banking/swift"
)

var (
    testSwift = swift.MustParse("DEUTDEFF500")
)

func main() {
    swft, err := swift.Parse("DEUTDEFF500")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(swft.BankCode())
    fmt.Println(swft.CountryCode())

    err = swift.Validate("DEUTDEFF500")
    if err != nil {
        log.Fatal(err)
    }
}
```
