### do-loop

Actually, there is only this kind of loop in fortran, other kind of loop should be implemented by using "if" and "goto"

This is corresponding to "for-loop"

```
      integer i, n, sum
 
      sum = 0
      do 10 i = 1, n
         sum = sum + i
         write(*,*) 'i =', i
         write(*,*) 'sum =', sum
  10  continue
```

  number "10" here is a statement label

The general form of the do loop is as follows:

```
      do label  var =  expr1, expr2, expr3
         statements
label continue
```

  * var is the loop variable (often called the loop index) which must be integer.

  * expr1 specifies the initial value of var, 

  * expr2 is the terminating bound, and expr3 is the increment (step).

  * Note: The do-loop variable must never be changed by other statements within the loop!

### while-loop

The most intuitive way to write a while-loop is

```
      while (logical expr) do
         statements
      enddo
```
or alternatively,

```
      do while (logical expr) 
         statements
      enddo
```

Correct way:
```
label if (logical expr) then
         statements
         goto label
      endif 
```

### until-loop

```
      do
         statements
      until (logical expr)
```
