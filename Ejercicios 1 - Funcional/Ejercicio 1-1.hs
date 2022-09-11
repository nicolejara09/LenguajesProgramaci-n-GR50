module Main (main) where

-- Nicole Jara González

{-1) Haciendo uso de la función filter, implemente una función que a partir de una lista de cadenas de 
parámetro, filtre aquellas que contengan una subcadena que el usuario indique en otro argumento.

Por ejemplo: 
sub_cadenas “la” [“la casa, “el perro”, “pintando la cerca”]
[“la casa, “pintando la cerca”]-}

filterCadena :: (x -> Bool) -> [x] -> [x]

filterCadena dato [] = []

filterCadena dato (z:zs)
    | dato z = z : filterCadena dato zs
    | otherwise = filterCadena dato zs

cadena :: String -> String -> Bool
cadena (_:_) [] = False
cadena zs ws
    | dato1 zs ws = True
    | cadena zs (tail ws) = True
    | otherwise = False

dato1 :: String -> String -> Bool
dato1 [] _ = True
dato1 (_:_) [] = False
dato1 (z:zs) (w:ws) = (z == w) && dato1 zs ws

sub_cadenas :: String -> [String] -> [String]
sub_cadenas str texto = filterCadena (cadena str) texto

main :: IO ()
main = do 
    
    print(sub_cadenas "la" ["la casa", "el perro", "pintando la cerca"])

{-Resultados:

["la casa","pintando la cerca"]

-}