package maths

var Myname string = "James"

func Factorial(n uint64) uint64 {
	f := n // Holds the first value of n
	var i uint64
	for i = 1; i < n; i++ {
		f *= uint64(n - i)
	}
	return f
}
