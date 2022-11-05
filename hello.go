package main

import (
	"fmt"

	"math"

	"rsc.io/quote/v4"
)

func describeTheStringSlice(s []string) {
	fmt.Println("length:", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Println("*", s[i])
	}
}

func koleczko() {
	i := 2

	for i < 20 {
		fmt.Println(math.Pow(float64(i), 2.0))
		i++
	}
}

func divisibleByThirteen() {
	for i := 0; i < 100; i++ {
		if i%13 == 0 {
			fmt.Println("Dvisible by 13: ", i)
		} else if i%17 == 0 {
			fmt.Println("Dvisible by 17: ", i)
		} else {
			fmt.Println(i, "not divisible by 13 nor 17 :(")
		}
	}
}

func matma() {
	const a = 20e5
	const b = 0.713
	const c = a / b

	fmt.Println("c:", c, "sinus:", math.Sin(b))
}

func summy() {
	var a, b int = 13, 45
	fmt.Println(a + b)
	a = 19
	fmt.Println(a + b)
}

func doit() {
	var a = "Adrian"
	b := "Gonciarz"
	fmt.Println(a, b)
}

func tempSwitch(temp int) {
	switch {
	case temp < 0:
		fmt.Println("zimno kurrrrr....")
	case temp < 25:
		fmt.Println("elegancko")
	case temp >= 25:
		fmt.Println("caliente")
	}
}

func simpleSwitch(value float32) {
	switch value {
	case 0.0:
		fmt.Println("pustka")
	case 1.0, 3.0:
		fmt.Println("dobro")
	default:
		fmt.Println("no tak nie wiadomo do końca")
	}
}

func arrayFun() {
	var a [5]int
	fmt.Println("empty:", a)

	a[0] = 5
	fmt.Println("curent:", a)
	fmt.Println("Len of array:", len(a))

	b := [5]int{0, 11, 33, 49, 39}
	fmt.Println(b)

	var twoDimArray [3][3]int

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			twoDimArray[i-1][j-1] = i * j
		}
	}
	fmt.Println(twoDimArray)

}

func slicesFun() {
	students := make([]string, 5)
	students[0] = "adam@example.com"
	students[1] = "katarzyna@example.com"
	students[2] = "mikhail@example.com"

	fmt.Println(students)
	fmt.Println(len(students))

	students = append(students, "laura@example.com")
	describeTheStringSlice(students)

	fmt.Println("First two students", students[:2])
	fmt.Println("Middle students", students[3:6])
}

func arrayReferences() {
	var ar1 [4]int
	ar2 := ar1
	ar1[0] = 42
	ar2[0] = 24

	fmt.Println(ar1, ar2)
	dynArr := [...]string{"cos", "coss", "asd", "Asd2134"}
	fmt.Println(len(dynArr))
}

func slicing() {
	ar1 := [4]string{"f", "oo", "ba", "r"}
	sl1 := ar1[1:3]
	sl1 = append(sl1, "du", "pa")
	fmt.Println(sl1)

	var emptySlice []string
	fmt.Println(len(emptySlice))
	emptySlice = append(emptySlice, "d", "r", "b")
	fmt.Println(emptySlice)
}

func mapFun() {
	scores := make(map[string]int)
	scores["adi"] = 12
	scores["pat"] = 8
	scores["sam"] = 32
	scores["mar"] = 52

	delete(scores, "sam")

	fmt.Println(scores)

	users := map[string]bool{"a": true, "b": false, "c": true}
	fmt.Println(users)
}

func rangeTest() {
	scores := []float32{3.4, 3.3, 7.8, 9.0, 1.9, 4.2}
	for idx, score := range scores {
		fmt.Printf("Score %f at index %d\n", score, idx)
	}

	users := map[string]bool{"a": true, "b": false, "c": true}
	for k, v := range users {
		fmt.Printf("\tMap key %s is %t\n", k, v)
	}

}

func main() {
	fmt.Println(quote.Go())
	fmt.Println("Foo" + "Bar")
	fmt.Println(1 + 2)
	fmt.Println(true && false)
	fmt.Println(7.0 / 3)

	doit()
	summy()
	matma()
	koleczko()
	divisibleByThirteen()
	tempSwitch(-5)
	tempSwitch(0)
	tempSwitch(15)
	tempSwitch(25)
	tempSwitch(26)
	simpleSwitch(1.0)
	simpleSwitch(2.0)
	simpleSwitch(3.0)

	arrayFun()
	slicesFun()
	arrayReferences()
	slicing()
	mapFun()
	rangeTest()
}
