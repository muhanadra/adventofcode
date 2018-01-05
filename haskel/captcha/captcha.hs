import System.Environment

-- A function to convert string of numbers into an array of integers
digitize x = map (read . (:"")) x :: [Int]

-- Create a new list with that unshift every element by 1 pos and put the 1st element in the end.
sol1 a = zip a ( tail a ++ take 1 a)

sol2 a = let len = length a `div` 2
   in zip (take len a) (reverse $ take len $ reverse a)
isEqual x =  if fst x == snd x then True else False

captcha1 x =  sum $ map fst $ filter isEqual ( sol1 (digitize x))
captcha2 x =  sum $ map fst $ filter isEqual ( sol2 (digitize x))
--captchaSum x = a <- (prepInput x)
main = do x <- getArgs
          print (captcha2 $ head x)
