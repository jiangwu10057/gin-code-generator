package util

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"gin-code-generator/internal/pkg/util"
)

func TestCheckFileIsExist(t *testing.T) {
	except := false
	result := util.CheckFileIsExist("a.go")
	assert.Equal(t, except, result)
}
