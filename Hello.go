// Hello

package main

import "fmt"


func f(a,b int){
	ret:=0
	
	for i:=a; i<=b; i++{
		bool ok:=true
		if i==1 {ok=false}
		for j:=2; j*j ; j<=i; ++j{
			if i%j ==0{
				ok=false
			}
		}
		if ok==true {
			ret++
		}
	}
	return ret
}

func main() {

	fmt.Println(f(1,20))
}

/*

import (
	"fmt"
	"sort"
	"strings"
)

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type CaseInsensitive []string

func (c CaseInsensitive) Len() int {
	return len(c)
}

func (c CaseInsensitive) Less(i, j int) bool {
	return strings.ToLower(c[i]) < strings.ToLower(c[j])
}

func (c CaseInsensitive) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func main() {
	apple := CaseInsensitive([]string{"ccc", "aaa", "bbb", "a"})
	sort.Sort(apple)
	fmt.Println(apple)
}


import (
	"fmt"
)

func fact(n int) int {

	var sum int = 1
	for i := 1; i <= n; i++ {
		sum *= i

	}
	return sum
}
func main() {

	fmt.Printf("fact val : %d\n", fact(5))
	//fmt.Print(fact(4))
}



package main

import (
	"fmt"
	"math/rand"
)

func main() {

	a := 40
	b := "hh"
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println("Hello World!")
}


package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite Number is ", rand.Intn(10))

}



package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(math.Pi)

}



package main

import (
	"fmt"
)

func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	a, b := swap("jss", "world")

	fmt.Println("a=", a)
	fmt.Println("b=", b)

}


Naked return

package main

import (
	"fmt"
)

func split(sum int) (x, y int) {

	x = sum + 3
	y = sum - 3

	return
}

func main() {

	fmt.Println(split(40))
}


package main

import (
	"fmt"
)

var java, c, r bool

func main() {
	var i int

	fmt.Println(java, c, r, i)

}


var initialize

package main

import (
	"fmt"
)

var i, j int = 4, 5

func main() {

	var a, b, c = 3, 4, "hh"

	k := false

	fmt.Println(a, b, c, i, j)

	fmt.Println(k)
}


package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBE   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBE, ToBE)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

}


*/
