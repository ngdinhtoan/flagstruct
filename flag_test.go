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

func TestAllType(t *testing.T) {
	type testStruct struct {
		Int        int     `flag:"int"`
		IntDef     int     `flag:"int-def,-1"`
		Int64      int64   `flag:"int64"`
		Int64Def   int64   `flag:"int64-def,1"`
		Uint       uint    `flag:"uint"`
		UintDef    uint    `flag:"uint-def,1"`
		Uint64     uint64  `flag:"uint64"`
		Uint64Def  uint64  `flag:"uint64-def,1"`
		String     string  `flag:"string"`
		StringDef  string  `flag:"string-def,abc"`
		Boolean    bool    `flag:"boolean"`
		BooleanDef bool    `flag:"boolean-def,true"`
		Float64    float64 `flag:"float64"`
		Float64Def float64 `flag:"float64-def,3.4"`
	}

	ts := testStruct{}
	fs := flag.NewFlagSet("test", flag.PanicOnError)
	if err := ParseByFlagSet(&ts, fs, []string{"--int=1"}); err != nil {
		t.Fatal(err)
	}

	if ts.Int != 1 || ts.BooleanDef != true {
		t.Fail()
	}
}

func TestErrUnsupportType(t *testing.T) {
	type ExportStruct struct{}
	type testStruct struct {
		ExportStruct         // anonymous field
		StringPtr    *string `flag:"string-ptr"`
	}

	err := Parse(&testStruct{})

	if err != ErrUnsupportType {
		t.Fail()
	}
}

func TestNotPointerToStruct(t *testing.T) {
	err := Parse("text")

	if err != ErrNotPointer {
		t.Fail()
	}
}

func ExampleParseByFlagSet() {
	type dbConfig struct {
		Host         string `flag:"host,localhost,hostname of database server"`
		Port         int64  `flag:"port,3306,port of database server"`
		DbName       string `flag:"db_name,test_db,database name"`
		Slave        bool   `flag:"slave"`
		MaxConnetion uint   `flag:"max_connection,50"`
	}

	dc := dbConfig{}
	fs := flag.NewFlagSet("test", flag.PanicOnError)
	if err := ParseByFlagSet(&dc, fs, []string{"-host=localhost", "-slave"}); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Host:", dc.Host)
	fmt.Println("Port:", dc.Port)

	// Output:
	// Host: localhost
	// Port: 3306
}
