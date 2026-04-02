package headers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeaderParse(t *testing.T) {
	// Test: Valid single header
	headers := NewHeaders()
	data := []byte("Host: localhost:42069\r\nFoofoo:    barbar\r\n\r\n")
	n, done, err := headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)

	host, ok := headers.Get("host")
	assert.True(t, ok)
	assert.Equal(t, "localhost:42069", host)

	foo, ok := headers.Get("Foofoo")
	assert.True(t, ok)
	assert.Equal(t, "barbar", foo)

	missingKey, ok := headers.Get("host")
	assert.True(t, ok)
	assert.Equal(t, "", missingKey)

	assert.Equal(t, 44, n)
	assert.True(t, done)

	// Test: Invalid spacing header
	headers = NewHeaders()
	data = []byte("       Host : localhost:42069       \r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	// Test: Invalid char in fieldname
	headers = NewHeaders()
	data = []byte("       H°st: localhost:42069       \r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	// Test: more than one field value
	headers = NewHeaders()
	data = []byte("Host: localhost:42069\r\nHost: localhost:42069       \r\n")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)

	host, ok = headers.Get("host")
	assert.True(t, ok)
	assert.Equal(t, "localhost:42069,localhost:42069", host)
	assert.False(t, done)
}
