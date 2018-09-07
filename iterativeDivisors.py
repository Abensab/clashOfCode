import sys
import math

# Auto-generated code below aims at helping you parse
# the standard input according to the problem statement.

n = float(raw_input())
l = int(raw_input())

# Write an action using print
# To debug: print >> sys.stderr, "Debug messages..."

print >> sys.stderr, str(n), str(l)
i=0
res = []
string ="["

while i<l:
    if n//1>0:
        res.append(int(n//1))
        n=n-n//1
        i+=1
        print >> sys.stderr, str(int(n))
    else:
        n=1/n
for x in range len(res)-1:
    string+=str(res[x])+','
    
string+=res[len(res)-1]+']'

print string
