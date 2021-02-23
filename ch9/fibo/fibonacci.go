package fibo

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	f0, f1 := 0, 1
	for i := 2; i <= n; i++ {
		f0, f1 = f1, f0+f1
	}
	return f1
}
