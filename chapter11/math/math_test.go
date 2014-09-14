package math

import "testing"

type testpair struct {
	values  []float64
	average float64
	min     float64
	max     float64
}

var tests = []testpair{
	{[]float64{1, 2, 3}, 2, 1, 3},
	{[]float64{2, 4, 6, 8, 10}, 6, 2, 10},
	{[]float64{-1, 1}, 0, -1, 1},
}

type Inspector struct {
	t *testing.T
}

func (i *Inspector) expectEqual(values []float64, expected float64, actual float64) {
	if expected != actual {
		i.t.Error("For", values, "Expected ", expected, ", but got ", actual)
	}
}

func TestAverage(t *testing.T) {
	i := Inspector{t: t}
	for _, pair := range tests {
		values := pair.values
		i.expectEqual(values, pair.average, Average(values))
	}
}

func TestMax(t *testing.T) {
	i := Inspector{t: t}
	for _, pair := range tests {
		values := pair.values
		i.expectEqual(values, pair.max, Max(values))
	}
}

func TestMin(t *testing.T) {
	i := Inspector{t: t}
	for _, pair := range tests {
		values := pair.values
		i.expectEqual(values, pair.min, Min(values))
	}
}
