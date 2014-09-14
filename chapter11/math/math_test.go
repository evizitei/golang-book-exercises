package math

import "testing"

var data []float64 = []float64{1, 2, 3}

type Inspector struct {
	t *testing.T
}

func (i *Inspector) expectEqual(expected float64, actual float64) {
	if expected != actual {
		i.t.Error("Expected ", expected, ", but got ", actual)
	}
}

func TestAverage(t *testing.T) {
	i := Inspector{t: t}
	i.expectEqual(2, Average(data))
}

func TestMax(t *testing.T) {
	i := Inspector{t: t}
	i.expectEqual(3, Max(data))
}

func TestMin(t *testing.T) {
	i := Inspector{t: t}
	i.expectEqual(1, Min(data))
}
