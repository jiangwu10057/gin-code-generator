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

func TestReplaceAll(t *testing.T) {
	source := "auth_center_auth_center"
	target := "authcenterauthcenter"
	result := util.ReplaceAll(source, "_", "")
	assert.Equal(t, result, target)
}

func TestLowerFirst(t *testing.T) {
	source := "ASDASJDHASKD"
	target := "aSDASJDHASKD"
	result := util.LowerFirst(source)
	assert.Equal(t, result, target)
}
