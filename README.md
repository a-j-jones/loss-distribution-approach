# Loss Distribution Approach - Monte-Carlo simulation:

A basic Go implementation [LDA](http://www.thierry-roncalli.com/download/lda.pdf), used for calculating operational risk.


## Overview

The code executes a Monte-Carlo simulation, where for each simulation:
 1. The frequency of losses for a given year is taken from the Poisson distribution.
 2. The severity (£) of each of those losses is taken from a given distribution (default: lognormal)
 3. The total loss for the year is added to an array of simulated losses.

Once the Monte-Carlo simulation is complete, the 99.9th percentile of losses (1 in 1000 year losses) is printed to the console.

## Build

With Go installed on your machine
```cmd
go build main.go
```

## Execution

After building the exe, you can run:
```cmd
main --help
```
Which will show the optional parameters of the tool, an example execution would be:

```cmd
main --lambda 1.2 --mu 10 --sigma 1 --simulations 2000000
```

This project does not handle the fitting of distributions to actual data, only the execution of the Monte-Carlo simulation.