package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hailsryan/introClass/method_vs_function/answerkey"
)

func main() {
	s := answerkey.NewSummer()
	rand.Seed(time.Now().Unix())
	executeMethod(s)
	executeFunction(answerkey.Sum)
}

type sumInterface interface {
	Add(int)
	Sum() int
}

func executeMethod(s sumInterface) {
	fmt.Println("Testing methods Add and Sum")
	count := 0
	for i := 0; i < rand.Intn(15); i++ {
		toAdd := rand.Intn(10000)
		count += toAdd
		s.Add(toAdd)
	}
	if s.Sum() != count {
		panic(fmt.Sprintf("Expected %d, Got %d", count, s.Sum()))
	}
	fmt.Println(fmt.Sprintf("Success! Expected %d, Got %d", count, s.Sum()))
}

func executeFunction(f func(a, b int) int) {
	fmt.Println("Testing the function Sum")
	for i := 0; i < rand.Intn(15); i++ {
		a := rand.Intn(100000)
		b := rand.Intn(100000)
		if f(a, b) != a+b {
			panic(fmt.Sprintf("Expected %d, Got %d", a+b, f(a, b)))
		}
		fmt.Println(fmt.Sprintf("Success! Expected %d, Got %d", a+b, f(a, b)))
	}
}
