package main

import "fmt"

func main() {
	var m1 = make(map[string]string)

	m1["cc"] = "ddd"
	m1["dd"] = "ccc"

	if v, ok := m1["zz"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("key err")
	}
}
