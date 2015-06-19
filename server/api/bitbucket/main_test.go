package bitbucket

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	var err error
	assert := assert.New(t)

	data, err := ioutil.ReadFile("test.payload.json")
	assert.NoError(err)

	apireq := ApiReq{}
	err = json.Unmarshal(data, &apireq)
	assert.NoError(err)

	config, err := NewConfigFromFile("test.conf")
	assert.NoError(err)

	buffer := bytes.NewBuffer(nil)
	err = config.Execute(buffer, apireq.Push)
	assert.NoError(err)

	t.Log(buffer.String())
}
