package main

// 斐波那契数列
func fib(n int) int {
	const MOD = 1e9 + 7
	if n <= 1 {
		return n
	}
	a := 0
	b := 1
	c := 0
	for i := 2; i <= n; i++ {
		c = (a + b) % MOD
		a = b
		b = c
	}
	return c
}
