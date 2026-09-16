package mtgJson

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const exampleJson = `{
	"meta": {
		"ignore": "me"
	}, 
	"data": {
		"code1": {
			"code": "code1",
			"name": "Example Set 1"
		},
		"code2": {
			"code": "code2",
			"name": "Example Set 2"
		}
	}
}`

type demoStruct struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func TestParseDataFile(t *testing.T) {
	type myDataFile dataFile[string, demoStruct]

	jsonDec := json.NewDecoder(strings.NewReader(exampleJson))
	tok, err := jsonDec.Token()
	require.NoError(t, err)
	assert.Equal(t, json.Delim('{'), tok)

	tok, err = jsonDec.Token()
	require.NoError(t, err)
	assert.Equal(t, "meta", tok)

	var meta any
	err = jsonDec.Decode(&meta)
	require.NoError(t, err)
	tok, err = jsonDec.Token()
	require.NoError(t, err)
	assert.Equal(t, "data", tok)
	tok, err = jsonDec.Token()
	require.NoError(t, err)
	assert.Equal(t, json.Delim('{'), tok)
	tok, err = jsonDec.Token()
	require.NoError(t, err)
	assert.Equal(t, "code1", tok)
	var val demoStruct
	err = jsonDec.Decode(&val)
	require.NoError(t, err)
	assert.Equal(t, "code1", val.Code)
	assert.Equal(t, "Example Set 1", val.Name)

	tok, err = jsonDec.Token()
	require.NoError(t, err)
	assert.Equal(t, "code2", tok)
	err = jsonDec.Decode(&val)
	require.NoError(t, err)
	assert.Equal(t, "code2", val.Code)
	assert.Equal(t, "Example Set 2", val.Name)

	tok, err = jsonDec.Token()
	require.NoError(t, err)
	assert.Equal(t, json.Delim('}'), tok) // End of "data" object

	tok, err = jsonDec.Token()
	require.NoError(t, err)
	assert.Equal(t, json.Delim('}'), tok) // End of root object
}
