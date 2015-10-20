package flagstruct

import (
	"flag"
	"fmt"
	"testing"
)

type dbconfig struct {
	Host         string `flag:"host,localhost,hostname of database server"`
	Port         int64  `flag:"port,3306,port of database server"`
	DbName       string `flag:"db_name,test_db,database name"`
	Slave        bool   `flag:"slave"`
	MaxConnetion uint   `flag:"max_connection,50"`

	DontParse string
}

func TestParse(t *testing.T) {
	dc := dbconfig{}
	fs := flag.NewFlagSet("test", flag.PanicOnError)
	if err := ParseByFlagSet(&dc, fs, []string{"-host=localhost", "--port", "1234", "-slave"}); err != nil {
		t.Fatal(err)
	}

	var expectedHost = "localhost"
	if dc.Host != expectedHost {
		t.Fatalf("Host name must be %q, get %q", expectedHost, dc.Host)
	}

	var expectedMaxConn = uint(50)
	if dc.MaxConnetion != expectedMaxConn {
		t.Fatalf("Max connection value must be %d, get %d", expectedMaxConn, dc.MaxConnetion)
	}

	fmt.Println("Test Data: ")
	fmt.Printf("%+v\n\n", dc)

	fs.PrintDefaults()
}
