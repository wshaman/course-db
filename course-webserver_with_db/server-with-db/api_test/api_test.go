package api_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/wshaman/server-with-db/datasource"
	"github.com/wshaman/server-with-db/server"
)

const baseURL = "http://localhost:8080"

func getPage(method, url string, body io.Reader) (*http.Response, error) {
	path := baseURL + url
	r, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	c := http.Client{}
	w, err := c.Do(r)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func testInit(m *testing.M) int {
	ds := datasource.NewMock()
	go server.Run("localhost:8080", ds)
	time.Sleep(1 * time.Second) // Usually we want some kind of waitfor.it, keep in mind sleep is just a simplification
	return m.Run()
}

func TestMain(m *testing.M) {
	os.Exit(testInit(m))
}

func TestUsers_01(t *testing.T) {
	r, err := getPage(http.MethodGet, "/users", nil)
	require.NoError(t, err)
	d, err := ioutil.ReadAll(r.Body)
	require.NoError(t, err)
	var out = []string{}
	json.Unmarshal(d, &out)
	assert.Equal(t, 3, len(out))
	//["James Alan Hetfield","Till Lindemann","Pär Sundström"]
}
