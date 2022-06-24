package universe_test

import (
	"testing"
	"time"

	"github.com/mvn-trinhnguyen2-dn/flux"
	"github.com/mvn-trinhnguyen2-dn/flux/array"
	"github.com/mvn-trinhnguyen2-dn/flux/arrow"
	"github.com/mvn-trinhnguyen2-dn/flux/execute"
	"github.com/mvn-trinhnguyen2-dn/flux/execute/executetest"
	"github.com/mvn-trinhnguyen2-dn/flux/memory"
	"github.com/mvn-trinhnguyen2-dn/flux/querytest"
	"github.com/mvn-trinhnguyen2-dn/flux/stdlib/influxdata/influxdb"
	"github.com/mvn-trinhnguyen2-dn/flux/stdlib/universe"
)

func TestCount_NewQuery(t *testing.T) {
	tests := []querytest.NewQueryTestCase{
		{
			Name: "from with range and count",
			Raw:  `from(bucket:"mydb") |> range(start:-4h, stop:-2h) |> count()`,
			Want: &flux.Spec{
				Operations: []*flux.Operation{
					{
						ID: "from0",
						Spec: &influxdb.FromOpSpec{
							Bucket: influxdb.NameOrID{Name: "mydb"},
						},
					},
					{
						ID: "range1",
						Spec: &universe.RangeOpSpec{
							Start: flux.Time{
								Relative:   -4 * time.Hour,
								IsRelative: true,
							},
							Stop: flux.Time{
								Relative:   -2 * time.Hour,
								IsRelative: true,
							},
							TimeColumn:  "_time",
							StartColumn: "_start",
							StopColumn:  "_stop",
						},
					},
					{
						ID: "count2",
						Spec: &universe.CountOpSpec{
							SimpleAggregateConfig: execute.DefaultSimpleAggregateConfig,
						},
					},
				},
				Edges: []flux.Edge{
					{Parent: "from0", Child: "range1"},
					{Parent: "range1", Child: "count2"},
				},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			querytest.NewQueryTestHelper(t, tc)
		})
	}
}

func TestCountOperation_Marshaling(t *testing.T) {
	data := []byte(`{"id":"count","kind":"count"}`)
	op := &flux.Operation{
		ID:   "count",
		Spec: &universe.CountOpSpec{},
	}

	querytest.OperationMarshalingTestHelper(t, data, op)
}

func TestCount_Process(t *testing.T) {
	testCases := []struct {
		name string
		data func() *array.Float
		want int64
	}{
		{
			name: "zero",
			data: func() *array.Float {
				return arrow.NewFloat([]float64{0, 0, 0}, nil)
			},
			want: 3,
		},
		{
			name: "nonzero",
			data: func() *array.Float {
				return arrow.NewFloat([]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, nil)
			},
			want: 10,
		},
		{
			name: "empty",
			data: func() *array.Float {
				return arrow.NewFloat(nil, nil)
			},
			want: 0,
		},
		{
			name: "with nulls",
			data: func() *array.Float {
				b := arrow.NewFloatBuilder(nil)
				defer b.Release()
				b.AppendValues([]float64{0, 1, 2, 3}, nil)
				b.AppendNull()
				b.AppendValues([]float64{5, 6}, nil)
				b.AppendNull()
				b.AppendValues([]float64{8, 9}, nil)
				return b.NewFloatArray()
			},
			want: 10,
		},
		{
			name: "only nulls",
			data: func() *array.Float {
				b := arrow.NewFloatBuilder(nil)
				defer b.Release()
				b.AppendNull()
				b.AppendNull()
				return b.NewFloatArray()
			},
			want: 2,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			data := tc.data()
			defer data.Release()

			executetest.AggFuncTestHelper(
				t,
				new(universe.CountAgg),
				data,
				tc.want,
			)
		})
	}
}
func BenchmarkCount(b *testing.B) {
	data := arrow.NewFloat(NormalData, &memory.ResourceAllocator{})
	executetest.AggFuncBenchmarkHelper(
		b,
		new(universe.CountAgg),
		data,
		int64(len(NormalData)),
	)
}
