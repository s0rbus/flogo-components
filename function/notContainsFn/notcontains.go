package notcontains

import (
	"fmt"
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnNotContains{})
}

type fnNotContains struct {
}

func (fnNotContains) Name() string {
	return "notcontains"
}

func (fnNotContains) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnNotContains) Eval(params ...interface{}) (interface{}, error) {
	str, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.notcontains function first parameter [%+v] must be a string", params[0])
	}

	substr, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.notcontains function second parameter [%+v] must be a string", params[1])
	}

	return !strings.Contains(str, substr), nil
}
