import System.IO  
import System.Directory  
import Data.List 
import Data.Char

main = do
  sumTo20 [4,9,10,2,8] 
  where  
    sumTo20 :: [Int] -> Int
    sumTo20 nums = go 0 nums   -- the acc. starts at 0
      where go :: Int -> [Int] -> Int
            go acc [] = acc   -- empty list: return the accumulated sum
            go acc (x:xs)
             | acc >= 20 = acc
             | otherwise = go (acc + x) xs

