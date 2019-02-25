// make a new type MyInt that is an integer
type MyInt int

// attach a method to MyInt to square a number
func(n MyInt) sqr MyInt {
  return n*n
}

// make a new MyInt-type variable called "number" and set it to 5
var number MyInt = 5

// use the sqr() method
var square = number.sqr()

