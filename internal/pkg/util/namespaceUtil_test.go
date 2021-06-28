package util

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gin-code-generator/internal/pkg/util"
)

func TestGetNameSpace(t *testing.T) {
	except := "gin-code-generator"
	result, _ := util.GetNameSpace("/golang/")
	assert.Equal(t, except, result)
}
