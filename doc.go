// Package flagstruct provide an easy way to register and parse flag into a struct
/*
Install

    go get -u github.com/ngdinhtoan/flagstruct

Tag syntax

	`flag:"name" default:"value" usage:"description"`

Tag `default` and `usage` can be omit.

Example

    package main

    import (
    	"fmt"

    	"github.com/ngdinhtoan/flagstruct"
    )

    type dbConfig struct {
    	Hostname string `flag:"hostname" default:"localhost" usage:"Hostname"`
    	Port     uint64 `flag:"port" default:"3306"`
    	DbName   string `flag:"db_name" usage:"Database name"`
    }

    func main() {
    	conf := dbConfig{}
    	flagstruct.Parse(&conf)

    	fmt.Println("Hostname:", conf.Hostname)
    	fmt.Println("Port:", conf.Port)
    	fmt.Println("DB Name:", conf.DbName)
    }

Run with some options:

    go run main.go -hostname=127.0.0.1 -db_name=test_db

Output:

    Hostname: 127.0.0.1
    Port: 3306
    DB Name: test_db
*/
package flagstruct
