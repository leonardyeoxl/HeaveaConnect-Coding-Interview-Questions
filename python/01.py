# Time Complexity: O(H*N). Given H is the number of elements in the haystack and N is the number of elements in the needle
# Space Complexity: O(1).
def strStr(haystack, needle):
    if len(haystack) == 0 or len(needle) == 0:
        return 0
    
    for index in range(len(haystack)):
        right, count = index, 0
        while count < len(needle)+1:
            if haystack[index:right+1] == needle:
                return index
            right += 1
            count += 1
    
    return -1

haystack = "hello"
needle = "ll"
assert strStr(haystack, needle) == 2, "Wrong haystack='hello', needle='ll'"

haystack = "aaaaa"
needle = "bba"
assert strStr(haystack, needle) == -1, "Wrong haystack='aaaaa', needle='bba'"

haystack = ""
needle = ""
assert strStr(haystack, needle) == 0, "Wrong haystack='', needle=''"