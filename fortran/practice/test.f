        program test
c          parameter (PI = 3.14)

          real x, r, PI

          r = 1
          PI = 3.14
          x = area(PI, r)

          write(*,*) x

          write(*,*) PI 

        end

        function area(p, r)
          p = 12
          area = p * r**2
          return
        end
