#### If statement:

Simplest:

  if (logical expression) executable statement

  Example:
```
          if (x .LT. 0) x = -x
```

Common:

```
      if (logical expression) then
         statements
      endif
```

General:

```
      if (logical expression) then
         statements
      elseif (logical expression) then
         statements
       :
       :
      else
         statements
      endif
```

Nested:

```
      if (x .GT. 0) then
         if (x .GE. y) then
            write(*,*) 'x is positive and x >= y'
         else
            write(*,*) 'x is positive but x'
```
