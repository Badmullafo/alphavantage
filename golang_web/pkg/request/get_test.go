package request

import (
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

			//	json.MarshalIndent(d, "   ", "    ")
			fmt.Println(d)
			//fmt.Println(d.DD)

			//fmt.Println(d.MetaData.Info)
		})
	}
}
