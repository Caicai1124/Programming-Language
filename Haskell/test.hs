import System.IO     
import Data.Char 

tt :: String -> IO()
tt "TEST"
  print TEST

main = do  
    putStrLn "Hello, what's your name?"  
    name <- getLine  
    putStrLn $ "Read this carefully, because this is your future: "
    tt ("ABC")

