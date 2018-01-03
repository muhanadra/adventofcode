import System.Environment

-- A function to convert string of numbers into an array of integers
digitize x = map (read . (:"")) x :: [Int]


-- Create a new list with that unshift every element by 1 pos and put the 1st element in the end.
tuppleIt a = zip a ( tail a ++ take 1 a)

isEqual x =  if fst x == snd x then True else False

captcha x =  sum $ map fst $ filter isEqual ( tuppleIt (digitize x))
--captchaSum x = a <- (prepInput x)
main = do x <- getArgs
          print (captcha $ head x)

