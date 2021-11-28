package request

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	//"gopkg.in/square/go-jose.v2/json"
)

// Custom type that allows setting the func that our Mock Do func will run instead
type MockDoType func(req *http.Request) (*http.Response, error)

// MockClient is the mock client
type MockClient struct {
	MockDo MockDoType
}

// Overriding what the Do function should "do" in our MockClient
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

type testWeb struct {
	url        string
	filebody   string
	timeout    time.Duration
	statuscode int
	reqval     string
}

func TestGitHubCallSuccess(t *testing.T) {

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
				url:        "http://faketest.com",
				filebody:   "mocks/today.json",
				timeout:    time.Second * 20,
				statuscode: 200,
				reqval:     "495.04",
			},
		},
		{
			name:   "Timeout",
			apiKey: "RABZYXWVHB8MX5GO",
			Symbol: "IBM",
			Ndays:  5,
			web: testWeb{
				url:        "http://faketest.com",
				filebody:   "mocks/today.json",
				timeout:    time.Millisecond * 1,
				statuscode: 400,
				reqval:     "59.00",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			content, err := ioutil.ReadFile(tt.web.filebody)
			require.NoError(t, err)

			// create a new reader with that JSON
			r := ioutil.NopCloser(bytes.NewReader([]byte(content)))

			Client := &MockClient{
				MockDo: func(*http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: 200,
						Body:       r,
					}, nil
				},
			}

			nr := NewRequest(Client, tt.apiKey, tt.Symbol, tt.Ndays, time.Second*2)

			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(tt.web.timeout))
			defer cancel()

			dd, err := nr.GetJson(ctx)
			require.NoError(t, err)

			res := &Result{}

			res.Getot(dd, "high")

			assert.Equal(t, res.String(), tt.web.reqval)

			fmt.Println(res)

		})
	}
}

/*

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
*/
