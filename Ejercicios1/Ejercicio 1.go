/**
Nicole Jara González

Haga un programa que cuente e indique el número de caracteres, el número de palabras y el
número de líneas de un texto cualquiera (obtenido de cualquier forma que considere).
Considere hacer el cálculo de las tres variables en el mismo proceso.
*/

package main

import (
	"fmt"
	"strings"
)

func datos(texto string) {
	caracteres := 0
	caracter := strings.Fields(texto)
	numcaracter := strings.Join(caracter, "")
	palabras := strings.Split(texto, " ")
	lineas := strings.Split(texto, "\n")

	for range texto {
		caracteres++
	}

	fmt.Println("El texto dado cuenta con", caracteres, "caracteres,", len(palabras), "palabras y", len(lineas), "líneas. En este texto los espacios son contados como caracter.")

	fmt.Println("\nEl texto dado cuenta con", len(numcaracter), "caracteres,", len(palabras), "palabras y", len(lineas), "líneas. En este texto los espacios no son contados como caracter.")

}

func main() {
	texto := `Dos hermanas, dos ciudades y un descubrimiento que cambiará el mundo para siempre.
			En las ciudades de Piltover y Zaun, se palpa el desasosiego en el ambiente:
			inventores, ladrones, políticos y señores del crimen buscan liberarse de las
			ataduras de una sociedad fragmentada. Mientras la rebelión va cobrando fuerza,
			dos hermanas roban un artefacto de poder inimaginable.
			Los descubrimientos y el peligro son el trasfondo sobre el que nacerán héroes y se
			romperán vínculos. ¿Servirá este poder para cambiar el mundo o lo llevará a la ruina?
			Este es el mundo de Arcane.`

	datos(texto)
}

/** Resultados en consola:

El texto dado cuenta con 587 caracteres, 86 palabras y 8 líneas. En este texto los espacios son contados como caracter.

El texto dado cuenta con 485 caracteres, 86 palabras y 8 líneas. En este texto los espacios no son contados como caracter.

*/
