c       ********************************************************
c       * Programming Language Fall 2015                       *
c       *   Programming Assignment 1                           *
c       *                                                      *
c       *   Caicai CHEN                                        *
c       *     caicai.chen@nyu.edu                              *
c       *     NetId: cc4584                                    *
c       ********************************************************

        program  assignmentV2

        parameter (INT_MAX = 2147483647)
        parameter (MAX_STR_LEN = 255)

        character num_str*(MAX_STR_LEN) 
        integer num, stt, fin, length
        double precision D_ZERO

        D_ZERO = 0.00

        read(UNIT=*, FMT='(A)', IOSTAT=KODE) num_str
        if (KODE .NE. 0) stop

c       -------------------------------------------------------
c         CHECK WHETHER THE FIRST INTEGER IS VALID INPUT
c         Valid input:    WS INTEGER WS
c           1. Strip head WS
c           2. Strip tail WS 
c           3. Check is_negative
c           4. Check is_invalid_int
c           5. Check is_zero
c           6. Check is_overflow
c           7. Convert into integer
c       -------------------------------------------------------

c       STRIP HEAD WS

        stt = strip_head_ws(num_str, len(num_str))

c       STRIP TAIL WS

        fin = strip_tail_ws(num_str, stt, len(num_str))

        length = fin - stt + 1

        if ( length .le. 0 ) then
          goto 88
        endif

c       CHECK IS_NEGATIVE  

        if (is_negative(num_str(stt:fin)) .eq. 1) then
          goto 88
        endif

c       CHECK IS_INVALID_INT

        if (is_invalid_int(num_str(stt:fin), length) .eq. 1) then
          goto 88
        endif

c       CHECK IS_ZERO
        if (is_zero(num_str(stt:fin), length) .eq. 1) then
c          call print_result(0.00,0.00,0.00,0.00)
          call print_result(D_ZERO, D_ZERO, D_ZERO, D_ZERO)
          goto 100
        endif

c       CHECK OVERFLOW
        if (is_overflow(num_str(stt:fin), length). eq. 1) then
          goto 88
        endif

        read(num_str(stt:fin), '(I10)') num

        call Calculation(num)

        goto 100  
  88    call print_error()
  100   stop
        end

c       ********************************************************
c       * Strip head and tail WS                               *
c       ********************************************************

        function strip_head_ws(s, length)
          character *(*) s
          integer length
          do 10 i = 1, length 
            if ((s(i:i) .ne. ' ') .and. (ichar(s(i:i)) .ne. 9) ) then
              goto 40
            endif
  10      continue

  40      strip_head_ws = i
          return
        end function

        function strip_tail_ws(s, stt, length)
          character *(*) s
          integer stt, length
          do 20 i = length, stt, -1
            if ((s(i:i) .ne. ' ') .and. (ichar(s(i:i)) .ne. 9)) then
              goto 80
            endif
  20      continue

  80      strip_tail_ws = i
          return

          end function
        
c       ********************************************************
c       * Check whether the integer is negative                *
c       ********************************************************

        function is_negative(s)
          character*(*) s

          if (s(1:1) .eq. '-') then
            is_negative = 1
          else
            is_negative = 0
          endif

          return
        end function

c       ********************************************************
c       * Check whether the integer invalid                    *
c       ********************************************************

        function is_invalid_int(s, len)
          character*(*) s
          integer len, num

          do 33 i = 1, len
            num = ichar(s(i:i))
            if ((num .lt. 48) .or. (num .gt. 57)) then
              goto 34
            endif
  33      continue

          is_invalid_int = 0
          return

  34      is_invalid_int = 1
          return
        end function

c       ********************************************************
c       * Check whether the integer is 0                       *
c       ********************************************************

        function is_zero(s,len)
          character*(*) s
          integer len, num

          do 44 i = 1, len
            num = ichar(s(i:i))
            if ( num .ne. 48 ) then
              goto 43
            endif
  44      continue

          is_zero = 1
          return

  43      is_zero = 0
          return
          end function
            
c       ********************************************************
c       * Check whether the integer is 0                       *
c       ********************************************************

        function is_overflow(s,len)
          character*(*) s
          integer len, num, cur_num, n_max
          integer INT_MAX
          INT_MAX = 2147483647

          num = 0
          if (len .lt. 10) then
            goto 45
          else if (len .gt. 10) then
            goto 46
          else
            do 440 i = 1, (len - 1)
              cur_num = ichar(s(i:i)) - 48
              num = num * 10 + cur_num
              n_max = INT_MAX / (10**(10-i))
              if (num .gt. n_max) then
                goto 46
              endif
  440       continue
            if (ichar(s(10:10))-48 .gt. 7) then
              goto 46
            endif
          endif

  45      is_overflow = 0
          return

  46      is_overflow = 1
          return
          end function

c       ********************************************************
c       * Convert String to Integer                            *
c       ********************************************************

        function convert_integer(s,len)
          character*(*) s
          integer len, num, cur_num

          num = 0
          do 58 i = 1, len
            cur_num = ichar(s(i:i)) - 48
            num = num*10 + cur_num
  58      continue

          convert_integer = num

          return
        end function

c       ********************************************************
c       * Begin Calculation                                    *
c       ********************************************************

        subroutine Calculation(num)

c       READ IN ALL ELEMENTS
          parameter (INT_MAX = 2147483647)
          integer num, k
          integer stt, fin, length
          integer sign, pnt

          character kbd_in*255
          character cur_num_buf*255
          double precision cur_num, sum 
          double precision cur_max, cur_min
          double precision pos(num), neg(num)
          integer n_pos, n_neg

          n_pos = 0
          n_neg = 0
          do 10 k = 1, num 
            read(UNIT=*, FMT='(A)', IOSTAT=KODE) kbd_in

            if ( KODE .ne. 0) then
              goto 888
            endif

            stt = strip_head_ws(kbd_in, len(kbd_in))
            fin = strip_tail_ws(kbd_in, stt, len(kbd_in))
            length = fin - stt + 1
            if (length .le. 0) then
              goto 888
            endif

c           Check Input is positive or negative
            if(kbd_in(stt:stt) .eq. '+') then
              sign = 1
              stt = stt+strip_head_ws(kbd_in(stt+1:fin),fin-stt)
            else if(kbd_in(stt:stt) .eq. '-') then
              sign = -1
              stt = stt+strip_head_ws(kbd_in(stt+1:fin),fin-stt)
            else
              sign = 1
            endif

            if(INDEX(kbd_in(stt:fin), '.') .ne. 0) then
              pnt = stt + INDEX(kbd_in(stt:fin), '.') - 1

              if(is_invalid_int(kbd_in(stt:pnt-1),pnt-stt).eq.1)then
                goto 888
              endif

              if(is_overflow(kbd_in(stt:pnt-1),pnt-stt).eq.1)then
                goto 888
              endif

              if(is_invalid_int(kbd_in(pnt+1:fin),fin-pnt).eq.1)then
                goto 888
              endif
            else
              if(is_invalid_int(kbd_in(stt:fin),fin-stt+1).eq.1)then
                goto 888
              endif

              if(is_overflow(kbd_in(stt:fin),fin-stt+1).eq.1)then
                goto 888
              endif
            endif

            if (sign .eq. -1) then
              cur_num_buf(2:) = kbd_in(stt:fin)
              cur_num_buf(1:1) = '-'
            else
              cur_num_buf = kbd_in(stt:fin)
            endif

            if(INDEX(kbd_in(stt:fin), '.') .eq. 0) then
              cur_num_buf((fin+1):(fin+1)) = '.'
            endif
            read(cur_num_buf, '(F60.4)') cur_num

            if ( k .eq. 1 ) then
              cur_min = cur_num
              cur_max = cur_num
            else
              cur_max = max(cur_max, cur_num)
              cur_min = min(cur_min, cur_num)
            endif

            if (sign .eq. 1) then
              n_pos = n_pos + 1
              pos(n_pos) = cur_num
            else
              cur_num_buf = kbd_in(stt:fin)
              if(INDEX(kbd_in(stt:fin), '.') .eq. 0) then
                cur_num_buf((fin+1):(fin+1)) = '.'
              endif
              read(cur_num_buf, '(F60.4)') cur_num
              n_neg = n_neg + 1
              neg(n_neg) = cur_num
            endif

  10      continue

          sum = 0 
          do 72 i = 1, num
            if (sum .gt. 0) then
              if (n_neg .ne. 0) then
                sum = sum - neg(n_neg)
                n_neg = n_neg - 1
              else
                if ( (INT_MAX - sum) .lt. pos(n_pos) ) then
                  goto 888
                else
                  sum = sum + pos(n_pos)
                  n_pos = n_pos - 1
                endif
              endif
            else
              if (n_pos .ne. 0) then
                sum = sum + pos(n_pos)
                n_pos = n_pos - 1
              else
                if ( (sum - INT_MAX) .gt. neg(n_neg)) then
                  goto 888
                else
                  sum = sum - neg(n_neg)
                  n_neg = n_neg - 1
                endif
              endif
            endif
  72      continue

          call print_result(sum, sum/num, cur_min, cur_max)
          return

  888     call print_error
          return

        end subroutine


        character*(*)function get_min(s1,sign,s2)
          character*(*) s1, s2
          integer sign
          get_min = s1
          return
        end function

        character*(*)function int_add(s1, s2)
          character*(*) s1, s2
          int_add = s1
          write(*,'(A)') s2
          return
        end function

        character*(*)function frac_add(s1,s2)
          character*(*) s1, s2
          write(*,'(A)') s2
          frac_add = s1
          return
        end function
          

c       ********************************************************
c       *  If program ends normally                            *
c       *  Print out final result                              *
c       ********************************************************
        subroutine print_result(sum, avg, cur_min, cur_max)
          double precision sum, cur_max, cur_min, avg

          write(*,'(A,F60.2)') 'Sum: ', sum 
          write(*,'(A,F60.2)') 'Average: ', avg
          write(*,'(A,F60.2)') 'Minimum: ', cur_min 
          write(*,'(A,F60.2)') 'Maximum: ', cur_max

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

