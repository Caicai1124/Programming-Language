c       ********************************************************
c       * Programming Language Fall 2015                       *
c       *   Programming Assignment 1                           *
c       *                                                      *
c       *   Caicai CHEN                                        *
c       *     caicai.chen@nyu.edu                              *
c       *     NetId: cc4584                                    *
c       ********************************************************

        program vi assignmentV1

          integer num

          read(*,*) num

          call Calculation(num)

        stop
        end


c       ********************************************************
c       * Begin Calculation                                    *
c       ********************************************************

        subroutine Calculation(num)

c       READ IN ALL ELEMENTS

          integer num, k
          real a(num), sum, cur_max, cur_min, avg
          
          cur_max = 0
          cur_min = 0
          sum = 0
          do 10 k = 1, num 
            read(*,*) a(k)
            sum = sum + a(k)
            if ( k .EQ. 1) then
              cur_min = a(k)
              cur_max = a(k)
            else
              cur_max = max(cur_max, a(k))
              cur_min = min(cur_min, a(k))
            endif

  10      continue

          write(*,*) 'THE NUMBERS I READ: '
          do 20 k = 1, num 
            write(*,'(F64.2)') a(k)
  20      continue;

          avg = sum/num

          call print_result(sum, cur_max, cur_min, avg)

        end subroutine

c       ********************************************************
c       *  If program ends normally                            *
c       *  Print out final result                              *
c       ********************************************************
        subroutine print_result(sum, cur_max, cur_min, avg)
          real sum, cur_max, cur_min, avg

          write(*,'(A,F0.2)') 'Sum: ', sum 
          write(*,'(A,F0.2)') 'Average: ', avg
          write(*,'(A,F0.2)') 'Minimum: ', cur_min 
          write(*,'(A,F0.2)') 'Maximum: ', cur_max

          return
        end subroutine

c       ********************************************************
c       *  If program has invalid input                        *
c       *  Print out ERR and return                            *
c       ********************************************************

        subroutine print_error
          write(*,'(A)') 'ERR'
          return
        end subroutine

