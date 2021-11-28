package request

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	//"gopkg.in/square/go-jose.v2/json"
	"io"
	"net/http"
	"net/http/httptest"
)

type testWeb struct {
	url  string
	body string
}

func TestGetJson(t *testing.T) {
	tests := []struct {
		name, apiKey, Symbol string
		Ndays                int
		web                  testWeb
	}{
		{
			name:   "first",
			apiKey: "RABZYXWVHB8MX5GO",
			Symbol: "IBM",
			Ndays:  25,
			web: testWeb{
				url:  "http://faketest.com",
				body: "Some stuff here",
			},
		},
		{
			name:   "first",
			apiKey: "RABZYXWVHB8MX5GO",
			Symbol: "IBM",
			Ndays:  25,
			web: testWeb{
				url:  "http://faketest.com",
				body: "Some stuff here",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			d, err := GetJson(tt.apiKey, tt.Symbol, tt.Ndays)

			require.NoError(t, err)

			fmt.Println(PrettyPrint(d))

			require.NoError(t, err)

			resp := ExampleResponseRecorder(tt.web.url, tt.web.body)

			require.Equal(t, resp.StatusCode, 200)

			//fmt.Println(d.MetaData.Info)
		})
	}
}

func TestGetJsonReal(t *testing.T) {
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
			r := Result{
				Symbol: "IBM",     // the stock name .e.g. FORG
				Ndays:  2,         // Number of days data to get
				Dtype:  "average", // data type - total/average
			}

			dm, err := r.getInRange(d)

			require.NoError(t, err)

			fmt.Printf("Type %T, value %v", dm, dm)

			require.NotEmpty(t, d.DD)

			r.Getot(dm, "high")

			fmt.Println("hi")
			fmt.Println(r)

			//fmt.Println(d.MetaData.Info)
		})
	}
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func ExampleResponseRecorder(url, rbody string) *http.Response {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, rbody)
	}

	req := httptest.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	handler(w, req)

	return w.Result()
	//body, _ := io.ReadAll(resp.Body)

	//fmt.Println(resp.StatusCode)
	//fmt.Println(resp.Header.Get("Content-Type"))
	//return fmt.Sprintf(string(body))

	// Output:
	// 200
	// text/html; charset=utf-8
	// <html><body>Hello World!</body></html>
}
