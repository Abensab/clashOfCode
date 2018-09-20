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

The main objective of this challenge was chek if an input string is a secure password, for that it was accomplish 4 principles:

1. May be have almost 8 characters
2. One character mus be uppercase
3. One character must be lowercase
4. One character must be a number

If the string accomplish these characteristics, we mus return "true", else "false".
#### * Position:4 Time:10:40
______________________________________________

