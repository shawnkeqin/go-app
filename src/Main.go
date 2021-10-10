package main 

import ("fmt"
"strconv")


var (
 actorName string = "Elizabeth"
 companion string = "Sarah"
 doctorName int = 3
 season int = 11
)

var (counter int = 0)

func main(){
	var i int = 42
	fmt.Printf("%v, %T\n",i,i )

	var j string 
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n",j,j)
	// var theHTTP string = "https://google.com"
//	i := 42
//	var j float32 = 27
	
}

