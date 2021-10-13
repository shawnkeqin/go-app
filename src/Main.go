package main 

import ("fmt"
)


// var (
//  actorName string = "Elizabeth"
//  companion string = "Sarah"
//  doctorName int = 3
//  season int = 11
// )

// var (counter int = 0)

const a int16 = 27 

const (
	x = iota 
	y = iota
	z = iota 
)

const (
	_ = iota + 5
	catSpz
	dogSpz
	snakeSpz
)



func main(){
	// a := make([]int, 3, 100)
	a := []int{}
	a = append(a,1)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	// a = append(a,2,3,4,5)
	a = append(a, []int{2,3,4,5}...)
	g := []int{1,2,3,4,5}
	fmt.Println(g)
	k := append(g[:2], g[3:]...)
	fmt.Println(k)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	// var spzType int
	// fmt.Printf("%v\n", spzType == catSpz)
	// fmt.Printf("%v\n", catSpz)
	x := []int{1,2,3,4,5,6,7,8,9,10}
	y := x[:]
	z := x[3:]
	m := x[:6]
	n := x[3:6]
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(m)
	fmt.Println(n)
	var idtMatrix [3][3]int = [3][3]int{ [3]int{1,0,0}, [3]int{0,0,1}}
	fmt.Println(idtMatrix)
	// a := [...]int{1,2,3}
	// a := []int{1,2,3}
	// b := a
	b := a
	// b[1] = 5 
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Students: %v\n", cap(b))
	var students [3]string
	// grades := [3]int{97,85,93}
	// fmt.Printf("Students: %v", students)
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"
	gradesx := [...]int{97,85,93}
	fmt.Printf("Students: %v\n", students[1])
	fmt.Printf("Students: %v\n", len(students))

	fmt.Printf("Grade: %v", gradesx)
	// const a int = 42
	// const b string = "foo"
	// const c float32 = 3.14
	// const d bool = true
	// const e int = 39
	// fmt.Printf("%v, %T\n", a + e, a + e)
	// fmt.Printf("%v\n", a)
	// fmt.Printf("%v\n", b)
	// fmt.Printf("%v\n", c)
	// fmt.Printf("%v\n", d)
	// fmt.Printf("%v\n, %T\n", x, x)
	// fmt.Printf("%v\n, %T\n", y, y)
	// fmt.Printf("%v\n, %T\n", z, z)
	// var r rune= 'a'
	// fmt.Printf("%v, %T\n", r, r)
	// var n complex128 = complex(5,12)
	// fmt.Printf("%v, %T\n", n,n)

	// s := "this is a string"
	// fmt.Printf("%v, %T\n", string(s[2]), s[2])
	// a := 10.2
	// b := 3.7
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)
	// n := 42
	// var n uint16 = 42
	// fmt.Printf("%v, %T\n",n,n)
	// a := 10
	// b := 3
	// fmt.Println(a+b)
	// fmt.Println(a-b)
	// fmt.Println(a*b)
	// fmt.Println(a+b)
	// fmt.Println(a%b)
	// var a int = 10
	// var b int8 = 3
	// fmt.Println(a + int(b))
	// a := 10
	// b := 3
	// fmt.Println(a & b)
	// fmt.Println( a | b)
	// fmt.Println(a ^  b)
	// fmt.Println(a &^ b)
	// a := 8
	// fmt.Println(a << 3)
	// fmt.Println(a >> 3)
	// var n bool = true 

	// n := 1 == 1
	// m := 1 == 2
	// fmt.Printf("%v, %T\n",n,n)
// 	var i int = 42
// 	fmt.Printf("%v, %T\n",i,i )

// 	var j string 
// 	j = strconv.Itoa(i)
// 	fmt.Printf("%v, %T\n",j,j)
	// var theHTTP string = "https://google.com"
//	i := 42
//	var j float32 = 27
	
}

