/**
Nicole Jara González

Escriba un programa que utilice una estructura y un arreglo/slice para almacenar en memoria
un inventario de una tienda que vende zapatos. Cada zapato debe contar con la información
de su modelo (marca), su precio y la talla del mismo que debe ir únicamente del 34 al 44.
La estructura debe llamarse "calzado". El programa debe poder almacenar la información
“quemada” para 10+ zapatos del inventario y un stock de N cantidad de unidades de dicho
zapato (quiere decir por ejemplo que puede existir en inventario el zapato marca Nike,
talla 42 y que cuesta 50000 colones, pero además que puede existir no solo un par de
esos, si no N pares del mismo zapato).
*/

package main

import "fmt"

type calzado struct {
	marca  string
	precio int
	talla  int
	stock  int
}

type lista_zapatos []calzado

var lZapatos lista_zapatos

func (c *lista_zapatos) agregarZapatos(marca string, precio int, talla int, stock int) {
	if talla >= 32 && talla <= 44 {
		lZapatos = append(lZapatos, calzado{marca: marca, precio: precio, talla: talla, stock: stock})
	} else {
		fmt.Println("\n--Los zapatos no se agregaron ya que la talla debe estar entre los números 34 y 44")
	}
}

func (c *lista_zapatos) buscarZapatos(marca string) int {
	var result = -1
	var i int
	for i = 0; i < len(*c); i++ {
		if (*c)[i].marca == marca {
			result = i
		}
	}
	return result
}

func (c *lista_zapatos) venderZapatos(marca string, talla int, stock int) {
	var cal = c.buscarZapatos(marca)
	if cal != -1 && stock > 0 {
		if (*c)[cal].stock >= stock && (*c)[cal].talla == talla {
			(*c)[cal].stock = (*c)[cal].stock - stock
			fmt.Println("\n--La venta ha sido realizada exitosamente")
		} else {
			fmt.Println("\n--No se puede vender mayor cantidad de zapatos de los que hay en existencia")
			c.eliminarZapatos()
		}
	}
}

func (c *lista_zapatos) eliminarZapatos() {
	var x int
	for x = 0; x < len(*c); x++ {
		if (*c)[x].stock == 0 {
			*c = append((*c)[:x], (*c)[x+1:]...)
			fmt.Println("\n--El zapato fue desechado ya que no quedaban de esa marca")
		}
	}
}

func llenar() {
	lZapatos.agregarZapatos("New Balance", 25000, 40, 15)
	lZapatos.agregarZapatos("Vans", 33000, 38, 8)
	lZapatos.agregarZapatos("Adidas", 35000, 37, 16)
	lZapatos.agregarZapatos("Puma", 37500, 34, 20)
	lZapatos.agregarZapatos("Reebok", 29500, 38, 31)
	lZapatos.agregarZapatos("Converse", 37000, 39, 29)
	lZapatos.agregarZapatos("Channel", 50000, 39, 16)
	lZapatos.agregarZapatos("Dior", 52000, 39, 2)
	lZapatos.agregarZapatos("Louis Vuitton", 58000, 39, 34)
	lZapatos.agregarZapatos("Prada", 60500, 39, 84)
}

func main() {
	llenar()
	fmt.Println("\n--Los zapatos fueron agregados con éxito.") //Agrega los zapatos con tallas acetadas
	fmt.Println("\n--Los zapatos que con los que cuenta la tienda son los siguientes:", lZapatos)

	lZapatos.agregarZapatos("Dior", 50500, 30, 4) // No se agregarán porque no tienen tallas correctas

	lZapatos.agregarZapatos("Nike", 30000, 44, 10) // Se agregarán  tienen tallas correctas
	fmt.Println("\n--Se agregó una nueva marca de zapatos.", lZapatos)

	lZapatos.venderZapatos("Puma", 34, 4)
	lZapatos.venderZapatos("Dior", 39, 2)
	fmt.Println("\n--Los zapatos que con los que cuenta la tienda luego de las ventas son los siguientes:", lZapatos)

	lZapatos.venderZapatos("Dior", 39, 2)
	fmt.Println("\n--Los zapatos que con los que cuenta la tienda luego de la eliminación son los siguientes:", lZapatos)
}

/** Resultados en consola:

--Los zapatos fueron agregados con éxito.

--Los zapatos que con los que cuenta la tienda son los siguientes: [{New Balance 25000 40 15} {Vans 33000 38 8} {Adidas 35000 37 16} {Puma 37500 34 20} {Reebok 29500 38 31} {Converse 37000 39 29} {Channel 50000 39 16} {Dior 52000 39 2} {Louis Vuitton 58000 39 34} {Prada 60500 39 84}]

--Los zapatos no se agregaron ya que la talla debe estar entre los números 34 y 44

--Se agregó una nueva marca de zapatos. [{New Balance 25000 40 15} {Vans 33000 38 8} {Adidas 35000 37 16} {Puma 37500 34 20} {Reebok 29500 38 31} {Converse 37000 39 29} {Channel 50000 39 16} {Dior 52000 39 2} {Louis Vuitton 58000 39 34} {Prada 60500 39 84} {Nike 30000 44 10}]

--La venta ha sido realizada exitosamente

--La venta ha sido realizada exitosamente

--Los zapatos que con los que cuenta la tienda luego de las ventas son los siguientes: [{New Balance 25000 40 15} {Vans 33000 38 8} {Adidas 35000 37 16} {Puma 37500 34 16} {Reebok 29500 38 31} {Converse 37000 39 29} {Channel 50000 39 16} {Dior 52000 39 0} {Louis Vuitton 58000 39 34} {Prada 60500 39 84} {Nike 30000 44 10}]

--No se puede vender mayor cantidad de zapatos de los que hay en existencia

--El zapato fue desechado ya que no quedaban de esa marca

--Los zapatos que con los que cuenta la tienda luego de la eliminación son los siguientes: [{New Balance 25000 40 15} {Vans 33000 38 8} {Adidas 35000 37 16} {Puma 37500 34 16} {Reebok 29500 38 31} {Converse 37000 39 29} {Channel 50000 39 16} {Louis Vuitton 58000 39 34} {Prada 60500 39 84} {Nike 30000 44 10}]

*/
