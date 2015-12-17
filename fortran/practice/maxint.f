        program maxint

        integer j, k
        integer i

        j = 32 
        k = -1 + 2**30 + 2**30
        write(*,*) k

        do 10 i = 1, j
          k = 2**i
          write(*,*) k
  10    continue 

        read(*,*) i
        write(*,*) i

        stop
        end
