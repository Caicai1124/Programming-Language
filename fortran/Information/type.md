#### Most common declaration

      integer   list of variables
      real      list of variables
      double precision  list of variables
      complex   list of variables
      logical   list of variables
      character list of variables

#### Numbers casting

      int
      real
      dble
      ichar
      char

Following is different:

w = dble(x)*dble(y)

w = dble(x*y)

#### Logical Expression

      .LT.  meaning <.LE. <=".GT.">
      .GE.          >=
      .EQ.          =
      .NE.          /=

      logical a, b
      a = .TRUE.
      b = a .AND. 3 .LT. 5/2

The order of precedence is important, as the last example shows. The rule is that arithmetic expressions are evaluated first, then relational operators, and finally logical operators. Hence b will be assigned .FALSE. in the example above. Among the logical operators the precedence (in the absence of parenthesis) is that .NOT. is done first, then .AND., then .OR. is done last.

Logical variables are seldom used in Fortran. But logical expressions are frequently used in conditional statements like the if statement.
