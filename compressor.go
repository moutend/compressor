package compressor

import (
	"time"
)

type Parameter struct {
	Attack    time.Duration
	Release   time.Duration
	Threshold float64
	Ratio     float64
}

func (p *Parameter) Apply(input []float64) []float64 {
	length := len(input)
	output := make([]float64, length)

	for i := 0; i < length; i++ {
		if input[i] >= p.Threshold {
			output[i] = p.Threshold + (input[i]-p.Threshold)*p.Ratio
		} else if input[i] <= -p.Threshold {
			output[i] = -p.Threshold - (input[i]+p.Threshold)*p.Ratio
		} else {
			output[i] = input[i]
		}
	}
	return output
}
