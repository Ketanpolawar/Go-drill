// addition of two numbers
// var a int =10 ; static varable declaration
package main

import (
	"fmt"
	"io"
	"os"
)

const (
	C  = "ketan"
	PI = 3.14
)

//f:=10  global variable cannot be declared using':='

func main() {

	// n1 := 10 //inferred data type
	// n2 := 20
	// var sum, sum1, sum2, sum3, sum4 int
	// //Addition
	// sum = n1 + n2
	// fmt.Println("sum is", sum)
	// fmt.Println("sum n1 n2 is ", sum, n1, n2)
	// //Substraction
	// sum1 = n2 - n1
	// fmt.Println("sum is", sum1)
	// fmt.Println("sum n1 n2 is ", sum1, n1, n2)
	// //Multiplication
	// sum2 = n2 * n1
	// fmt.Println("sum is", sum2)
	// fmt.Println("sum n1 n2 is ", sum2, n1, n2)
	// //Division
	// sum3 = n2 / n1
	// fmt.Println("sum is", sum3)
	// fmt.Println("sum n1 n2 is ", sum3, n1, n2)
	// //Modulus
	// sum4 = n2 % n1
	// fmt.Println("sum is", sum4)
	// fmt.Println("sum n1 n2 is ", sum4, n1, n2)

	//day 2

	// const PI = 3.34
	// //PI = PI + 1  //will throw the error as changin the value of const "not addressable error/or apping error"
	// fmt.Println("PI is ", PI)

	//Typed and Untyped constant
	//Typed which is statically declared via type name
	//Untyped is inferred using '='
	// const(
	// 	A = 10
	// 	B = 20
	// )"Declaring multiple constants at simultaneously"

	// fmt.Println("Const ", C)

	// //Print f function

	// const L = 10
	// const W = 5
	// var area int
	// area = L * W
	// fmt.Printf("Area of the triangle:%d", area) //p in small will give undefined
	// // var s float32 = 123456789012
	// // s = s + 9
	// // fmt.Printf("\nArea of the triangle:%f", s)
	// //go will throw error in c ase of float overflow ,while in C goes on -ve axis like 255+3 will result in -3

	// var a uint8 // Analogy if need 1 lit of water we do not use the 10 lit container as it is inefficient same as byte  rune as int 32
	// a = 30      //will show error of overflowif above 255 in go (only +ve)
	// fmt.Printf("\nValue of Unsigned int:%d", a)
	// //uintptr unsigned ptr varaible

	// var student1 string = "John"
	// var student2 = "Jane"
	// x := 10
	// fmt.Printf("\nValue of variables :%s and %s and %T\n", student1, student2, x) // F in Printf stands for format specifier
	// fmt.Println("Value of variables :", student1)
	// fmt.Println("Value of variables :", student2) //default prints in next line '\n' not supported

	// fmt.Print("Value of variables  :", x) //print the text and value but not in new line even if two print statement are used

	// f := "foo"
	// fmt.Printf("\n%T", f)

	// p, y, z, k := 0, 2.2, "ketan", true

	// fmt.Println("the values variables are ", p, y, z, k)

	// var s, e = 6, "Hello"
	// c, d := 7, "world"
	// fmt.Println("the values variables are ", s, e, c, d)

	// var (
	// 	o     = "polawar"
	// 	v int = 1
	// )
	// fmt.Println("the values variables are ", o, v)
	// fmt.Printf("the vales variable are %s ,%d\n", o, v)
	//var fname, lname, Age, Branch, College string = "Ketan", "Polawar", "20", "Data Science", "RCOEM"
	// fmt.Printf("Hey enter your fname\n")
	// fmt.Scanf("%s", &fname)
	// fmt.Printf("Hey enter your lname\n")
	// //fflush use karan padega else woh 2nd input nahi lega
	// fmt.Scanln(&lname)

	// fmt.Printf("name is %s,%s ", fname, lname)
	// var i, b int = 10, 20
	// fmt.Println(i, b)
	// fmt.Println(i + b)
	// fmt.Println(fname + lname)

	// fmt.Println("Name :", fname, lname)
	// fmt.Println("Age:", Age)
	// fmt.Println("Branch:", Branch)
	// fmt.Println("CollegeName:", College)

	// fmt.Printf("\nName : '%s %s'\tAge:'%s'\n", fname, lname, Age)
	// fmt.Printf("Branch:'%s'\tCollege:'%s'\n", Branch, College)

	// const name, dept = "DASCA", "DS"
	// fmt.Printf("%s is a %s Portal", name, dept)

	// var str = "Ketan_Polawr"
	// fmt.Printf("The string is %s \n", str) //string
	// var num1 int = 20
	// fmt.Printf("The string is %d \n", num1) //int
	// fmt.Printf("The string is %b \n", num1) //prints binary value
	// var num3 float32 = 1.22
	// fmt.Printf("The string is %g \n", num3) //float num %f can also be employeed
	// var num2 = 4 + 1i
	// fmt.Printf("The string is %e \n", num2) //complex number

	// fmt.Printf("The string is %v,%v,%v,%v \n", str, num3, num2, num1) //it uses the original type and prints the values
	//C prints the garbage values in case of undefined values  but In Go prints the defined value
	//default values
	//int 0
	//float 0.0
	//string nil or blank space
	//Boolean False
	//Format specifiers in go
	//%d integer
	//%f float
	//%T type of varable

	//var can be used inside ans outside the variable , also assignemet and declaration can be done seperately
	//':=' cannot be used to declare the global declaration ,also assignemet and declaration can be done simultaneously
	//literals (something having fixed value)are always rvalue x=10 :: x is lvalue and 10 is rvalue
	//integer literals has decimal(10 starts without the 0) octal(8 starts 0,0o,0O) hexa(16  starts with 0x or 0X) binary(2 starts with 0b or 0B)
	// fmt.Println(15 == 017)
	// fmt.Println(15 == 0xF)
	// println(13 == 0xD)
	// var a int = 10
	// a++
	// fmt.Println(a)
	// fmt.Printf("hex of %d is %0x\n", a, a)
	// fmt.Printf("hex of %d is %0b\n", a, a)
	// fmt.Printf("hex of %d is %0o\n", a, a)
	// pre increment ie --a,++a is not defined in GO post increment is defined  C langauge lang supports the both
	//e stands as an exponent
	// a = 1.23e5
	// fmt.Print(a)
	//%T is type %t is for boolean
	// fmt.Printf("\nExpression 15==0xf is '%t' and has type '%T'", 15 == 0xf, 15 == 0xf)
	//\a,\t,\n are called escape specifier
	// fname := "Ketan"
	// lname := "Polawar"
	// fmt.Print(fname, "\n")
	// fmt.Print(lname, "\n")
	// fmt.Print("my name is ", fname, " and last name is ", lname)
	// fmt.Printf("\n%s\n", fname)
	// fmt.Printf("%s\n", lname)
	// fmt.Printf("%s  %s\n", fname, lname)
	// //print by defaults inserts the space between the literals if they are not string type else either any one is string  explictly mention space-
	// fmt.Print(fname, 20)
	//var i = 15
	//var txt = "ketan"
	// fmt.Printf("%v\n", i)  //will give Ketan
	// fmt.Printf("%#v\n", i) //will give "ketan"
	// fmt.Printf("%v%%\n", i)
	// fmt.Printf("%T\n", i)
	// fmt.Printf("\"ketan\"")
	// fmt.Printf("'Ketan'")
	// fmt.Printf("%b\n", i) //binary
	// fmt.Printf("%d\n", i) //prints the sign also
	// fmt.Printf("%+d\n", i)
	// fmt.Printf("%o\n", i)
	// fmt.Printf("%O\n", i)
	// fmt.Printf("%x\n", i)
	// fmt.Printf("%X\n", i)
	// fmt.Printf("%#x\n", i)
	// fmt.Printf("%4d\n", i)  //writes in 4 digit
	// fmt.Printf("%-4d\n", i) //left justify
	// fmt.Printf("%04d\n", i)
	//output
	//1111
	// 15
	// +15
	// 17
	// 0o17
	// f
	// F
	// 0xf
	//   15
	// 15
	// 0015
	// fmt.Printf("%s\n", txt)
	// fmt.Printf("%q\n", txt)
	// fmt.Printf("%8s\n", txt)
	// fmt.Printf("%-8s\n", txt)
	// fmt.Printf("%x\n", txt)
	// fmt.Printf("% x\n", txt)

	// var (
	// 	marks1     int
	// 	marks2     int
	// 	marks3     int
	// 	roll       int
	// 	percentage int
	// 	sum        int
	// 	name       string
	// 	branch     string
	// )

	// fmt.Printf("\n**Please Enter The Details**\n")
	// fmt.Println("**Please Enter Your Name**")
	// fmt.Scanf("%s", &name)
	// fmt.Println("**Please Enter Your Branch**")
	// fmt.Scanf("%s", &branch)
	// fmt.Println("**Please Enter Your Id**")
	// fmt.Scanf("%s", &roll)
	// fmt.Println("**Please Enter Your Marks1**")
	// fmt.Scanf("%s", &marks1)
	// fmt.Println("**Please Enter Your Marks2**")
	// fmt.Scanf("%s", &marks2)
	// fmt.Println("**Please Enter Your Marks3**")
	// fmt.Scanf("%s", &marks3)

	// sum = marks1 + marks2 + marks3
	// percentage = (sum / 300.0) * 100

	// println("SHRI RAMDEOBABA COLLEGE OF ENGINEERING AND MANAGEMENT")
	// fmt.Printf("\nName : %s\tId:'%d'\n", name, roll)
	// fmt.Printf("Branch:'%s'\tsum:'%d'\n", branch, sum)
	// fmt.Printf("Percentage:'%d'", percentage)

	// integer := 23
	// fmt.Printf("%T %T\n", integer, &integer)
	// fmt.Print(integer, &integer)

	// number := 1234.575883939
	// fmt.Printf("number:%5.2f", number)

	// fmt.Printf("%*s\n", 40, "text")//width of output is 40  can be specified "%040s"

	// ch := 'A'
	// bch := "ketan"
	// fmt.Printf("%T %T\n", ch, bch)

	// var val byte = 'A' //character define kar sakte hai with the help of
	// fmt.Printf("Character value:%c", val)
	// fmt.Printf("\nASCII value :%d", val)
	// var val float64=-19.25
	// fmt.Printf("Absolute value of %f is %f",val,math.Abs(val))//Abs value takes float input only
	// var num1 float64 = 10.34
	// var num2 float64 = 2.7
	// var large float64 = 0
	// large = math.Max(float64(num1), float64(num2)) //type conversion
	// fmt.Printf("Largest number is:%f", large)
	// large = math.Min(float64(num1), float64(num2)) //type conversion
	// fmt.Printf("\nSmallest number is:%f", large)
	// large = math.Pow(num1, num2) //type conversion
	// fmt.Printf("\nExponent number is:%g", large)
	// large = math.Sqrt(float64(num1)) //type conversion
	// fmt.Printf("\nSqure Root number is:%f", large)
	// large = math.Ceil(float64(num1)) //type conversion
	// fmt.Printf("\nCeil Root number is:%f", large)
	// large = math.Floor(float64(num1)) //type conversion
	// fmt.Printf("\nFoor Root number is:%f\n", large)

	// print("Using println instead of fmt.Println  ")
	// print("Using println instead of fmt.Println")

	//***************OPERATORS***********************
	// var a float64
	// var b string = "ketan"
	// var area float64
	// //var c float32 = 30.67
	// //var c int = 10
	// println("/nenter the number")
	// fmt.Scanf("%d", &a)
	// // print(a % 2)
	// // println(c)
	// // println(c & a)
	// // println(c | a)
	// // println(c ^ a)
	// // println(a > 5)
	// // println(a == 10)
	// //x>>=3

	// //sizeof is defined in 'unsafe' libraray
	// println(unsafe.Sizeof(a))
	// println(unsafe.Sizeof(b))
	// fmt.Println(reflect.TypeOf(a))
	// defer fmt.Println(reflect.TypeOf(b))
	// defer fmt.Println(reflect.ValueOf(a).Kind())
	// fmt.Println(reflect.ValueOf(b).Kind())
	// area = PI * a
	// println(area)

	// var ftemp float64
	// var ctemp float64
	// fmt.Print("Enter the temperatur in celcius:")
	// fmt.Scanf("%f", &ftemp)
	// ctemp = (ftemp * 1.8) + 32
	// fmt.Printf("The temperature in farenheit : %.2f", ctemp)
	// fmt.Print("hi its ketan")
	var dd int = 20
	var mm int = 02
	var yy int = 2024
	var str string
	str = fmt.Sprintf("%02d-%02d-%04d", dd, mm, yy)
	io.WriteString(os.Stdout, str)

	fmt.Printf("\nenter the string:")
	fmt.Scan(&str)
	fmt.Printf("\nResult:%s\n", str)

}
