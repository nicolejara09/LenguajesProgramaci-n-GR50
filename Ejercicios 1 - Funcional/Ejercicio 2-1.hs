module Main where

-- Nicole Jara González

{-2) Construya una función que se llame sub_conjunto. Esta función recibe dos listas y debe retornar 
True cuando el primer argumento es subconjunto completo del segundo y #f en caso contrario. 

Por ejemplo: 
sub_conjunto [] [1,2,3,4,5]
True
sub_conjunto [1,2,3] [1,2,3,4,5]
True
sub_conjunto [1,2,6] [1,2,3,4,5]
False-}
    
sub_conjunto :: (Eq a) => [a] -> [a] -> Bool

sub_conjunto (z:zs) (w:ws)

    | z == w = sub_conjunto zs ws
    | otherwise = sub_conjunto (z:zs) ws

sub_conjunto [] [] = True
sub_conjunto [] _ = True
sub_conjunto _ [] = False

main :: IO ()
main = do 

    print(sub_conjunto [1,2,3][1,2,3,4,5])
    print(sub_conjunto [1,2,6][1,2,3,4,5])
    print(sub_conjunto [][1,2,3,4,5])

{-Resultados:

True
False
True

-}