package util

import (
	"gin-code-generator/internal/pkg/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpperFirst(t *testing.T) {
	source := "adsadhajks"
	target := "Adsadhajks"
	result := util.UpperFirst(source)
	assert.Equal(t, result, target)
}

func TestCase2Camel(t *testing.T) {
	source := "auth_center"
	target := "AuthCenter"
	result := util.Case2Camel(source)
	assert.Equal(t, result, target)

	source = "auth"
	target = "Auth"
	result = util.Case2Camel(source)
	assert.Equal(t, result, target)
}
