module Main where

-- Nicole Jara González

{-3) Implemente la función aplanar. Esta función recibe una lista con múltiples listas anidadas
como elementos y devuelve una lista con los mismos elementos de manera lineal (sin listas).

Por ejemplo:

aplanar ‘(1 2 (3 (4 5) (6 7))))
(1 2 3 4 5 6 7)

En el caso de que una lista de listas a muchos niveles represente un reto complejo puede 
considerar al menos que haya listas de listas a un nivel y no más profundo.-}

aplanarlist :: [[z]] -> [z]

aplanarlist [] = []

aplanarlist (z:zs) = z ++ aplanarlist zs

main :: IO ()
main = do 
    
    print(aplanarlist[[1, 2], [3, 4, 5], [6, 7]])

{-Resultados:

[1,2,3,4,5,6,7]

-}