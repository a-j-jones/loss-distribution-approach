package main

import (
	"flag"
	"fmt"
	"sort"
	"sync"
	"time"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	// Command line flag variables
	var distribution string
	var lambda float64
	var mu float64
	var sigma float64
	var simulations int
	var iterations int

	// Parse command line arguments
	flag.StringVar(&distribution, "distribution", "lognormal", "Distribution type (lognormal/normal/...)")
	flag.Float64Var(&lambda, "lambda", 1.2, "Lambda value")
	flag.Float64Var(&mu, "mu", 10, "Mu value")
	flag.Float64Var(&sigma, "sigma", 1, "Sigma value")
	flag.IntVar(&simulations, "simulations", 1000, "Number of simulations")
	flag.IntVar(&iterations, "iterations", 1, "Number of iterations")
	flag.Parse()

	// Get uint64 random seed based on the current epoch time
	seed := uint64(time.Now().UnixNano())
	src := rand.New(rand.NewSource(seed))

	var wg sync.WaitGroup
	wg.Add(int(iterations))

	for sim := 0; sim < iterations; sim++ {
		go func() {
			defer wg.Done()

			frequency := distuv.Poisson{
				Lambda: lambda,
				Src:    src,
			}

			var severity distuv.Rander
			switch distribution {
			case "lognormal":
				severity = distuv.LogNormal{
					Mu:    mu,
					Sigma: sigma,
					Src:   src,
				}

			case "normal":
				severity = distuv.Normal{
					Mu:    mu,
					Sigma: sigma,
					Src:   src,
				}

			default:
				fmt.Println("Invalid distribution type")
				return
			}

			data := make([]float64, simulations)
			for i := range data {
				events := frequency.Rand()
				for j := 0; j < int(events); j++ {
					data[i] += severity.Rand()
				}
			}
			sort.Float64s(data)

			// Calculate the mean and standard deviation of the data

			percentile := stat.Quantile(0.999, stat.Empirical, data, nil)
			fmt.Printf("%0.2f\n", percentile)
		}()
	}

	wg.Wait()
}
