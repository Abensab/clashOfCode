import sys
s = raw_input()
print (str(any(char.istitle() for char in s) and any(char.isdigit() for char in s) and len(s)>7 and any(char.islower() for char in s)).lower())
