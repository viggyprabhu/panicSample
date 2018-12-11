package samplepackage

import (
	"fmt"
	"time"

	"github.com/viggyprabhu/go-handlepanic"
)

type ArgsSample struct {
	a int
	b int
}

func CreatePanic() {
	fmt.Println("Called create Panic of sample package")
	args := ArgsSample{
		a: 1,
		b: 0,
	}
	go handlepanic.Handle(panicRoutine, args)
	time.Sleep(10 * time.Second)
	fmt.Println("Finishing CreatePanic in samplepackage")
}

var panicRoutine = func(args interface{}) interface{} {
	arguments := args.(ArgsSample)
	fmt.Println("Ready for panic in sample package")
	i := arguments.a / arguments.b

	fmt.Println("Shouldnt come here in sample package", i)
	return 0
}
