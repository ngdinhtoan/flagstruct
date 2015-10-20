# FlagStruct - simple way to register and parse flag into struct

[![Build Status](https://travis-ci.org/ngdinhtoan/flagstruct.svg)](https://travis-ci.org/ngdinhtoan/flagstruct)
[![GoDoc](https://godoc.org/github.com/ngdinhtoan/flagstruct?status.svg)](https://godoc.org/github.com/ngdinhtoan/flagstruct)

## Install

    go get -u github.com/ngdinhtoan/flagstruct

## Tag syntax

	`flag:"flagname[,default_value[,usage]]"`

## Example

```go
package main

import (
	"fmt"

	"github.com/ngdinhtoan/flagstruct"
)

type dbConfig struct {
	Hostname string `flag:"hostname,localhost,Hostname"`
	Port     uint64 `flag:"port,3306"`
	DbName   string `flag:"db_name,,Database name"`
}

func main() {
	conf := dbConfig{}
	flagstruct.Parse(&conf)

	fmt.Println("Hostname:", conf.Hostname)
	fmt.Println("Port:", conf.Port)
	fmt.Println("DB Name:", conf.DbName)
}
```

Run with some options:

    go run main.go -hostname=127.0.0.1 -db_name=test_db

Output:

    Hostname: 127.0.0.1
    Port: 3306
    DB Name: test_db
