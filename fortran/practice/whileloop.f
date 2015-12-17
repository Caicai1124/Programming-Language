        program whileloop

        integer n

        n = 1 

  10    if ( n .le. 100 ) then
            write(*,*) n
            n = 2 * n
            goto 10
          endif

        stop
        end
