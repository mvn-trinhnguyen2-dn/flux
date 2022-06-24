package http

import (
	"context"
	"net/url"

	"github.com/mvn-trinhnguyen2-dn/flux/interpreter"
	"github.com/mvn-trinhnguyen2-dn/flux/runtime"
	"github.com/mvn-trinhnguyen2-dn/flux/values"
)

func init() {
	runtime.RegisterPackageValue("http", "pathEscape", pathEscapeFunc)
}

const inputStringArg = "inputString"

var pathEscapeFunc = values.NewFunction(
	"pathEscape",
	runtime.MustLookupBuiltinType("http", "pathEscape"),
	func(ctx context.Context, args values.Object) (values.Value, error) {
		return interpreter.DoFunctionCall(PathEncode, args)
	},
	false,
)

func PathEncode(args interpreter.Arguments) (values.Value, error) {
	inputString, err := args.GetRequiredString(inputStringArg)
	if err != nil {
		return nil, err
	}
	return values.NewString(url.PathEscape(inputString)), nil
}
