package twse

import (
	"fmt"
	"testing"
)

func TestNewServer(t *testing.T) {
	client := GetClient()
	twse, err := New(client)
	if err != nil {
		t.Fatal(err)
	}

	call := twse.Quotes.GetStockInfoTWSE("0050")
	stockInfo, err := call.Do()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", stockInfo)
}
func ExampleQuotesService_GetStockInfoTWSE() {
	client := GetClient()
	twse, err := New(client)
	if err != nil {
		panic(err)
	}

	call := twse.Quotes.GetStockInfoTWSE("0050")
	stockInfo, err := call.Do()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", stockInfo)
}
