import sys
import math

s = raw_input()
string=""
for character in s:
    if(character.lower()!=character.upper()):
        if(character==character.upper()):
            string+=character

print string
