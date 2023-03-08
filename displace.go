package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a float64, v_o float64, s_o float64) func(float64) float64 {
	ret := func(time float64) float64 {
		disp := (0.5)*a*math.Pow(time, 2) + v_o*time + s_o
		return disp
	}
	return ret
}

func displace() {
	var a, v_o, s_o, time float64
	fmt.Println("Enter the acceleration value")
	fmt.Scan(&a)
	fmt.Println("Enter the initial velocity value")
	fmt.Scan(&v_o)
	fmt.Println("Enter the initial displacement")
	fmt.Scan(&s_o)
	f1 := GenDisplaceFn(a, v_o, s_o)
	fmt.Println("Enter the time value")
	fmt.Scan(&time)
	disp := f1(time)
	fmt.Println("The Displacement is: ", disp)
}
