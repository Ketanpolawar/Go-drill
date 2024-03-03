func main() {
	var colors = []string{"Red", "Blue", "Green"}//slice wthout the size
	for index, val := range colors { //range always returns 2 values index and value as showm
		fmt.Println(index, ":-", val)
	}

	for _, val := range colors { //range always returns 2 values index and value as showm
		fmt.Println(val)
	}
}