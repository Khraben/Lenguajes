package main

import "fmt"

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	// Buscar si el producto ya existe
	index := l.buscarProducto(nombre)
	if index != -1 {
		// Si el producto ya existe, incrementar la cantidad y actualizar el precio
		(*l)[index].cantidad += cantidad
		if (*l)[index].precio != precio {
			(*l)[index].precio = precio
		}
	} else {
		// Si el producto no existe, agregarlo a la lista
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}
}

// Hacer una función adicional que reciba una cantidad potencialmente infinita de productos previamente creados y los agregue a la lista
func (l *listaProductos) agregarProductosInfinitos(productos ...producto) {
	for _, p := range productos {
		index := lProductos.buscarProducto(p.nombre)
		if index != -1 {
			// Si el producto ya existe, incrementar la cantidad y actualizar el precio
			lProductos[index].cantidad += p.cantidad
			if lProductos[index].precio != p.precio {
				lProductos[index].precio = p.precio
			}
		} else {
			// Si el producto no existe, agregarlo a la lista
			lProductos = append(lProductos, p)
		}
	}
}

func (l *listaProductos) buscarProducto(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

func (l *listaProductos) venderProducto(nombre string, cantidadVender int) {
	//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista" y notificar dicha acción
	//modifique posteriormente para que se ingrese de parámetro la cantidad de productos a vender
	var prod = l.buscarProducto(nombre)
	if prod != -1 {
		if (*l)[prod].cantidad == 0 {
			fmt.Println("Producto agotado")
			*l = append((*l)[:prod], (*l)[prod+1:]...)
		} else if (*l)[prod].cantidad < cantidadVender {
			fmt.Println("No cantidad suficiente de este producto")
		} else {
			(*l)[prod].cantidad = (*l)[prod].cantidad - cantidadVender
		}
	}
}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	var listaMinimos listaProductos
	for i := 0; i < len(*l); i++ {
		if (*l)[i].cantidad < existenciaMinima {
			listaMinimos = append(listaMinimos, (*l)[i])
		}
	}
	return listaMinimos
}

// modifique la función para que esta vez solo retorne el producto como tal y que retorne otro valor adicional "err" conteniendo un 0 si no hay error y números mayores si existe algún
// tipo de error como por ejemplo que el producto no exista. Una vez implementada la nueva función, cambie los usos de la anterior función en funciones posteriores, realizando los cambios
// que el uso de la nueva función ameriten
func (l *listaProductos) buscarProducto2(nombre string) (int, int) {
	var result int
	var err = 1 // Valor de error por defecto, indicando que el producto no se encontró
	for i := 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
			err = 0 // Cambiar el valor de error a 0, indicando que se encontró el producto
			break
		}
	}
	return result, err
}

func (l *listaProductos) venderProducto2(nombre string, cantidadVender int) {
	prod, err := l.buscarProducto2(nombre)
	if err == 0 {
		if (*l)[prod].cantidad == 0 {
			fmt.Println("Producto agotado")
			// Eliminar el producto de la lista
			*l = append((*l)[:prod], (*l)[prod+1:]...)
			fmt.Printf("El producto %s se ha agotado y ha sido eliminado de la lista.\n", nombre)
		} else if (*l)[prod].cantidad < cantidadVender {
			fmt.Println("No hay suficiente cantidad de este producto")
		} else {
			(*l)[prod].cantidad -= cantidadVender
		}
	} else {
		fmt.Println("Producto no encontrado")
	}
}

func (l *listaProductos) listarProductosMínimos2() listaProductos {
	var listaMinimos listaProductos
	for i := 0; i < len(*l); i++ {
		prod, err := l.buscarProducto2((*l)[i].nombre)
		if err == 0 && (*l)[prod].cantidad < existenciaMinima {
			listaMinimos = append(listaMinimos, (*l)[prod])
		}
	}
	return listaMinimos
}

// haga una función para a partir del nombre del producto, se pueda modificar el precio del mismo Pero utilizando la función buscarProducto MODIFICADA PREVIAMENTE
func (l *listaProductos) modificarPrecioProducto(nombre string, nuevoPrecio int) {
	prod, err := l.buscarProducto2(nombre)
	if err == 0 {
		(*l)[prod].precio = nuevoPrecio
	} else {
		fmt.Println("Producto no encontrado")
	}
}

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)

}

func main() {
	fmt.Println("\nLista de Productos Inicial:")
	llenarDatos()
	fmt.Println(lProductos)

	//probar la funcion de productos infinitos
	fmt.Println("\nLista Tras Agregar 2 Productos Con La Funcion agregarProductosInfinitos:")
	producto1 := producto{nombre: "pasta", cantidad: 20, precio: 1800}
	producto2 := producto{nombre: "azúcar", cantidad: 30, precio: 3000}
	lProductos.agregarProductosInfinitos(producto1, producto2)
	fmt.Println(lProductos)
	//venda productos
	fmt.Println("\nSe intenta vender 10 de café:")
	lProductos.venderProducto("café", 10)
	fmt.Println(lProductos)
	//verificar que no hay unidades suficientes
	fmt.Println("\nSe intenta vender 5 de café:")
	lProductos.venderProducto("café", 5)
	fmt.Println(lProductos)
	//Para dejar las unidades en 0
	fmt.Println("\nSe intenta vender 2 de café:")
	lProductos.venderProducto2("café", 2)
	fmt.Println(lProductos)
	//Para probar si se elimina si no hay unidades
	fmt.Println("\nSe intenta vender 8 de café:")
	lProductos.venderProducto2("café", 8)
	fmt.Println(lProductos)
	//listarProdcutosMinimos
	fmt.Println("\nProductos Con Existencias Minimas:")
	fmt.Print(lProductos.listarProductosMínimos())
	//modificar un precio de producto
	fmt.Println("\n\nSe Modifica el Precio de pasta a 1000:")
	lProductos.modificarPrecioProducto("pasta", 1000)
	fmt.Println(lProductos)
}
