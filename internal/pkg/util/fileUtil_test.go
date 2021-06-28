package util

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gin-code-generator/internal/pkg/util"
)

func TestCheckFileIsExist(t *testing.T) {
	except := false
	result := util.CheckFileIsExist("a.go")
	assert.Equal(t, except, result)
}
