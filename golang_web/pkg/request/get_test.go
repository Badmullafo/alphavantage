package request

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	//"gopkg.in/square/go-jose.v2/json"
)

func TestGetJson(t *testing.T) {
	type test struct {
		name   string
		apiKey string
		Symbol string
		Ndays  int
	}

	tests := []test{
		{"first", "RABZYXWVHB8MX5GO", "IBM", 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			d, err := GetJson(tt.apiKey, tt.Symbol, tt.Ndays)

			require.NoError(t, err)

			fmt.Println(PrettyPrint(d))

			h := d.DD["2021-06-24"]
			fmt.Printf("Type %T, value %v", h.High, h.High)

			//fmt.Println(d.MetaData.Info)
		})
	}
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
