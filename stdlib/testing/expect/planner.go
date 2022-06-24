package expect

import (
	"context"

	"github.com/mvn-trinhnguyen2-dn/flux/codes"
	"github.com/mvn-trinhnguyen2-dn/flux/dependencies/testing"
	"github.com/mvn-trinhnguyen2-dn/flux/internal/errors"
	"github.com/mvn-trinhnguyen2-dn/flux/interpreter"
	"github.com/mvn-trinhnguyen2-dn/flux/runtime"
	"github.com/mvn-trinhnguyen2-dn/flux/semantic"
	"github.com/mvn-trinhnguyen2-dn/flux/values"
)

const pkgpath = "testing/expect"

func init() {
	signature := runtime.MustLookupBuiltinType(pkgpath, "planner")
	runtime.RegisterPackageValue(pkgpath, "planner",
		values.NewFunction("planner",
			signature,
			func(ctx context.Context, args values.Object) (values.Value, error) {
				return interpreter.DoFunctionCallContext(Planner, ctx, args)
			},
			true,
		),
	)
}

func Planner(ctx context.Context, args interpreter.Arguments) (values.Value, error) {
	rules, err := args.GetRequiredDictionary("rules")
	if err != nil {
		return nil, err
	}
	rulesType := rules.Type()

	if keyType, err := rulesType.KeyType(); err != nil {
		return nil, err
	} else if got := keyType.Nature(); got != semantic.String {
		return nil, errors.Newf(codes.FailedPrecondition, "key type must be a string, got %s", got)
	}

	if valueType, err := rulesType.ValueType(); err != nil {
		return nil, err
	} else if got := valueType.Nature(); got != semantic.Int {
		return nil, errors.Newf(codes.FailedPrecondition, "value type must be an int, got %s", got)
	}

	rules.Dict().Range(func(key, value values.Value) {
		if err != nil {
			return
		}
		err = testing.ExpectPlannerRule(ctx, key.Str(), int(value.Int()))
	})
	if err != nil {
		return nil, err
	}
	return values.Void, nil
}
