        program stringtoint
        character  str*3 , test*5
        integer num
        read(UNIT= *, FMT = '(A)', IOSTAT = KODE) str
        if (KODE .eq. 0) write(*,*) KODE

        read(*, '(A)') test

c        test(3:5) = str(1:3)
        test(1:1) = 'a'
        test(2:) = str
        test = test(3:3)

        write(*,'(A)') test
        end program
