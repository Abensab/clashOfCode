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
| Position | TIME  |  LANGUAGE  |     MODE    |
| :------: | :---: | :--------: | :---------: |
|    4     | 10:40 |   Python   |Fastest code |

The main objective of this challenge was chek if an input string is a secure password, for that it was accomplish 4 principles:

1. May be have almost 8 characters
2. One character mus be uppercase
3. One character must be lowercase
4. One character must be a number

If the string accomplish these characteristics, we mus return "true", else "false".


______________________________________________
## 3.Counting ones

| Position |  TIME |  LANGUAGE  |    MODE    |
| :------: | :---: |  :-------: | :---------:|
|   7/8    | 15:00 |     Go     |Fastest code|

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

| Position |  TIME |  LANGUAGE  |    MODE     |
| :------: | :---: |  :-------: | :---------: |
|   4/4    | 6:21  |     Go     |Fastest code |

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
 
 _________________________________
## 5. The result angle
 
| Position |  TIME |  LANGUAGE  |     MODE    |
| :------: | :---: |  :-------: | :---------: |
|   7/8    | 4:37  |     Go     | Fastest code|

#### The program:
Given two angles of a triangle, calculate the third angle.

#### INPUT:
A single line containing the two angles a et b separated by a space.

#### OUTPUT:
Une seule ligne containing the last angle c.

#### CONSTRAINTS:
1≤a<180
1≤b<180

#### EXAMPLE:
|   INPUT   |   OUTPUT  | 
| :-------: | :-------: | 
|   1 1     |    178    | 
______________________________________________
## 6. The middel number

| Position |  TIME |  LANGUAGE  |       MODE        |  CHARACTERS |
| :------: | :---: |  :-------: | :---------------: | :---------: |
|   2/5    | 5:11  |     Go     |   Shortest code   |     362     |

#### The program:
You must output the number in MIDDLE value for each of given lines.
<pre>
Triplet → Output
1 5 6 → 5
8 4 6 → 6
3 1 3 → 3
2 9 2 → 2
6 6 6 → 6
</pre>

#### Input:
Line 1: An integer N for the lines of numbers to proceed.
Next N lines: Three numbers X, Y and Z separated in space.

#### Output:
N lines, each line is the MIDDLE number in value.

#### Constraints
1 ≤ N ≤ 20
-65535 ≤ X,Y,Z ≤ 65535

#### EXAMPLE: 

|                    INPUT                         |          OUTPUT         | 
| :----------------------------------------------: | :---------------------: | 
|5<br>1 2 3<br>5 3 8<br>9 1 7<br>4 2 4<br>6 6 6<br>|<br>2<br>5<br>7<br>4<br>6<br>|

________________________________________
## 7. Get uppercase characters

| Position |  TIME |  LANGUAGE  |       MODE        | 
| :------: | :---: |  :-------: | :---------------: | 
|   4/7    | 4:15  |   Python   |   Fastest code    |   

#### The program:
Your program must output only the capital letters of the string given as input.

#### INPUT:
A string S.

#### OUTPUT:
A string containing S stripped of all characters except capital letters.

#### CONSTRAINTS:
S contains at least 1 capital letter.
S contains less than 100 characters.

#### EXAMPLE:

|                    INPUT                         |          OUTPUT         | 
| :----------------------------------------------: | :---------------------: | 
|           .2A1N5Y64! §C*H*zAtrR|ANYCHAR|
 
 
 ____________________________________________
 ## 8.Get intersections between circles
 
| Position |  TIME |  LANGUAGE  |       MODE        |  CHARACTERS |
| :------: | :---: |  :-------: | :---------------: | :---------: |
|   2/4    | 4:34  |   Python   |   Shortest code   |      58     |

 
You are given a number of circles assuming that none are positioned on top of each other. What is the most number of points that they will intersect at? For example 3 circles: 6 points

#### Input
One int x the number of circles
Output
One int y the maximum number of intersecting points

#### Constraints
1<x<1000000000

#### Example
|                    INPUT                         |          OUTPUT         | 
| :----------------------------------------------: | :---------------------: | 
|3|6|
