# [UNDER CONSTRUCTION] estimate_deadline
## Inspired by "The Clean Coder - A Code of Conduct for Professional Programmers, Robert C. Martin"

The object of this script is to estimate the deadline for a task (or a project)
with trivariate input:
1. "O" for most optimistic case
2. "M" for most likely case
3. "P" for most pessimisitic case

[about PERT](https://en.wikipedia.org/wiki/program_evaluation_and_review_technique)

## Usage with source code:
```sh
#the input lengths for "o", "m", and "p" should be same
go run estimate_deadline.go -o 1,2,3 -m 2,5,7 -p 4,9,12
```
### Output
1. mean and standard deviation for individual tasks & all tasks
2. 1-std interval & 2-std interval

### Formula
```py
mu = (O + 4 * M + P) / 6
std = (P - O) / 6
total_mu = sum(mu)
total_std = math.sqrt(sum(std ** 2))
```
Beta distribution is presumed in this formula which makes sense 
because the minimum duration is often more certain than the maximum.

