# clashOfCode
Repo of clash of code challenges

# 1. Iterative Divisors challenge
In this challenge, we must get the 'x' divisors at a number like this example:


```
                  1
number = a0 + ----------------
                              1
                  a1 = ----------------
                                           1   
                              a2 = ----------------
                                           a3
```

The objective must have a list with 'ai' numbers.
π or e numbers have infinite divisors, this it why give us the lenght of the result list size.

The result for π number with a length result of 5 must be:
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


To emphasize, the code list that we must return must no have any blank spaces.
