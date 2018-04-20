// Package f64stat calculates rudamentary statistics for a single real random variable
package f64stat

import (
	"math"
)

// Stat is a container for tracking the random variable
type Stat struct {
	Sum   float64
	Sum2  float64
	Vmin  float64
	Vmax  float64
	Count float64
	Last  float64
}

// Reset zeros out the accumilated data
func (f *Stat) Reset() *Stat {
	f.Count = 0.0
	return f
}

// Add incorporates a sample into the statistical calculations
// It returns a pointer to the
func (f *Stat) Add(value float64) *Stat {
	if f.Count <= 0.0 {
		f.Sum = 0.0
		f.Sum2 = 0.0
		f.Vmin = value
		f.Vmax = value
		f.Count = 0.0
	}
	f.Last = value
	f.Count += 1.0
	f.Sum += value
	f.Sum2 += value * value
	if value < f.Vmin {
		f.Vmin = value
	}
	if value > f.Vmax {
		f.Vmax = value
	}
	return f
}

// Ave returns the arithmetic mean of the random variable samples
func (f *Stat) Ave() float64 {
	if f.Count <= 0.0 {
		return 0.0
	}
	return f.Sum / f.Count
}

// RMS returns the root mean squared average of the samples
func (f *Stat) RMS() float64 {
	if f.Count <= 0.0 {
		return 0.0
	}
	return math.Sqrt(f.Sum2 / float64(f.Count))
}

//Stddev returns standard deviation of the samples
func (f *Stat) Stddev() float64 {
	if f.Count <= 0.0 {
		return 0.0
	}
	ave := f.Sum / float64(f.Count)
	return math.Sqrt(f.Sum2/float64(f.Count) - ave*ave)
}

// Min returns the minimum sample value
func (f *Stat) Min() float64 {
	return f.Vmin
}

// Max returns the maximum sample value
func (f *Stat) Max() float64 {
	return f.Vmax
}

// New creates a new random variable ready to add samples
func New() *Stat {
	return &Stat{}
}
