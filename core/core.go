package core

//Fibonacci sequence
func Fibonacci(x int) int {
	var f, p int
	if x == 0 {
		return 0
	}
	for x, f, p = x-2, 1, 0; x >= 0; x-- {
		f, p = f+p, f
	}
	return f
}
