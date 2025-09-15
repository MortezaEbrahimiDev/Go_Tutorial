package main

func main() {
	x := 10
	p := &x

	println(p)
	println(*p)

	*p = 20
	println(p)
	println(*p)
	println(x)
}
