package main

func main() {
	fibLoop(30)
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fibLoop(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]

}

func fibLoop2(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	prev, curr := 0, 1

	for i := 2; i <= n; i++ {
		// temp := prev + curr
		// prev = curr
		// curr = temp
		prev, curr = curr, prev+curr
	}
	return curr
}
