package utils_test

import (
	"encoder/framework/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsJson(t *testing.T) {
	json := `{
				"id": "",
				"file_path": "",
				"status": ""
			}`

	err := utils.IsJson(json)

	require.Nil(t, err)

	json = `roger`

	err = utils.IsJson(json)

	require.Error(t, err)

}
