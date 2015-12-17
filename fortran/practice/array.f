        program array

        integer a(20)

        integer j 

        j = 1
        write(*,*) 'j = ', j
        a(j) = 10

        do 10  i = 1, 10
          write(*,*) 'a(1) : ' , a(1)
          a(1) = a(1) - 1
  10    continue

        stop
        end

