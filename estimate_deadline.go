package main

import (
	"fmt"

	"github.com/khb7840/estimate_deadline/calc"
)

/*
 PERT implementation
 This script is to estimate the deadline for a task (or a project)
 with trivariate input
	1) "O" for most optimistic case
	2) "M" for most likely case
	3) "P" for most pessimisitic case
 Usage with source code:
	go run estimate_deadline.go -o 1,2,3 -m 2,5,7 -p 4,9,12
	(the input lengths for "o", "m", and "p" should be same)
 Output:
	1) mean and standard deviation for individual tasks & all tasks
	2) 1-std interval & 2-std interval
 Formula:
	mu = (O + 4M + P) / 6
	std = (P - O) / 6
	total_mu = sum(mu)
	total_std = sqrt(sum(std^2))
 PERT presumes that this approximates a beta distribution.
 This makes sense since the minimum duration for a task
 is often much more certain than the maximum.
*/

func main() {
	a, _ := calc.GetMu(1, 5, 10)
	fmt.Print(a)
}

/*
func main() {

}
*/
