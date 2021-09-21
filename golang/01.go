package main

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 || len(needle) == 0 {
		return 0
	}

	for index:=0; index < len(haystack); index++ {
		right := index
		count := 1
		for {
			if count < len(needle)+1 {
				if right+1 <= len(haystack){
					if haystack[index:right+1] == needle {
						return index
					}
				}
			} else {
				break
			}
			right += 1
			count += 1
		}
	}

	return -1
}

func main() {
	haystack := "hello"
	needle := "ll"
	if strStr(haystack, needle) != 2 {
		panic("Wrong. haystack='hello', needle='ll'")
	}

	haystack = "aaaaa"
	needle = "bba"
	if strStr(haystack, needle) != -1 {
		panic("Wrong haystack='aaaaa', needle='bba'")
	}

	haystack = ""
	needle = ""
	if strStr(haystack, needle) != 0 {
		panic("Wrong haystack='', needle=''")
	}
}