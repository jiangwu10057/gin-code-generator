package util

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"gin-code-generator/internal/pkg/util"
)

func TestContains(t *testing.T) {
	arr := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	item := "a"
	i := "i"

	assert.Equal(t, true, util.Contains(arr, item))
	assert.Equal(t, false, util.Contains(arr, i))

}
