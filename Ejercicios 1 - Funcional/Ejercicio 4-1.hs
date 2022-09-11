module Main where

-- Nicole Jara González

{-4) Implemente el equivalente para la función anterior, pero utilizando la función map para dicho fin.-}

mapA :: (z -> w -> w) -> w -> [z] -> w

mapA z w = dato
    where
        dato [] = w
        dato (y:ys) = y `z` dato ys
    
aplanar :: [[x]] -> [x] 
aplanar = mapA (++) []

main :: IO ()
main = do 

    print(aplanar[[1, 2], [3, 4, 5], [6, 7]])

{-Resultados:

[1,2,3,4,5,6,7]

-}