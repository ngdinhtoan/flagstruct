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
