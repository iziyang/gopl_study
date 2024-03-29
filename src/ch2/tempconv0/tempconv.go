package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZero Celsius = -273.15
	FreezingC    Celsius = 0
	BoilingC     Celsius = 100
)

func CTOF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FTOC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)

}

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100"C
	BoilingF := CTOF(BoilingC)
	fmt.Printf("%g\n", BoilingF-CTOF(FreezingC)
	)
	//fmt.Printf("%g\n", BoilingF-FreezingC)

}
