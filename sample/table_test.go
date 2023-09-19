package sample

import (
	"math"
	"testing"
	"time"
)

func TestTableSimple(t *testing.T) {
	tcs := []struct {
		name string
		val  float64
		want float64
	}{
		{name: "the positive", val: 42, want: 42},
		{name: "the negative", val: -42, want: 42},
		{name: "zero", val: 0, want: 0},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := math.Abs(tc.val)
			if tc.want != got {
				t.Errorf("want: %f, got: %f", tc.want, got)
			}
		})
	}
}

func TestTableSlow(t *testing.T) {
	tcs := []struct {
		name  string
		sleep time.Duration
	}{
		{name: "1s", sleep: 1 * time.Second},
		{name: "2s", sleep: 2 * time.Second},
		{name: "3s", sleep: 3 * time.Second},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			time.Sleep(tc.sleep)
		})
	}
}

func TestTableParallel(t *testing.T) {
	tcs := []struct {
		name  string
		sleep time.Duration
	}{
		{name: "1s", sleep: 1 * time.Second},
		{name: "2s", sleep: 2 * time.Second},
		{name: "3s", sleep: 3 * time.Second},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			t.Log("running ", tc.name)
			time.Sleep(tc.sleep)
		})
	}
}

func TestTableParallelFixed(t *testing.T) {
	tcs := []struct {
		name  string
		sleep time.Duration
	}{
		{name: "1s", sleep: 1 * time.Second},
		{name: "2s", sleep: 2 * time.Second},
		{name: "3s", sleep: 3 * time.Second},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc // rebidding tc so the goroutines will not share it
			t.Parallel()
			t.Log("running ", tc.name)
			time.Sleep(tc.sleep)
		})
	}
}
