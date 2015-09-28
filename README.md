# FlagStruct - simple way to register and parse flag into struct

## Install

    go get -u github.com/ngdinhtoan/flagstruct

## Example

Tag structure: `flag:"flagname[,default_value[,usage]]"`

```go
package main

import "github.com/ngdinhtoan/flagstruct"

type DbConfig struct {
    Hostname string `flag:"hostname,localhost,Hostname"`
    Port     uint64 `flag:"port"`
    DbName   string `flag:"db_name,,Database name"`
}

func main() {
    conf := DbConfig{}
    flagstruct.Parse(&conf)

    fmt.Println("Hostname: ", conf.Hostname)
    fmt.Println("Port: ", conf.Port)
    fmt.Println("DB Name: ", conf.DbName)
}
```

Run with some options:

    go run main.go -hostname=127.0.0.1 -port=5000 -db_name=test_db
