// // addition of two numbers
// // var a int =10 ; static varable declaration
// package main

// import "fmt"

// const (
// 	C  = "ketan"
// 	PI = 3.14
// )

// //f:=10  global variable cannot be declared using':='

// func main() {

// 	// n1 := 10 //inferred data type
// 	// n2 := 20
// 	// var sum, sum1, sum2, sum3, sum4 int
// 	// //Addition
// 	// sum = n1 + n2
// 	// fmt.Println("sum is", sum)
// 	// fmt.Println("sum n1 n2 is ", sum, n1, n2)
// 	// //Substraction
// 	// sum1 = n2 - n1
// 	// fmt.Println("sum is", sum1)
// 	// fmt.Println("sum n1 n2 is ", sum1, n1, n2)
// 	// //Multiplication
// 	// sum2 = n2 * n1
// 	// fmt.Println("sum is", sum2)
// 	// fmt.Println("sum n1 n2 is ", sum2, n1, n2)
// 	// //Division
// 	// sum3 = n2 / n1
// 	// fmt.Println("sum is", sum3)
// 	// fmt.Println("sum n1 n2 is ", sum3, n1, n2)
// 	// //Modulus
// 	// sum4 = n2 % n1
// 	// fmt.Println("sum is", sum4)
// 	// fmt.Println("sum n1 n2 is ", sum4, n1, n2)

// 	//day 2

// 	// const PI = 3.34
// 	// //PI = PI + 1  //will throw the error as changin the value of const "not addressable error/or apping error"
// 	// fmt.Println("PI is ", PI)

// 	//Typed and Untyped constant
// 	//Typed which is statically declared via type name
// 	//Untyped is inferred using '='
// 	// const(
// 	// 	A = 10
// 	// 	B = 20
// 	// )"Declaring multiple constants at simultaneously"

// 	// fmt.Println("Const ", C)

// 	// //Print f function

// 	// const L = 10
// 	// const W = 5
// 	// var area int
// 	// area = L * W
// 	// fmt.Printf("Area of the triangle:%d", area) //p in small will give undefined
// 	// // var s float32 = 123456789012
// 	// // s = s + 9
// 	// // fmt.Printf("\nArea of the triangle:%f", s)
// 	// //go will throw error in c ase of float overflow ,while in C goes on -ve axis like 255+3 will result in -3

// 	// var a uint8 // Analogy if need 1 lit of water we do not use the 10 lit container as it is inefficient same as byte  rune as int 32
// 	// a = 30      //will show error of overflowif above 255 in go (only +ve)
// 	// fmt.Printf("\nValue of Unsigned int:%d", a)
// 	// //uintptr unsigned ptr varaible

// 	// var student1 string = "John"
// 	// var student2 = "Jane"
// 	// x := 10
// 	// fmt.Printf("\nValue of variables :%s and %s and %T\n", student1, student2, x) // F in Printf stands for format specifier
// 	// fmt.Println("Value of variables :", student1)
// 	// fmt.Println("Value of variables :", student2) //default prints in next line '\n' not supported

// 	// fmt.Print("Value of variables  :", x) //print the text and value but not in new line even if two print statement are used

// 	// f := "foo"
// 	// fmt.Printf("\n%T", f)

// 	// p, y, z, k := 0, 2.2, "ketan", true

// 	// fmt.Println("the values variables are ", p, y, z, k)

// 	// var s, e = 6, "Hello"
// 	// c, d := 7, "world"
// 	// fmt.Println("the values variables are ", s, e, c, d)

// 	// var (
// 	// 	o     = "polawar"
// 	// 	v int = 1
// 	// )
// 	// fmt.Println("the values variables are ", o, v)
// 	// fmt.Printf("the vales variable are %s ,%d\n", o, v)
// 	//var fname, lname, Age, Branch, College string = "Ketan", "Polawar", "20", "Data Science", "RCOEM"
// 	// fmt.Printf("Hey enter your fname\n")
// 	// fmt.Scanf("%s", &fname)
// 	// fmt.Printf("Hey enter your lname\n")
// 	// //fflush use karan padega else woh 2nd input nahi lega
// 	// fmt.Scanln(&lname)

// 	// fmt.Printf("name is %s,%s ", fname, lname)
// 	// var i, b int = 10, 20
// 	// fmt.Println(i, b)
// 	// fmt.Println(i + b)
// 	// fmt.Println(fname + lname)

// 	// fmt.Println("Name :", fname, lname)
// 	// fmt.Println("Age:", Age)
// 	// fmt.Println("Branch:", Branch)
// 	// fmt.Println("CollegeName:", College)

// 	// fmt.Printf("\nName : '%s %s'\tAge:'%s'\n", fname, lname, Age)
// 	// fmt.Printf("Branch:'%s'\tCollege:'%s'\n", Branch, College)

// 	// const name, dept = "DASCA", "DS"
// 	// fmt.Printf("%s is a %s Portal", name, dept)

// 	// var str = "Ketan_Polawr"
// 	// fmt.Printf("The string is %s \n", str) //string
// 	// var num1 int = 20
// 	// fmt.Printf("The string is %d \n", num1) //int
// 	// fmt.Printf("The string is %b \n", num1) //prints binary value
// 	// var num3 float32 = 1.22
// 	// fmt.Printf("The string is %g \n", num3) //float num %f can also be employeed
// 	// var num2 = 4 + 1i
// 	// fmt.Printf("The string is %e \n", num2) //complex number

// 	// fmt.Printf("The string is %v,%v,%v,%v \n", str, num3, num2, num1) //it uses the original type and prints the values
// 	//C prints the garbage values in case of undefined values  but In Go prints the defined value
// 	//default values
// 	//int 0
// 	//float 0.0
// 	//string nil or blank space
// 	//Boolean False
// 	//Format specifiers in go
// 	//%d integer
// 	//%f float
// 	//%T type of varable

// 	//var can be used inside ans outside the variable , also assignemet and declaration can be done seperately
// 	//':=' cannot be used to declare the global declaration ,also assignemet and declaration can be done simultaneously
// 	//literals (something having fixed value)are always rvalue x=10 :: x is lvalue and 10 is rvalue
// 	//integer literals has decimal(10 starts without the 0) octal(8 starts 0,0o,0O) hexa(16  starts with 0x or 0X) binary(2 starts with 0b or 0B)
// 	// fmt.Println(15 == 017)
// 	// fmt.Println(15 == 0xF)
// 	// println(13 == 0xD)
// 	// var a int = 10
// 	// a++
// 	// fmt.Println(a)
// 	// fmt.Printf("hex of %d is %0x\n", a, a)
// 	// fmt.Printf("hex of %d is %0b\n", a, a)
// 	// fmt.Printf("hex of %d is %0o\n", a, a)
// 	// pre increment ie --a,++a is not defined in GO post increment is defined  C langauge lang supports the both
// 	//e stands as an exponent
// 	// a = 1.23e5
// 	// fmt.Print(a)
// 	//%T is type %t is for boolean
// 	// fmt.Printf("\nExpression 15==0xf is '%t' and has type '%T'", 15 == 0xf, 15 == 0xf)
// 	//\a,\t,\n are called escape specifier
// 	// fname := "Ketan"
// 	// lname := "Polawar"
// 	// fmt.Print(fname, "\n")
// 	// fmt.Print(lname, "\n")
// 	// fmt.Print("my name is ", fname, " and last name is ", lname)
// 	// fmt.Printf("\n%s\n", fname)
// 	// fmt.Printf("%s\n", lname)
// 	// fmt.Printf("%s  %s\n", fname, lname)
// 	// //print by defaults inserts the space between the literals if they are not string type else either any one is string  explictly mention space-
// 	// fmt.Print(fname, 20)
// 	//var i = 15
// 	//var txt = "ketan"
// 	// fmt.Printf("%v\n", i)  //will give Ketan
// 	// fmt.Printf("%#v\n", i) //will give "ketan"
// 	// fmt.Printf("%v%%\n", i)
// 	// fmt.Printf("%T\n", i)
// 	// fmt.Printf("\"ketan\"")
// 	// fmt.Printf("'Ketan'")
// 	// fmt.Printf("%b\n", i) //binary
// 	// fmt.Printf("%d\n", i) //prints the sign also
// 	// fmt.Printf("%+d\n", i)
// 	// fmt.Printf("%o\n", i)
// 	// fmt.Printf("%O\n", i)
// 	// fmt.Printf("%x\n", i)
// 	// fmt.Printf("%X\n", i)
// 	// fmt.Printf("%#x\n", i)
// 	// fmt.Printf("%4d\n", i)  //writes in 4 digit
// 	// fmt.Printf("%-4d\n", i) //left justify
// 	// fmt.Printf("%04d\n", i)
// 	//output
// 	//1111
// 	// 15
// 	// +15
// 	// 17
// 	// 0o17
// 	// f
// 	// F
// 	// 0xf
// 	//   15
// 	// 15
// 	// 0015
// 	// fmt.Printf("%s\n", txt)
// 	// fmt.Printf("%q\n", txt)
// 	// fmt.Printf("%8s\n", txt)
// 	// fmt.Printf("%-8s\n", txt)
// 	// fmt.Printf("%x\n", txt)
// 	// fmt.Printf("% x\n", txt)
// 	//**************************************************//
// 	// var (
// 	// 	marks1     int
// 	// 	marks2     int
// 	// 	marks3     int
// 	// 	roll       int
// 	// 	percentage int
// 	// 	sum        int
// 	// 	name       string
// 	// 	branch     string
// 	// )

// 	// fmt.Printf("\n**Please Enter The Details**\n")
// 	// fmt.Println("**Please Enter Your Name,Branch,Id**")
// 	// fmt.Scan(&name, &branch, &roll)
// 	// fmt.Println("**Please Enter Your marks1 marks2 marks 3**")
// 	// fmt.Scan(&marks1, &marks2, &marks3)
// 	// sum = marks1 + marks2 + marks3
// 	// percentage = (sum * 100) / 300
// 	// print("\n\n")
// 	// println("**SHRI RAMDEOBABA COLLEGE OF ENGINEERING AND MANAGEMENT**")
// 	// print("                Student Grade Card")
// 	// fmt.Printf("\nName : %s\tId:'%d'\n", name, roll)
// 	// fmt.Printf("Branch:'%s'\tscore:'%d/300'\n", branch, sum)
// 	// println("DBMS:", marks1)
// 	// println("OS:", marks2)
// 	// println("DHLV:", marks3)
// 	// fmt.Printf("Percentage:%d%%", percentage)
// 	// if percentage > 40 {
// 	// 	print("\nCongulations You Have Passed the Exam !!")
// 	// } else {
// 	// 	print("\nYou have Failed ,Better Luck Next Time :) ")
// 	// }
// 	//**************************************************************//
// 	// integer := 23
// 	// fmt.Printf("%T %T\n", integer, &integer)
// 	// fmt.Print(integer, &integer)

// 	// number := 1234.575883939
// 	// fmt.Printf("number:%5.2f", number)

// 	// fmt.Printf("%*s\n", 40, "text")//width of output is 40  can be specified "%040s"

// 	// ch := 'A'
// 	// bch := "ketan"
// 	// fmt.Printf("%T %T\n", ch, bch)

// 	// var val byte = 'A' //character define kar sakte hai with the help of
// 	// fmt.Printf("Character value:%c", val)
// 	// fmt.Printf("\nASCII value :%d", val)
// 	// var val float64=-19.25
// 	// fmt.Printf("Absolute value of %f is %f",val,math.Abs(val))//Abs value takes float input only
// 	// var num1 float64 = 10.34
// 	// var num2 float64 = 2.7
// 	// var large float64 = 0
// 	// large = math.Max(float64(num1), float64(num2)) //type conversion
// 	// fmt.Printf("Largest number is:%f", large)
// 	// large = math.Min(float64(num1), float64(num2)) //type conversion
// 	// fmt.Printf("\nSmallest number is:%f", large)
// 	// large = math.Pow(num1, num2) //type conversion
// 	// fmt.Printf("\nExponent number is:%g", large)
// 	// large = math.Sqrt(float64(num1)) //type conversion
// 	// fmt.Printf("\nSqure Root number is:%f", large)
// 	// large = math.Ceil(float64(num1)) //type conversion
// 	// fmt.Printf("\nCeil Root number is:%f", large)
// 	// large = math.Floor(float64(num1)) //type conversion
// 	// fmt.Printf("\nFoor Root number is:%f\n", large)

// 	// print("Using println instead of fmt.Println  ")
// 	// print("Using println instead of fmt.Println")

// 	//***************OPERATORS***********************
// 	// var a float64
// 	// var b string = "ketan"
// 	// var area float64
// 	// //var c float32 = 30.67
// 	// //var c int = 10
// 	// println("/nenter the number")
// 	// fmt.Scanf("%d", &a)
// 	// // print(a % 2)
// 	// // println(c)
// 	// // println(c & a)
// 	// // println(c | a)
// 	// // println(c ^ a)
// 	// // println(a > 5)
// 	// // println(a == 10)
// 	// //x>>=3

// 	// //sizeof is defined in 'unsafe' libraray
// 	// println(unsafe.Sizeof(a))
// 	// println(unsafe.Sizeof(b))
// 	// fmt.Println(reflect.TypeOf(a))
// 	// defer fmt.Println(reflect.TypeOf(b))
// 	// defer fmt.Println(reflect.ValueOf(a).Kind())
// 	// fmt.Println(reflect.ValueOf(b).Kind())
// 	// area = PI * a
// 	// println(area)

// 	// var ftemp float64
// 	// var ctemp float64
// 	// fmt.Print("Enter the temperatur in celcius:")
// 	// fmt.Scanf("%f", &ftemp)
// 	// ctemp = (ftemp * 1.8) + 32
// 	// fmt.Printf("The temperature in farenheit : %.2f", ctemp)
// 	// fmt.Print("hi its ketan")
// 	// var dd int = 20
// 	// var mm int = 02
// 	// var yy int = 2024
// 	// var str string
// 	// var flt float32 = 90.78
// 	// str = fmt.Sprintf("%02d-%02d-%04d", dd, mm, yy)
// 	// io.WriteString(os.Stdout, str)

// 	// fmt.Printf("\nenter the string:")
// 	// fmt.Scan(&str, &dd, &flt)
// 	// fmt.Printf("\nResult:%s,%d\n", str, dd)

// 	//Why type casting is important??
// 	// var a int = 123
// 	// var b uint = 0
// 	// b = uint(a) //type cast nahi kiya tho error
// 	// c := float64(-10.00)
// 	// fmt.Printf("\t%g", math.Abs(c))
// 	// fmt.Printf("\t%g", math.Sqrt(math.Abs(c)))
// 	// //assign value to b of a
// 	// fmt.Printf("\na=%d,b=%d\n", a, b)

// 	// var x int = 225
// 	// var r float32
// 	// r = float32(math.Sqrt(float64(x)))
// 	// print(r)

// 	// var x int = 42
// 	// var y float64 = float64(x)
// 	// var z uint = uint(y)
// 	// fmt.Printf("\nValue of x is %d and type is %T\n", x, x)
// 	// fmt.Printf("\nValue of y is %2f and type is %T\n", y, y)
// 	// fmt.Printf("\nvalue of z is %d and type is %T\n", z, z)
// 	// var num1 int
// 	// var num2 int
// 	// var avg float64
// 	// num1 = 10
// 	// num2 = 5
// 	// avg = (float64(num1) + float64(num2)) / 2
// 	// //avg = (num1+num2) / 2 will give error as int cant be stored in float
// 	// fmt.Printf("Average of %d and %d is %.2f\n", num1, num2, avg)

// 	//****Selection Iterations Sequencing ****
// 	/*if condition/expression {
// 		code to be executed
// 	}*/

// 	//a := 20
// 	// b := 10

// 	// if !(a > b) {
// 	// 	fmt.Println("Hi its ketan 1")
// 	// }
// 	// if 20 > 10 {
// 	// 	fmt.Println("Hi its ketan 2")
// 	// }
// 	// if true {
// 	// 	fmt.Println("Hi its ketan 3")
// 	// }

// 	// if b := 10; a > b { //b would be block variable ,only one ststement kar sakte a,b=10,20 works
// 	// 	fmt.Println("a is less than b")
// 	// }
// 	//var age int
// 	//var name string
// 	// print("enter the values")
// 	// if _, err := fmt.Scan(&name, &age); err != nil {
// 	// 	fmt.Println(err)
// 	// 	os.Exit(1)
// 	// }
// 	// println("Name:", name)
// 	// println("Age:", age)

// 	// print("enter the values")
// 	// if a, err := fmt.Scan(&name, &age); err == nil {//fmt.scan jo value return karta hai voh return karata hai hum determine nahi kar sakta hai
// 	// 	fmt.Println("Hey", a)
// 	// }
// 	// println("Name:", name)
// 	// println("Age:", age)

// 	// var a, b, c int
// 	// println("Enter the values of a,b,c")
// 	// fmt.Scanln(&a, &b, &c) //eak line mai hi inputs likho
// 	// // fmt.Scanln(&a)//harr statement enter k badd input fmt.Scanln(&a,&b) dono eak line mai likh sakte
// 	// // fmt.Scanln(&b)
// 	// // fmt.Scanln(&c)
// 	// if (a > b) && (a > c) {
// 	// 	println("The Largerst value is :", a)
// 	// }
// 	// if (b > a) && (b > c) {
// 	// 	println("The Largerst value is :", b)
// 	// }
// 	// if (c > b) && (c > a) {
// 	// 	println("The Largerst value is :", c)
// 	// }
// 	/*if condn {

// 	}else{ //else must be written in same line of ending } of if

// 	}*/

// 	// a := 20
// 	// b := 10

// 	// if a > b {
// 	// 	println("a is greater than b")
// 	// } else {
// 	// 	println("b is greater the a")
// 	// }

// 	// var a int
// 	// println("enter the number")
// 	// fmt.Scan(&a)
// 	// // if a%2 == 0 {
// 	// // 	println("It is even number")
// 	// // } else {
// 	// // 	println("it is odd number")
// 	// // }

// 	// if a < 0 {
// 	// 	a = -1 * a
// 	// 	fmt.Printf("Abs value of a is: %d", a)
// 	// } else {
// 	// 	println("the value of a is ", a)
// 	// }

// 	/*func functionname (argument list){

// 	}*/

// 	//22-02-2024//
// 	//******************Functions************************//

// 	/*func functionName(){

// 	}*/
// 	var x int
// 	var y int
// 	//var sum int
// 	print("/nEnter the numbaers:")
// 	fmt.Scan(&x, &y)
// 	// greet()
// 	// addition()
// 	// addition1(x, y)S
// 	// z := addition2(x, y)
// 	// print("\nreturned sum :", z)
// 	// sum, difference := calculate(x, y)
// 	// fmt.Printf("\nreturned sum is %d difference is %d:", sum, difference)
// 	// s := square(x)
// 	// fmt.Printf("\nreturned sum is :%d ", s)
// 	// sum2()
// 	// fmt.Printf("Sum is:", sum)
// 	//countdown(x)
// 	// fmt.Println(factorial(3))
// 	// fmt.Println(factorial(4))
// 	// fmt.Println(factorial(5))
// 	if x < 0 {
// 		print("enter valid input,Try again with vaild input >0")
// 	} else {
// 		fmt.Print("the sum is :", sum(x))
// 	}

// 	welcome := greet
// 	welcome()
// 	sum1(x, y)
// 	fmt.Printf("%T", sum1)

// 	var result = sum3(x, y)
// 	fmt.Println("resulted Sum is:", result)
// 	///Function having Function as argument
// 	expected := square(sum3(x, y))
// 	fmt.Println("resulted Sum is:", expected)
// 	expected1 := square(sub3(x, y))
// 	fmt.Println("resulted Sum is:", expected1)
// 	expected2 := square(mul3(x, y))
// 	fmt.Println("resulted Sum is:", expected2)
// 	expected3 := square(div3(x, y))
// 	fmt.Println("resulted Sum is:", expected3)

// }

// func greet1() {
// 	print("Hello World!! kasa kay :)")
// }

// func addition() {
// 	n1 := 10
// 	n2 := 20
// 	sum := n1 + n2
// 	print("\nsum is :", sum)
// }

// func addition1(n1 int, n2 int) {
// 	sum := n1 + n2
// 	print("\nsum is :", sum)

// }

// func addition2(n1 int, n2 int) int { //return type bhi likho
// 	sum := n1 + n2
// 	return (sum)
// 	//print("hi")

// }

// func calculate(n1 int, n2 int) (int, int) {
// 	sum := n1 + n2
// 	difference := n2 - n1
// 	return sum, difference
// }

// func square(n1 int) int {
// 	s := n1 * n1
// 	return s

// }

// func sum2() {
// 	var sum int
// 	sum = 5 + 9
// 	print(sum)
// }

// //******************** Recursion *********************

// /*func recurse(){
// 	---
// 	---
// 	recurse()
// }*/

// // func countdown(n1 int) {
// // 	fmt.Println(n1)
// // 	if n1 > 1 {
// // 		countdown(n1 - 1)
// // 	}

// // }

// // func factorial(num int) int {
// // 	if num == 1 || num == 0 {
// // 		return num
// // 	}
// // 	return num * factorial(num-1)
// // }

// func sum(number int) int {
// 	if number == 0 {
// 		return 0
// 	} else {
// 		return number + sum(number-1)
// 	}
// }

// /*func(){

// }*/

// var greet = func() {
// 	fmt.Println("\nHello,how are you")
// }
// var sum1 = func(n1, n2 int) {
// 	result := n1 + n2
// 	fmt.Println("Sum is:", result)
// }

// var sum3 = func(n1, n2 int) int {
// 	result := n1 + n2
// 	return result
// }
// var sub3 = func(n1, n2 int) int {
// 	result := n1 - n2
// 	return result
// }
// var mul3 = func(n1, n2 int) int {
// 	result := n1 * n2
// 	return result
// }
// var div3 = func(n1, n2 int) int {
// 	result := n1 / n2
// 	return result
// }

// //kpi bhi function function mai define nahi kar sakte ,but ananymous functions can be defined

// //********************************24/02/24*****************************

// package main

// import "fmt"

// func calculate() func() int {
// 	num := 1
// 	return func() int {
// 		num = num + 2
// 		return num
// 	}
// }

// func main() {
// 	odd := calculate()
// 	fmt.Println(odd())
// 	fmt.Println(odd())
// 	fmt.Println(odd())

// 	odd2 := calculate()
// 	fmt.Println(odd2())
// }

// package main

// import "fmt"

// func greet() func() string {
// 	name := "john"
// 	return func() string {
// 		name = "Hi " + name
// 		return name
// 	}
// }

// func main() {
// 	messages := greet()
// 	//ketan
// 	fmt.Println(messages())

// }
// package main

// import "fmt"

// func displayNumbers() func() int {
// 	number := 0
// 	return func() int {
// 		number++
// 		return number
// 	}
// }
// func main() {
// 	num1 := displayNumbers()
// 	fmt.Println(num1())
// 	num2 := displayNumbers()
// 	fmt.Println(num2())
// 	num3 := displayNumbers()
// 	fmt.Println(num3())

// package main

// import "fmt"

// func displayNumbers() func() int {
// 	number := 0
// 	return func() int {
// 		number++
// 		return number
// 	}
// }
// func main() {
// 	num1 := displayNumbers()
// 	fmt.Println(num1())
// 	num2 := displayNumbers()
// 	fmt.Println(num2())
// 	fmt.Println(displayNumbers()())
// 	num3 := displayNumbers()
// 	fmt.Println(displayNumbers())
// 	fmt.Println(num3())

// }

//*********febo using recursion and ananymous***********

// package main

// import "fmt"

// func main() {

// 	var fib func(n int) int
// 	fib = func(n int) int {
// 		if n < 2 {
// 			return n
// 		}
// 		return fib(n-1) + fib(n-2)
// 	}
// 	fmt.Println(fib(7))
// }

// ***********************************************02-03-24************************************************
// Nested if else block
package main

import (
	"time"
)

// func main() {
// 	var year int
// 	println("enter the value for year:")
// 	fmt.Scanln(&year)
// 	if (year%4 == 0 && year%100 != 0) || (year%4 == 0 && year%100 == 0 && year%400 == 0) {
// 		if year == 0 {
// 			println("invalid year")
// 		} else {
// 			println(year, " is a leap year")
// 		}
// 	} else {
// 		println(year, " is not a leap year")
// 	}
// }

// func main() {
// 	x := 8
// 	y := 10

// 	if x != y {
// 		if x < y {
// 			println(x, "is less than", y)
// 		} else if x > y {
// 			println(x, "is more than", y)
// 		} else {
// 			println(x, " is equal to ", y)
// 		}

// 	}
// }

// func main() {
// 	var time int
// 	println("Enter the time:")
// 	fmt.Scanln(&time)
// 	if 0 < time && time < 24 && time != 0 {
// 		if time < 10 {
// 			println("good moring")
// 		} else if time > 10 && time <= 20 {
// 			println("good day")
// 		} else {
// 			println("good evening")
// 		}

// 	}

// }

//**************************************************Switch Statements****************************************************

// func main() {
// 	M := time.Now()    //get system date/month/year
// 	Month := M.Month() //corp the Month from the date returned
// 	fmt.Printf("%T\n", M)
// 	switch Month {
// 	case 1:
// 		println("January")
// 	case 2:
// 		println("Febrary")
// 	case 3, 4:
// 		println("March", "April")
// 	case 5:
// 		println("May")
// 	}

// 	switch Month {
// 	case 1, 3, 5, 7, 8, 10:
// 		fmt.Println("31 days")
// 		fallthrough //goes to next case ,and direcly runs the next case without even verifing the nect case
// 	case 4, 6, 9, 11:
// 		fmt.Println("30 days")
// 	default:
// 		fmt.Println("this")

// 	}

// 	//var x interface{} = "RKNEC"
// 	var x interface{}
// 	switch i := x.(type) {
// 	case nil:
// 		fmt.Printf("type of x is %T", i)
// 	case int:
// 		fmt.Printf("type of x is int")
// 	case float32:
// 		fmt.Printf("type of x is float32")
// 	case bool, string:
// 		fmt.Print("Type of x is string or bool")
// 	default:
// 		fmt.Print("dont know the type")
// 	}

// }
//
//	func main() {
//		i := 0
//
// loop:
//
//		fmt.Println(i)
//		i++
//		if i < 5 {
//			goto loop
//		}
//		println("ends the loop")
//	}
// func main() {
// 	var str string = "Hello World"
// 	var substr string = "Wor"
// 	if strings.Contains(str, substr) { //checks the contains of substr in str
// 		fmt.Printf("String (%s) contains substring (%s)\n", str, substr)
// 	} else {
// 		fmt.Printf("String (%s) do not contain substring (%s)\n", str, substr)
// 	}
// 	println(strings.ToUpper(str)) //to upper
// 	println(strings.ToLower(str)) //to lower
// 	if (strings.Index(str, "W")) < 0 {
// 		println("does not exists")
// 	} else {
// 		println("exists st the position", (strings.Index(str, "W")))
// 	} //position of the "W" in str
// }

//*********************************03-03-2024************************************

//style one : assigning function to a variable

// func test(x int) {
// 	println("Hello", x)
// }
// func main() {
// 	x := test
// 	x(5)
// 	test(5)
// }

//style two : Anonnamous (must be defined inside a function only)
//inferred variable or deeclaration cant be declared gobally
// func main() {
// 	test := func(x int) {
// 		println(x)
// 	}
// 	test(5)
// }

// var test = func(x int) { //inferred like this is possiable but ':=' do not work
// 	println(x)
// }

// func main() {
// 	test(5)
// }

//style three : return type of anomnmous with diffent calling strategy

// func main() {
// 	test := func(x int) int {
// 		return x * x
// 	}(5) //default argument
// 	fmt.Println(test)
// }

// func calculate(x , y int) (int, int) {//last arg ka type mention kiya tho bhi it works
// 	return x + y, x - y
// }
// func main() {
// 	sum, difference := calculate(10, 20)
// 	println("sum is ", sum, "differnce is ", difference)
// }

//******************Date and Time Function **************************
// func main() {
// 	Y, M, D := time.Now().Date()
// 	println(D, "/", M, "/", Y)
// 	println("Date is :", D)
// 	println("Month is :", M)
// 	println("Year is :", Y)
//

// func main() {
// 	currentdatetime := time.Now()
// 	day := currentdatetime.Day()
// 	month := currentdatetime.Month()
// 	year := currentdatetime.Year()
// 	Hour := currentdatetime.Hour()
// 	Min := currentdatetime.Minute()
// 	Sec := currentdatetime.Second()

// 	fmt.Printf("The report is printed on %d/%d/%d at %d:%d:%d", day, month, year, Hour, Min, Sec)

// }

// func main() {
// 	println("HELLO!!")
// 	time.Sleep(4* time.Second)
// 	println("My Name is Ketan :)")
// 	time.Sleep(4* time.Second)
// 	println("I'am Budding Data Scientiest")
// }
