# ClashOfCode

The objective at this challenges is to be the fastest in pass all test cases and the final case, between a group with a maximun of 10 members.

I emphasize that my code couldn't be efficient or the best solution, but i've tried to done that.


## 1. Iterative Divisors challenge
In this challenge, we must get the 'ai' divisors at number 'n' like this example:


```
                  1
n = a0 + ----------------
                              1
                  a1 = ----------------
                                           1   
                              a2 = ----------------
                                           a3
```

The objective must have a list with 'ai' numbers.
π or e numbers have infinite divisors, this it why give us the lenght of the result list size.

### Example:

The result for π number with a length result of 5 must be like this (**without blank spaces**):
<pre>
[<b>3</b>,<b>7</b>,<b>15</b>,<b>1</b>,<b>292</b>]
</pre>

because the formula is:

<pre>
π = 3.141592653589793<br>
π = <b>3</b> + 0.141592653589793<br>
π = <b>3</b> + (1/0.141592653589793)<br>
π = <b>3</b> + (1/(<b>7</b>+0.062513306))<br>
π = <b>3</b> + (1/(<b>7</b>+(1/(<b>15</b>+0.996594389)))<br>
π = ...
</pre>
______________________________________________
## 2.Checking passwords
| Position | TIME  |  LANGUAGE  |
| :------: | :---: | :--------: |  
|    4     | 10:40 |   Python   |

The main objective of this challenge was chek if an input string is a secure password, for that it was accomplish 4 principles:

1. May be have almost 8 characters
2. One character mus be uppercase
3. One character must be lowercase
4. One character must be a number

If the string accomplish these characteristics, we mus return "true", else "false".


______________________________________________
## 3.Counting ones

| Position |  TIME |  LANGUAGE  | 
| :------: | :---: |  :-------: | 
|   7/8    | 15:00 |     Go     |

The game mode is REVERSE: You do not have access to the statement. You have to guess what to do by observing the following set of tests:

|    NAME    |                            INPUT                                | RETURN |
|  :------:  | :------------------------------------------------------------:  | :----: |
|01 Test 1   |                         3 <br>1 0 1                             |    0   |
|02 Test 2   |                         3 <br>0 1 0                             |    1   |
|03 Test 3   |                        5<br>1 0 1 0 0                           |    0   |
|04 Test 4   |                  10<br>1 0 1 0 0 1 0 1 0 0                      |    0   |
|05 Test 5   |                12<br>1 1 1 1 1 1 1 1 1 1 1 1                    |    0   |
|06 Test 6   |               13<br>1 1 1 1 1 1 1 1 1 1 1 1 1                   |    1   |
|07 Test 7   |     24<br>0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1       |    1   |
|08 Test 8   |30<br>0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 1 0 0 0 0 0|    0   |

The objective of this challenge was return "0" if the counting of 1 was even, 1 if was odd.
______________________________________________
## 4. Ascii to string

| Position |  TIME |  LANGUAGE  | 
| :------: | :---: |  :-------: | 
|   4/4    | 6:21  |     Go     |

The program:
Your program must convert a sequence of integers into a string of ASCII characters.

#### INPUT:
Line 1: An integer charCount for the number of integers to convert.
Line 2: charCount integers charCode, separated by spaces.

#### OUTPUT:
<pre>
A single line containing a string of characters with the given ASCII codes.
</pre>

#### CONSTRAINTS:
<pre>
0 < charCount < 1000 
32 ≤ charCode ≤ 126
</pre>

#### EXAMPLE:
|                      INPUT                    |     OUTPUT    | 
| :-------------------------------------------: | :-----------: |  
|11 <br>72 101 108 108 111 32 87 111 114 108 100|  Hello World  |
 
 
