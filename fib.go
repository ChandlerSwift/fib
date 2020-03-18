package fib

// Fib calculates the n-th Fibonnaci number.
func Fib(n int) int {
	if n > 2 {
		return Fib(n-1) + Fib(n-2)
	}
	return 1 // BUG(chandler): This doesn't correctly handle numbers less than one.
}
