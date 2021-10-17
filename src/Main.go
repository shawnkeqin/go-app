package main 

import ("fmt" 
		"reflect"
		"math"
)

type Doctor struct {
	Number int 
	ActorName string
	Companions []string
	Episodes []string
}

// type Animal struct {
// 	Name string
// 	Origin string
// }

type Animal struct {
	Name string `required max: "100"`
	Origin string
}




type Bird struct {
	Animal
	SpeedKPH 	float32 
	CanFly 	bool 
}




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

func returnTrue() bool {
	fmt.Println("returning true")
	return true 
}

func main(){
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")

	s := "hello Go!"
	for k,v := range s {
		fmt.Println(k,string(v))
	}

	// s := []int{1,2,3}
	// for k,v := range s {
	// 	fmt.Println(k,v)
	// }


// Loop: 
// 	for i := 1; i <= 3; i++{
// 		for j := 1; j <= 3; j++{
// 			fmt.Println(i*j)
// 			if i*j >= 3 {
// 				break Loop
// 			}
// 		}
// 	}



	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0{
	// 		continue 
	// 	}
	// 	fmt.Println(i)
	// }

	// i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// for ; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// 	if i%2 == 0 {
	// 		i /= 2
	// 	} else {
	// 		i = 2*i + 1 
	// 	}
	// }


	for i, j := 0, 0; i < 5; i,j = i+ 1, j + 1 {
		fmt.Println(i)
	}

// 	var i interface{} = [3]int{}
// 	switch i.(type) {
// 	case int:
// 		fmt.Println("i is an int")
// 	case float64:
// 		fmt.Println("i is float64")
// 	case string:
// 		fmt.Println("i is string")
// 	case [3]int:
// 		fmt.Println("i is array")
// 		break
// 		fmt.Println("this will print too ")
// 	default:
// 		fmt.Println("i is another type")
// }



	switch i := 2 + 3; i {
		case 1, 5 , 10:
			fmt.Println("one")
			fallthrough
		case 2, 4, 6:
			fmt.Println("two")
		default:
			fmt.Println("not one or two")
	}


	myNum := 0.1234556
	if math.Abs(myNum / math.Pow(math.Sqrt(myNum), 2) - 1) < 0.001 {
		fmt.Println("these are the same")
	} else {
		fmt.Println("these are different!")
	}

	number := 50
	guess := 101
	if guess < 1 || returnTrue() || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100!")

	}else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
		fmt.Println(number <= guess, number >= guess, number != guess)

	}




	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
	// b := Bird{
	// 	Animal: Animal{Name: "Emu", Origin: "Australia"},
	// 	SpeedKPH: 48,
	// 	CanFly: false,
	// }
	// fmt.Println(b)
	// b := Bird{}
	// b.Name = "Emu"
	// b.Origin = "Australia"
	// b.SpeedKPH = 48
	// b.CanFly = false
	// fmt.Println(b)
	aNurse := struct{name string}{name: "brreg"}
	anotherNurse := aNurse
	anotherNurse.name = "lord"
	fmt.Println(aNurse)
	fmt.Println(anotherNurse)
	aDoctor := Doctor{
		Number: 3,
		ActorName: "bird",
		Companions: []string{
			"bferf",
			"referf",
			"fereg",
		},
	}
	fmt.Println(aDoctor)

	// statePopulations := make(map[string]int)
	// a := make([]int, 3, 100)
	statePopulations := map[string]int{
		"Cali": 392545346,
		"Texas": 335938595,
		"New York": 348394975,
		"Ohio": 3345459045,
	}
	statePopulations["Georgia"] = 10232311
	fmt.Println(statePopulations["Georgia"])
	pop, ok := statePopulations["Ohio"]
	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(pop, ok)
	fmt.Println(len(statePopulations))
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

