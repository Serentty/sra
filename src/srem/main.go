package main

import (
"fmt"
)

func main() {
	fmt.Println("SRA Emulator")
	fmt.Println("Creating Virtual CPU...")
	c := CPU{}
	fmt.Println("Running Test...")
	test(&c)
}

func test(c *CPU) {
	c.store(5, 42)
	fmt.Println(c.load(5))
}