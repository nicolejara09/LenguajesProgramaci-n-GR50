package main

import (
	"fmt"
	"sort"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	g := (*l).buscarProducto(nombre)
	if g == -1 {
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	} else {
		if (*l)[g].nombre == nombre {
			i := (*l)[g].cantidad + cantidad
			(*l)[g].cantidad = i
			if (*l)[g].precio != precio {
				(*l)[g].precio = precio
			} else {
				k := (*l)[g].precio
				(*l)[g].precio = k
			}
		}
		fmt.Println("El proceso de inserción de datos fue realizado con éxito.")
	}
}

func (l *listaProductos) buscarProducto(nombre string) int {
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

func (l *listaProductos) venderProducto(nombre string, cant int) {
	var prod = l.buscarProducto(nombre)
	if prod != -1 && cant > 0 {
		if (*l)[prod].cantidad >= cant {
			(*l)[prod].cantidad = (*l)[prod].cantidad - cant
		} else if (*l)[prod].cantidad == 0 {
			l.remove()
		} else {
			fmt.Println("No se puede vender mayor cantidad de productos de los que hay en existencia")
			return
		}
	}
}

func (l *listaProductos) remove() {
	var x int
	for x = 0; x < len(*l); x++ {
		if (*l)[x].cantidad == 0 {
			*l = append((*l)[:x], (*l)[x+1:]...)
			fmt.Println("\nEl producto fue desechado ya que no quedaban unidades disponibles")
		}
	}
}

func (l *listaProductos) listarProductosMinimos() listaProductos {
	var productosMin listaProductos
	var x = 0

	for ; x < len(*l); x++ {
		if (*l)[x].cantidad <= existenciaMinima { // La existencia mínima va de 0 a 10
			productosMin = append(productosMin, (*l)[x])
		}
	}
	fmt.Println("\nLa siguiente lista está compuesta por productos que cuentan con existencias mínimas")
	return productosMin
}

/**

Nicole Jara González

a.	A partir de la lista de productos con mínimas existencias de inventario, ampliar dicho
inventario con la compra de más unidades de dicho producto hasta que cumplan con el mínimo
establecido de manera constante.
*/

func (l *listaProductos) aumentarInventarioDeMinimos() {
	var x = 0

	for ; x < len(*l); x++ {
		if (*l)[x].cantidad < existenciaMinima {
			(*l)[x].cantidad = existenciaMinima
			fmt.Println("El número de unidades de", (*l)[x].nombre, "aumentó a el número máximo de existencias mínimas (", existenciaMinima, ")")
		}
	}
	return
}

/**
b.	Crear una función que ordene la lista de productos usando como llave para el ordenamiento
cualquiera de los elementos de la estructura producto. La lista/slice debe quedar modificada al
finalizar el método. Se solicita investigar y hacer uso de alguna función predefinida de alguna
librería del lenguaje Go.
*/

func (l *listaProductos) ordenarProductos() {
	sort.Slice(*l, func(i, j int) bool {
		return (*l)[i].precio < (*l)[j].precio
	})
	fmt.Println("El ordenamiento ascendente de los productos fue realizado priorizando el precio de los mismos")
	fmt.Print("\n")
	fmt.Println(lProductos)
}

func llenarDatos() {
	lProductos.agregarProducto("Arroz", 15, 2500)
	lProductos.agregarProducto("Frijoles", 4, 2000)
	lProductos.agregarProducto("Leche", 9, 1200)
	lProductos.agregarProducto("Café", 12, 5500)
	lProductos.agregarProducto("Maruchan", 50, 500)
	lProductos.agregarProducto("Cerveza", 20, 1000)
	lProductos.agregarProducto("Uvas", 17, 850)
	lProductos.agregarProducto("Chocolate", 24, 1500)
	fmt.Println("Los datos fueron subidos correctamente")
}

func main() {
	llenarDatos()
	fmt.Println(lProductos, "\n")
	lProductos.agregarProducto("Café", 12, 5500)
	lProductos.agregarProducto("Almendras", 10, 5000)
	lProductos.agregarProducto("Naranja", 28, 350)
	lProductos.agregarProducto("Vinagre", 14, 750)
	fmt.Println(lProductos)
	lProductos.venderProducto("Café", 24)
	lProductos.venderProducto("Café", 1) // Se elimina
	fmt.Println(lProductos)
	fmt.Println(lProductos.listarProductosMinimos())
	fmt.Print("\n")
	lProductos.aumentarInventarioDeMinimos()
	fmt.Println(lProductos)
	fmt.Print("\n")
	lProductos.ordenarProductos()

}

// si ha sobrado tiempo antes del receso, el ejercicio se podría ampliar para que los los productos se almacenen en archivos de texto
// que al inicio se carguen del archivo a la lista y que al final se actualice el archivo con el contenido de la lista

/** Resultados en consola:

Los datos fueron subidos correctamente
[{Arroz 15 2500} {Frijoles 4 2000} {Leche 9 1200} {Café 12 5500} {Maruchan 50 500} {Cerveza 20 1000} {Uvas 17 850} {Chocolate 24 1500}]

El proceso de inserción de datos fue realizado con éxito.
[{Arroz 15 2500} {Frijoles 4 2000} {Leche 9 1200} {Café 24 5500} {Maruchan 50 500} {Cerveza 20 1000} {Uvas 17 850} {Chocolate 24 1500} {Almendras 10 5000} {Naranja 28 350} {Vinagre 14 750}]

El producto fue desechado ya que no quedaban unidades disponibles
[{Arroz 15 2500} {Frijoles 4 2000} {Leche 9 1200} {Maruchan 50 500} {Cerveza 20 1000} {Uvas 17 850} {Chocolate 24 1500} {Almendras 10 5000} {Naranja 28 350} {Vinagre 14 750}]

La siguiente lista está compuesta por productos que cuentan con existencias mínimas
[{Frijoles 4 2000} {Leche 9 1200} {Almendras 10 5000}]

El número de unidades de Frijoles aumentó a el número máximo de existencias mínimas ( 10 )
El número de unidades de Leche aumentó a el número máximo de existencias mínimas ( 10 )
[{Arroz 15 2500} {Frijoles 10 2000} {Leche 10 1200} {Maruchan 50 500} {Cerveza 20 1000} {Uvas 17 850} {Chocolate 24 1500} {Almendras 10 5000} {Naranja 28 350} {Vinagre 14 750}]

El ordenamiento ascendente de los productos fue realizado priorizando el precio de los mismos.
[{Naranja 28 350} {Maruchan 50 500} {Vinagre 14 750} {Uvas 17 850} {Cerveza 20 1000} {Leche 10 1200} {Chocolate 24 1500} {Frijoles 10 2000} {Arroz 15 2500} {Almendras 10 5000}]

*/
