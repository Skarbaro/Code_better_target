package main

import (
	"fmt"
	"log"
)

func hello() {
	fmt.Println("hello")
}

func main() {
	fmt.Println("hello test")
}

func test() {
	log.Println("added line")
	log.Println("added line2")
	log.Println("some new text")
	fmt.Println("hello 1")
	fmt.Println("new line")
	fmt.Println("new line2")
	fmt.Println("new line3")
}

func test2() {
	log.Println("added anothr method")
}

func testNewFunc() {
	log.Println("test new func")
}

func test3() {
	log.Println("test3")
	fmt.Println("new line in conflict")
	fmt.Println("new line in conflict2")
}
