package main

import (
	"fmt"
	"sort"
	"time"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	const simulations int32 = 1_000_000

	// Get uint64 random seed based on the current epoch time
	seed := uint64(time.Now().UnixNano())
	src := rand.New(rand.NewSource(seed))

	frequency := distuv.Poisson{
		Lambda: 1.2,
		Src:    src,
	}

	severity := distuv.LogNormal{
		Mu:    10,
		Sigma: 1,
		Src:   src,
	}

	data := make([]float64, simulations)
	prev_loss := 0.0
	comparison := 0.0
	for i := range data {
		events := frequency.Rand()
		for j := 0; j < int(events); j++ {
			data[i] += severity.Rand()
		}

		if i%5000 == 0 {
			datacopy := make([]float64, i+1)
			copy(datacopy, data)
			sort.Float64s(datacopy)
			loss := stat.Quantile(0.999, stat.Empirical, datacopy, nil)
			if prev_loss != 0.0 {
				comparison = ((loss - prev_loss) / prev_loss) * 100
				fmt.Printf("%0.2f\n", comparison)
			}
			prev_loss = loss
		}
	}
	sort.Float64s(data)

	// Calculate the mean and standard deviation of the data

	percentile := stat.Quantile(0.999, stat.Empirical, data, nil)
	fmt.Printf("%0.2f\n", percentile)
}
