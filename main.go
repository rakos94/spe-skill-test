package main

import(
	"log"
	"strconv"
	"math"
)

func main(){
	log.Println(narcissistic(153))
}

func narcissistic(value int) bool {
	realValue := value
	slice := strconv.Itoa(value)
	l := len(slice)
	sum := 0
	
	// log.Println(value / 10)
	for value%10 > 0 {
		lastDigit := value%10
		sum += int(math.Pow(float64(lastDigit),float64(l)))
		value = value / 10
	}
	
	if sum == realValue {
		return true
	}

	return false
}