package main

func main() {
	//way 1
	var a, b, c, d int
	a = 1
	b = 2
	c = 3
	d = 4

	//way 2
	e, f, g, h := 1, 2, 3, 4

	//way 3
	var i, j, k, l = 1, 2, 3, 4

	//way 4
	var (
		m int    = 101
		n string = "Hello"
		o bool   = false
		p int
	)

	print(a, b, c, d)
	print(e, f, g, h)
	print(i, j, k, l)
	print(m, n, o, p)

}