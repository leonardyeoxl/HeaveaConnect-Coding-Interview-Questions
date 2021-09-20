# HeveaConnect-Coding-Interview-Questions

## Question 01 - Implement strStr()

- Implement `strStr()`.
- Return the index of the first occurrence of `needle` in `haystack`, or `-1` if `needle` is not part of `haystack`.

### Example 1:
```
Input: haystack = "hello", needle = "ll"
Output: 2
```
### Example 2:
```
Input: haystack = "aaaaa", needle = "bba"
Output: -1
```
### Example 3:
```
Input: haystack = "", needle = ""
Output: 0
```

### Constraints:
- `0 <= haystack.length`, `needle.length <= 5 * 10^4`
- `haystack` and `needle` consist of only lower-case English characters.

---

## Question 02 - Add Strings

- Given two non-negative integers, `num1` and `num2` represented as string, return the sum of `num1` and `num2` as a string.
- You must solve the problem without using any built-in library for handling large integers (such as BigInteger). You must also not convert the inputs to integers directly.

### Example 1:
```
Input: num1 = "11", num2 = "123"
Output: "134"
```
### Example 2:
```
Input: num1 = "456", num2 = "77"
Output: "533"
```
### Example 3:
```
Input: num1 = "0", num2 = "0"
Output: "0"
```

### Constraints:

- `1 <= num1.length`, `num2.length <= 10^4`
- `num1` and `num2` consist of only digits.
- `num1` and `num2` don't have any leading zeros except for the zero itself.

## Reference
- https://leetcode.com/problems/implement-strstr/
- https://leetcode.com/problems/add-strings/