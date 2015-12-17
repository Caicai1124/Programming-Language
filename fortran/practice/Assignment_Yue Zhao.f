       PROGRAM ASSIGNMENT1
*array with 1000 numbers
       REAL X(10000)
       CHARACTER CNPTS*255
       INTEGER NPTS
*read number of data points NPTS
       READ(UNIT=*, FMT='(A)', IOSTAT=NZE) CNPTS
       
       INDFIRST = LFIRSTNSPACE(CNPTS)
       INDLAST = LASTNSPACE(CNPTS)

       
       IF (INDEX(CNPTS(INDFIRST:INDLAST),' ') .GT. 0) THEN
           WRITE(UNIT=*,FMT=80)
80         FORMAT('ERR') 
           STOP 
           
       END IF

       IF (INDEX(CNPTS,'+') .GT. 0) THEN
           WRITE(UNIT=*,FMT=81)
81         FORMAT('ERR')
           STOP
       END IF

       READ(CNPTS,'(I20)', IOSTAT=NZE) NPTS



       IF (NZE .NE. 0) THEN
       GOTO 60
       END IF
       IF(NPTS .LT. 0) THEN
          WRITE (UNIT=*, FMT=65)
          STOP
       ELSE IF(NPTS .EQ. 0) THEN
          STOP
       END IF


       DO 90, I=1,NPTS
           READ (UNIT=*, FMT=*, ERR=60) X(I)

90     CONTINUE
       CALL STATS(X, NPTS, SUM, AVG, AMAX, AMIN)
       IF (AMIN .LE. -0.005 .OR. SUM .LE. -0.005) THEN
          WRITE(UNIT=*,FMT=18)
18        FORMAT('ERR')
          STOP
       ELSE IF (AMIN .GT. -0.005 .AND. AMIN .LT. 0 .AND.
     $          AMAX .GT. 0.004) THEN
          AMIN = 0
          WRITE(UNIT=*, FMT=24) SUM
24        FORMAT('Sum:',1X, F20.2)
          WRITE(UNIT=*, FMT=34) AVG
34        FORMAT('Average:',1X, F20.2)
          WRITE(UNIT=*, FMT=44) AMIN
44        FORMAT('Minimum:',1X, F20.2)
          WRITE(UNIT=*, FMT=54) AMAX
54        FORMAT('Maximun:',1X, F20.2)
       ELSE IF (AMIN .GT. -0.005 .AND. AMIN .LT. 0 .AND.
     $          AMAX .GT. -0.005 .AND. AMAX .LT. 0) THEN
          AMAX = 0
          AMIN = 0
          SUM = 0
          AVG = 0
          WRITE(UNIT=*, FMT=28) SUM
28        FORMAT('Sum:',1X, F20.2)
          WRITE(UNIT=*, FMT=38) AVG
38        FORMAT('Average:',1X, F20.2)
          WRITE(UNIT=*, FMT=48) AMIN
48        FORMAT('Minimum:',1X, F20.2)
          WRITE(UNIT=*, FMT=58) AMAX
58        FORMAT('Maximun:',1X, F20.2)



       ELSE

          WRITE(UNIT=*, FMT=26) SUM
26        FORMAT('Sum:',1X, F23.2)
          WRITE(UNIT=*, FMT=36) AVG
36        FORMAT('Average:',1X, F20.2)
          WRITE(UNIT=*, FMT=46) AMIN
46        FORMAT('Minimum:',1X, F20.2)
          WRITE(UNIT=*, FMT=56) AMAX
56        FORMAT('Maximun:',1X, F20.2)
       END IF
       STOP

60     WRITE(UNIT=*,FMT=65)
65     FORMAT('ERR')  
       STOP

       END



       SUBROUTINE STATS(X, NPTS, SUM, AVG, AMAX, AMIN)
       INTEGER NPTS
       REAL X(NPTS), SUM, AVG
       SUM = 0.00
       DO 10, I = 1, NPTS
           SUM = SUM + X(I)
10     CONTINUE
       
       AVG = SUM / NPTS

       AMIN = X(1)
       DO 25, I = 2, NPTS
           IF(X(I) .LE. AMIN) THEN
               AMIN = X(I)
           END IF
25     CONTINUE
       
       AMAX = X(1)
       DO 35, I = 2, NPTS
           IF(X(I) .GE. AMAX) THEN
               AMAX = X(I)
           END IF
35     CONTINUE
       END


             
       INTEGER FUNCTION LFIRSTNSPACE(STRING)
       CHARACTER*(*) STRING
       DO 15, I = 1, LEN(STRING)
          IF(STRING(I:I) .NE. ' ') GO TO 20
15     CONTINUE
20     LFIRSTNSPACE = I
       END

       INTEGER FUNCTION LASTNSPACE(STRING)
       CHARACTER*(*) STRING
       DO 15, I = LEN(STRING),1,-1
          IF(STRING(I:I) .NE. ' ' ) GO TO 20
15     CONTINUE
20     LASTNSPACE = I
       END



