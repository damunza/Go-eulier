package main

import (
	"fmt"
	"math/rand"
	"time"
	"math"
)
// dont use caps in describing types unless they are to be exported 
type pythog struct{
	a [3] int
	sum float64
}

func main(){
	if calcPythog(generteRandom(1000),generteRandom(1000)) == 1000.0{

	}
	
}

func generteRandom (end int) int{
	/*
	this is a function to generate a random numbers between 0 and 1000
	*/
	rand.Seed(time.Now().UnixNano()) // prevents the repetition of a similar value by ensiring a new one is generated each time rand is called
	start := 1
	x := rand.Intn(end-start)+start
	return x
}

func calcPythog (h, k int) float64{
	/*
	function to calculate the posible c value of our Pythog struct
	*/
	var x pythog
	x.a[0] = h
	x.a[1] = k
	a := float64(x.a[0])
	b := float64(x.a[1])
	c := math.Pow(a, 2.0)
	d := math.Pow(b, 2.0)
	e := c + d
	x.a[2] =int(math.Sqrt(e))
	x.sum = 0
	for _, i := range x.a{
		x.sum += float64(i)
	}
	return x.sum

}
