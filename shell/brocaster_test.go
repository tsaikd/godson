package shell

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_brocaster(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(assert)
	require := require.New(t)
	require.NotNil(require)

	bcaster := newBrocaster()
	buffer1 := bytes.Buffer{}
	buffer2 := bytes.Buffer{}

	bcaster.AddOutput(&buffer1)
	bcaster.AddOutput(&buffer2)

	n, err := bcaster.Write([]byte("test string"))
	require.NoError(err)
	require.NotZero(n)
	require.NotZero(buffer1.Len())
	require.NotZero(buffer2.Len())
	require.Equal(buffer1.String(), buffer2.String())
}
