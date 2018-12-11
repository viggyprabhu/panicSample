package main

import (
	"fmt"
	"panicSample/samplepackage"
	"time"

	"github.com/viggyprabhu/go-handlepanic"
)

type Args struct {
	a int
	b int
}

func main() {

	fmt.Println("Starting go routine")
	args := Args{
		a: 1,
		b: 0,
	}
	handlepanic.CreateInstance(rescue)
	go handlepanic.Handle(panicRoutine, args)
	samplepackage.CreatePanic()

	time.Sleep(10 * time.Second)
	fmt.Println("Finishing main")
}

var panicRoutine = func(args interface{}) interface{} {
	arguments := args.(Args)
	fmt.Println("Ready for panic")
	i := arguments.a / arguments.b

	fmt.Println("Shouldnt come here", i)
	return 0
}

var rescue = func() {
	//TODO: HAndle panic even in defer
	fmt.Println("Called rescue")
	if r := recover(); r != nil {
		fmt.Println("Recovered from ", r)
	}

}
