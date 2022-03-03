package hd

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	dbpath := "../test.db"
	assert := assert.New(t)

	ts := httptest.NewServer(Handler(dbpath))
	defer ts.Close()
	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("test msg", string(data))
}
