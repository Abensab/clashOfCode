import sys
import math

#My code:
def poor_code():
    s = raw_input()
    string=""
    for character in s:
        if(character.lower()!=character):
            string+=character

    print(string)

#Better code (RedDead):
def better_code():
    s = input()
    b = [el for el in s if el.lower() != el] 
    print("".join(b))
    
#Bettest code (Me, after seeing the code of RedDead):
def bettest_code():
    print("".join([c for c in input() if c.lower() != c]))
