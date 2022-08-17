/**
Nicole Jara González

3)	Escriba una función que permita rotar una secuencia de elementos N posiciones a la izquierda o a la derecha
según sea el caso en función al parámetro que se reciba. Los parámetros deben ser un puntero a un arreglo
previamente creado, la cantidad de movimiento de rotación y la dirección (0 = izq y 1 = der).

A partir de esta función, escriba un programa que haga varias rotaciones cualesquiera de una secuencia
de elementos previamente creada que sea inmutable. Al final debe imprimirse el resultado de cada
rotación además de la secuencia original.

*/

package main

import "fmt"

func rotar(array *[]string, rotacion int, direccion int) {
	lista := []string{}
	for x, y := range *array {
		lista = append(lista, y)
		_ = x
	}

	var x = 0

	for ; x < rotacion; x++ {
		if direccion == 0 { // Izquierda
			lista = append(lista[:0+1], lista[0:]...)
			z := len(lista) - 1
			lista[0] = lista[z]
			lista = append(lista[:z], lista[z+1:]...)
		} else if direccion == 1 { // Derecha
			lista = append(lista, lista[0])
			lista = append(lista[:0], lista[0+1:]...)
		} else {
			fmt.Println("\nLa dirección de rotación está únicamente asignada a los números 0 y 1, debido a que ingresó un número distinto, el proceso en donde se dio dicha asignación no será efectuado.")
			return
		}
	}
	fmt.Println("Nuevo arreglo", lista, "con su respectiva rotación de", rotacion, "en dirección", direccion, "(0 = izq / 1 = der)")
}

func main() {
	letras := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	fmt.Println("\nArreglo original", letras)
	rotar(&letras, 3, 0)
	rotar(&letras, 3, 1)
	rotar(&letras, 3, 2)

}

/** Resultados en consola:

Arreglo original [a b c d e f g h]
Nuevo arreglo [f g h a b c d e] con su respectiva rotación de 3 en dirección 0 (0 = izq / 1 = der)
Nuevo arreglo [d e f g h a b c] con su respectiva rotación de 3 en dirección 1 (0 = izq / 1 = der)

La dirección de rotación está únicamente asignada a los números 0 y 1, debido a que ingresó un número distinto, el proceso en donde se dio dicha asignación no será efectuado.

*/
