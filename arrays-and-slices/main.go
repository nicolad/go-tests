package main
 
func Sum(a [5]int) int {
	sum:=0

	// for i:=0;i<len(a);i++{
	// 	sum+=a[i]
	// }

	for _,v:=range a{
		sum+=v
	}

	return sum
}

func main() {
	var a [5]int
	Sum(a)
}