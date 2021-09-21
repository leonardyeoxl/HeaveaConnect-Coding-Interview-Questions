package main

import (
	"strconv"
)

func Max(x int, y int) int {
    if x < y {
        return y
    }
    return x
}

func reverse(s string) string {
    rns := []rune(s) // convert to rune
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
  
        // swap the letters of the string,
        // like first with last and so on.
        rns[i], rns[j] = rns[j], rns[i]
    }
  
    // return the reversed string.
    return string(rns)
}

func AddStrings(num1 string, num2 string) string {
	result := ""
	max := Max(len(num1), len(num2))
	diff := -1
	new_num1 := ""
	new_num2 := ""

	if len(num1) < max {
		diff = max - len(num1)	
		for i:=0; i<diff; i++ {
			new_num1 += "0"
		}
		new_num1 += num1
		num1 = new_num1
	}
	if len(num2) < max {
		diff = max - len(num2)	
		for i:=0; i<diff; i++ {
			new_num2 += "0"
		}
		new_num2 += num2
		num2 = new_num2
	}

	num1 = reverse(num1)
	num2 = reverse(num2)

	bring_forward := 0
	for index := 0; index < max; index++ {
		int1, err1 := strconv.Atoi(num1[index:index+1])
		int2, err2 := strconv.Atoi(num2[index:index+1])
		if err1 == nil && err2 == nil {
			sum := int1 + int2 + bring_forward
			if sum > 9 {
				result += strconv.Itoa(sum%10)
				bring_forward = int(sum/10)
			} else {
				result += strconv.Itoa(sum)
			}
		}
	}

	return reverse(result)
}

func main() {
	num1 := "11"
	num2 := "123"
	if AddStrings(num1, num2) != "134" {
		panic("Wrong answer for num1='11', num2='123'")
	}

	num1 = "456"
	num2 = "77"
	if AddStrings(num1, num2) != "533" {
		panic("Wrong answer for num1='456', num2='77'")
	}

	num1 = "0"
	num2 = "0"
	if AddStrings(num1, num2) != "0" {
		panic("Wrong answer for num1='0', num2='0'")
	}
}