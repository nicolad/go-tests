package main

func Sum(a []int) int {
	sum := 0

	for _, v := range a {
		sum += v
	}

	return sum
}

func SumAll(a ...[]int) []int {
	var sums []int
	for _, v := range a {
		sums = append(sums, Sum(v))
	}

	return sums
}

func SumAllTails(a ...[]int) []int {
	var sums []int
	for _, v := range a {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			tail := v[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

func main() {}
