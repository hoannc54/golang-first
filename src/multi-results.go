package main

import "fmt"

func swap(x, y, z string) (string, string, string) {
	return z, y, x
}

func main() {
	a, b, c := swap("Nguyen", "Cong", "Hoan")
	fmt.Println(a, b, c)
}
