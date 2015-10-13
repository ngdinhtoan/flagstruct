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
	if err := flagstruct.Parse(&conf); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Hostname:", conf.Hostname)
	fmt.Println("Port:", conf.Port)
	fmt.Println("DB Name:", conf.DbName)
}
