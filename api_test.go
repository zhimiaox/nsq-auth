package main

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPluginRootSecret_Authorization(t *testing.T) {
	get, err := http.Get("http://localhost:1325/auth?secret=123")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 200, get.StatusCode)
	body, _ := ioutil.ReadAll(get.Body)
	_ = get.Body.Close()
	t.Log(string(body))
}

func BenchmarkApi_Auth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		get, err := http.Get("http://localhost:1325/auth?secret=123")
		if err != nil {
			b.Fatal(err)
		}
		assert.Equal(b, 200, get.StatusCode)
	}
}
