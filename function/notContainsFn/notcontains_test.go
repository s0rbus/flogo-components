package notcontains

import (
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFnNotContains_Eval(t *testing.T) {
	f := &fnNotContains{}

	v, err := function.Eval(f, "foo", "Bar")
	assert.Nil(t, err)
	assert.True(t, v.(bool))

	v, err = function.Eval(f, "foobar", "foo")
	assert.Nil(t, err)
	assert.False(t, v.(bool))
}
