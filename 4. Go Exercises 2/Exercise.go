package main

import (
	"fmt"
	"sort"
	"strings"
)

type infoCliente struct {
	nombre string
	correo string
	edad   int32
}
type listaClientes []infoCliente

var lClientes listaClientes

// funcion para agregar clientes a la lista
func (l *listaClientes) agregarCliente(nombre string, correo string, edad int32) {
	*l = append(*l, infoCliente{nombre: nombre, correo: correo, edad: edad})
}

// funciones genericas hechas en clase
func genericFilter[P1 any](list []P1, f func(P1) bool) []P1 {
	filtered := make([]P1, 0)
	for _, e := range list {
		if f(e) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

func genericMap[P1, P2 any](list []P1, f func(P1) P2) []P2 {
	mapped := make([]P2, len(list))
	for i, e := range list {
		mapped[i] = f(e)
	}
	return mapped
}

func genericReduce[P1, P2 any](list []P1, f func(P2, P1) P2) P2 {
	var result P2
	for _, e := range list {
		result = f(result, e)
	}
	return result
}

func (l *listaClientes) listaClientes_ApellidoEnCorreo() []infoCliente {
	resultado := make([]infoCliente, 0)
	for _, cliente := range *l {
		// Extraer el apellido del nombre del cliente y convertirlo a minúsculas
		nombreSeparado := strings.Split(cliente.nombre, " ")
		if len(nombreSeparado) > 1 {
			apellido := strings.ToLower(nombreSeparado[len(nombreSeparado)-1])
			// Convertir el correo electrónico a minúsculas
			correo := strings.ToLower(cliente.correo)
			// Verificar si el apellido está presente en el correo
			if strings.Contains(correo, apellido) {
				resultado = append(resultado, cliente)
			}
		}
	}
	return resultado
}

func (l *listaClientes) cantidadCorreosCostaRica() int {
	// Filtrar clientes cuyos correos terminan en ".cr"
	clientesCR := genericFilter(*l, func(cliente infoCliente) bool {
		return strings.HasSuffix(cliente.correo, ".cr")
	})

	// Mapear los clientes filtrados a un valor de 1
	mappedClientes := genericMap(clientesCR, func(cliente infoCliente) int {
		return 1
	})

	// Reducir la lista de clientes mapeados sumando la cantidad
	cantidad := genericReduce(mappedClientes, func(acc, val int) int {
		return acc + val
	})

	return cantidad
}

func (l *listaClientes) clientesSugerenciasCorreos() listaClientes {
	// Filtrar los clientes cuyo correo no contiene el nombre
	clientesSinNombre := genericFilter(*l, func(cliente infoCliente) bool {
		nombreSeparado := strings.Split(cliente.nombre, " ")
		nombreCorto := strings.ToLower(string(nombreSeparado[0][0]) + nombreSeparado[len(nombreSeparado)-1])
		return !strings.Contains(cliente.correo, nombreCorto)
	})

	// Generar los nuevos correos sugeridos y mapearlos a la nueva lista de clientes
	nuevosClientes := genericMap(clientesSinNombre, func(cliente infoCliente) infoCliente {
		nombreSeparado := strings.Split(cliente.nombre, " ")
		nombreCorto := strings.ToLower(string(nombreSeparado[0][0]) + nombreSeparado[len(nombreSeparado)-1])
		nuevoCorreo := nombreCorto + "@" + strings.Split(cliente.correo, "@")[1]
		return infoCliente{
			nombre: cliente.nombre,
			correo: nuevoCorreo,
			edad:   cliente.edad,
		}
	})

	return nuevosClientes
}

func (l *listaClientes) correosOrdenadosAlfabeticos() []string {
	// Extraer los correos de los clientes
	correoList := make([]string, 0)
	for _, cliente := range *l {
		correoList = append(correoList, cliente.correo)
	}

	// Ordenar los correos alfabéticamente
	sort.Strings(correoList)

	return correoList
}

func main() {
	fmt.Println("SE REALIZAN LAS 10 INSERCIONES SOLICITADAS EN LA PARTE 1 USANDO 'agregarCliente':")
	fmt.Println("Se inserta a Oscar Viquez:")
	lClientes.agregarCliente("Oscar Viquez", "oviquez@tec.ac.cr", 44)
	fmt.Println(lClientes)
	fmt.Println("Se inserta a Pedro Perez:")
	lClientes.agregarCliente("Pedro Perez", "elsegundo@gmail.com", 30)
	fmt.Println(lClientes)
	fmt.Println("Se inserta a Maria Lopez:")
	lClientes.agregarCliente("Maria Lopez", "mlopez@hotmail.com", 18)
	fmt.Println(lClientes)
	fmt.Println("Se inserta a Juan Rodriguez:")
	lClientes.agregarCliente("Juan Rodriguez", "jrodriguez@gmail.com", 25)
	fmt.Println(lClientes)
	fmt.Println("Se inserta a Luisa Gonzalez:")
	lClientes.agregarCliente("Luisa Gonzalez", "muyseguro@tec.ac.cr", 67)
	fmt.Println(lClientes)
	fmt.Println("Se inserta a Marco Rojas:")
	lClientes.agregarCliente("Marco Rojas", "loquesea@hotmail.com", 47)
	fmt.Println(lClientes)
	fmt.Println("Se inserta a Marta Saborio:")
	lClientes.agregarCliente("Marta Saborio", "msaborio@gmail.com", 33)
	fmt.Println(lClientes)
	fmt.Println("Se inserta a Camila Segura:")
	lClientes.agregarCliente("Camila Segura", "csegura@ice.co.cr", 19)
	fmt.Println(lClientes)
	fmt.Println("Se inserta a Fernando Rojas:")
	lClientes.agregarCliente("Fernando Rojas", "frojas@estado.gov", 27)
	fmt.Println(lClientes)
	fmt.Println("Se inserta a Rosa Ramirez:")
	lClientes.agregarCliente("Rosa Ramirez", "risuenna@gmail.com", 50)
	fmt.Println(lClientes)

	fmt.Println("\nSE PRUEBA LA FUNCION 'listaClientes_ApellidoEnCorreo' COMO SE SOLICITA EN LA PARTE2:")
	fmt.Println(lClientes.listaClientes_ApellidoEnCorreo())

	fmt.Println("\nSE PRUEBA LA FUNCION 'cantidadCorreosCostaRica' COMO SE SOLICITA EN LA PARTE3:")
	fmt.Println("Cantidad de clientes con correos de Costa Rica:", lClientes.cantidadCorreosCostaRica())

	fmt.Println("\nSE PRUEBA LA FUNCION 'clientesSugerenciasCorreos' COMO SE SOLICITA EN LA PARTE 4:")
	fmt.Println(lClientes.clientesSugerenciasCorreos())

	fmt.Println("\nSE PRUEBA LA FUNCION 'correosOrdenadosAlfabeticos' COMO SE SOLICITA EN LA PARTE 5:")
	fmt.Println(lClientes.correosOrdenadosAlfabeticos())
}
