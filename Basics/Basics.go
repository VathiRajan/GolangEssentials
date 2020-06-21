//Hello
package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println("---------Declare without variable type------")
	x, y := 101.1, 200.25

	fmt.Printf("x=%v, type of %T\n, y=%v\n", x, x, y)
	fmt.Printf("y=%v, type of %T\n", y, y)

	mean := (x + y) / 2
	fmt.Printf("result of mean = %v, of type %T\n", mean, mean)
	fmt.Println("---------If Else ------")

	if mean > 5 {
		fmt.Println("Yes bigger")
	} else {
		fmt.Println("less")
	}

	if mean > 100 && mean < 200 {
		fmt.Println("between 100-200")
	}
	fmt.Println("---------DSwitch case. You dont need to give break at the end of every case------")

	switch mean {
	case 150.675:
		fmt.Println("One Fifty")
	default:
		fmt.Println("printing defaut")
	}
	fmt.Println("---------Usual forloop-----")

	for i := 0; i < 3; i++ {
		fmt.Println("For loop")
	}
	fmt.Println("---------Function in GO------")

	fmt.Println(hello("hello", "world"))

	m, n := twoWay("hello", "world")
	fmt.Println(m, n)
	fmt.Println("---------Declare multiple variabes during initialisation------")

	var k, l = 1, "hello"
	f, n := 78, "gh"
	fmt.Println(k, l, f, n)
	fmt.Println("---------For loop with only ne condition like an while loop-----")

	i := 0
	for i < 3 {
		i++
		fmt.Println("i less than 3")
	}

	book := "Hello this is vathi"
	fmt.Println("---------determine length of a string-----")

	fmt.Println(len(book))
	fmt.Println(book[0])
	//book[0] = 116 //Strings are immutable in Go
	//Slicing a string
	fmt.Println(book[4:11])
	fmt.Println(book[5:])
	fmt.Println(book[:]) //print entire string

	fmt.Println("-----")
	//	To decimal string:

	no := int64(32)
	str := strconv.FormatInt(no, 10)
	fmt.Println(str) // Prints "32"
	//To hexadecimal string:

	ni := int64(32)
	stri := strconv.FormatInt(ni, 16)
	fmt.Println(stri) // Prints "20"

	//Multiline
	mul := `Hello
	This 
	is
	Vathi`
	fmt.Println(mul)

	numb := 41
	sh := fmt.Sprintf("%d", numb)
	fmt.Printf("y=%q, type of %T\n", sh, sh)
	//Time module
	fmt.Println(time.Now().Weekday())

	numset := []int{16, 8, 45, 12, 45}
	max := numset[0]
	//_ denotes the index. For each loop in Go (index,value)
	for _, value := range numset[1:] {
		if value > max {
			max = value
		}
	}
	fmt.Println(max)

	//Maps in Golang
	//two maps cannot be compared. Maps can only be compared with null
	newMap := map[string]float64{
		"one":   1.0,
		"two":   2.0,
		"three": 3.0,
	}
	//Check number of items in map
	fmt.Println("No of elements", len(newMap)) //it will retur 3

	//Get a value from Map
	fmt.Println(newMap["one"])

	//Check if the key is present or not
	fmt.Println(newMap["four"]) //it will return 0 bu default. Hence we have to check it.

	value, ok := newMap["four"]
	if !ok {
		fmt.Println(value, " four is not present")
	} else {
		fmt.Println("value present")
	}

	for key, value := range newMap {
		fmt.Println(key, value)
	}
	//set a value
	newMap["Four"] = 4.0
	delete(newMap, "one")
	newMap["Four"] = 44.0
	fmt.Println(newMap)

	// You can also create an empty slice and define the data type,
	// length (receive value of 0), capacity (max size)
	numSlice3 := make([]int, 5, 10)
	numSlice := []int{1, 2, 5}
	numSlice3 = append(numSlice3, 50, -1)
	copy(numSlice3, numSlice)
	fmt.Println("Created array", numSlice3)

	//Problem - Count the number of repeating words.

	//Use strings.fields to split the words.
	//use strings.ToLower to conver to ower case

	txt := `need the need
	human race is always a race
	but need still remains the race`

	words := strings.Fields(txt)
	fmt.Println(words)
	counts := map[string]int{} //declare an empty map
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}
	printGreen(counts)
	fmt.Println(counts)
	//Initilase the map
	//lookup := make(map[string]int, 200)

	//When you pass the function or arry you pass in the reference.
	//If you also wated to pass in the reference for variable you can use pointers
	pon := 10
	passingIn(&pon)
	fmt.Println(pon)
	//Exception Handling
	fmt.Println(sqroot(-1))
	fmt.Println("----")
	s1, err := sqrt(-1)
	if err != nil {
		fmt.Printf("Error : %s\n", err)
	} else {
		fmt.Println("Sqrt of number is", s1)
	}

	//HTT GET request
	cType, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Content type is %s", cType)
	}
	t := Person{"Vathi", 25, "M"}
	fmt.Println(t)
	t2 := Person{
		Name: "Vathirajan",
		Age:  26,
		Sex:  "Male",
	}
	fmt.Println(t2)
	trade := Trade{
		Stock:  "HDFC",
		Price:  99.98,
		Volume: 100,
		Buy:    true,
	}
	fmt.Println(trade.value())
	fmt.Println("--New Trade --")
	tr, err := NewTrade("ICCI", 95.35, 45, true)
	if err != nil {
		fmt.Printf("The trade %s", err)
	} else {
		fmt.Println(tr.value())
	}

	go say()

}

func say() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func (t *Trade) value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}
	return value
}

//NewTrade is
func NewTrade(stock string, price float64, volume float64, buy bool) (*Trade, error) {

	if stock == "" {
		return nil, fmt.Errorf("The stock is empty")
	}
	if volume <= 0.0 {
		return nil, fmt.Errorf("The volume is empty")
	}
	trade := &Trade{
		Stock:  stock,
		Price:  price,
		Volume: volume,
		Buy:    buy,
	}
	return trade, nil
}

//Trade is
type Trade struct {
	Stock  string  //Stock
	Price  float64 //Price
	Volume float64 //Volume
	Buy    bool    //Buy
}

//Person is
type Person struct {
	Name string //Name
	Age  int16  //Age
	Sex  string //Sex
}

func contentType(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	cType := res.Header.Get("Content-Type")
	if cType == "" {
		return "", fmt.Errorf("empty content Type")

	}
	return cType, nil

}

//	fmt.Println("---------First bracket is arguments and second bracket is Type of return------")
//func subtract(args ...) Var argument in Go
func hello(x, y string) string {
	return x + " " + y
}

//	fmt.Println("---------You can return two value from function------")

func twoWay(x, y string) (string, string) {
	return x + "Vathi", y
}

func printGreen(coun map[string]int) {
	// Handle map argument.
	coun["Mangatha"] = 5
}
func passingIn(pon *int) {
	*pon = 2
}
func sqrt(n float64) (float64, error) {

	if n < 0 {
		return 0.0, fmt.Errorf("Error in Sqrt (%f)", n)
	}
	return math.Sqrt(n), nil
}
func sqroot(n float64) float64 {

	return math.Sqrt(n)
}
