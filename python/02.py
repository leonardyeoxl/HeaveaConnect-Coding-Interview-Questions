# Time Complexity: O(max(N,M)). Given N is the number of elements in num1 and M is the number of elements in num2
# Space Complexity: O(1)
def AddStrings(num1, num2):
    _max = max(len(num1), len(num2))
    
    if len(num1) < _max:
        diff = _max-len(num1)
        new_num1 = "0"*diff + num1
        num1 = new_num1
    if len(num2) < _max:
        diff = _max-len(num2)
        new_num2 = "0"*diff + num2
        num2 = new_num2

    num1 = num1[::-1]
    num2 = num2[::-1]

    bring_forward = 0
    result = ""
    for index in range(_max): # O(max(N, M))
        _sum = int(num1[index]) + int(num2[index]) + bring_forward
        if _sum > 9:
            result += str(_sum%10)
            bring_forward = _sum//10
        else:
            result += str(_sum)

    return result[::-1]


num1 = "11"
num2 = "123"
AddStrings(num1, num2)
assert AddStrings(num1, num2) == "134", "Wrong answer for num1='11', num2='123'"

num1 = "456"
num2 = "77"
AddStrings(num1, num2)
assert AddStrings(num1, num2) == "533", "Wrong answer for num1='456', num2='77'"

num1 = "0"
num2 = "0"
AddStrings(num1, num2)
assert AddStrings(num1, num2) == "0", "Wrong answer for num1='0', num2='0'"