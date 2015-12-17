        program findspace
          character char*3
          read(*,'(A)') char

          write(*, *) INDEX(char, ' ')

        end
