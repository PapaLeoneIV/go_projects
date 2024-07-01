package defe

import (
	"fmt"
)

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func Wakeuproutine(name string) { 
	trace("Wakeup Routine")
	defer untrace("Wakeup Routine")
	fmt.Println("GoodMorning :", name)
}

func Simpleroutine() {
	trace("Simpleroutine")
	defer untrace("Simpleroutine")
	Wakeuproutine("Riccardo")

}