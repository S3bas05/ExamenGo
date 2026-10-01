package main

import "fmt"

var productosvendidos []string
var subtotales []float64

func RegistrarVenta(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)
	productosvendidos = append(productosvendidos, nombre)
	subtotales = append(subtotales, subtotal)

	fmt.Println("Venta registrada correctamente")
	fmt.Printf("Producto:%s\n", nombre)
	fmt.Printf("Subtotal:%.2f\n", subtotal)
	fmt.Printf("Cantidad:%d\n", cantidad)
}

func MostrarEstadisticas() {
	if len(productosvendidos) == 0 {
		fmt.Println("No se han registrado ventas aún.")
		return
	}

	total := 0.0

	fmt.Println("\n----------Estadisticas de Ventas----------")

	for i := 0; i < len(productosvendidos); i++ {
		fmt.Printf("%d. %s - $%.2f\n",
			i+1,
			productosvendidos[i],
			subtotales[i])

		total += subtotales[i]
	}

	fmt.Printf("Total recaudado: $%2f\n", total)

}

func main() {
	var opcion int

	fmt.Println("--------Bienvenido al programa----------")
	fmt.Println("1. Registrar una venta")
	fmt.Println("2. Mostrar estadisticas")
	fmt.Println("3. Salir")
	fmt.Println("Seleccione una opción")

	fmt.Scan(&opcion)

	switch {
	case 1:
		var producto int
		var cantidad int

		fmt.Println("\n------ PRODUCTOS ------")
		fmt.Println("1. Arroz  - $1.25")
		fmt.Println("2. Leche  - $0.95")
		fmt.Println("3. Pan    - $0.50")

		fmt.Print("Seleccione un producto: ")
		fmt.Scan(&producto)

		fmt.Print("Ingrese la cantidad vendida: ")
		fmt.Scan(&cantidad)

		if cantidad <= 0 {
			fmt.Println("La cantidad debe ser mayor a 0.")
		}

		switch producto {

		case 1:
			RegistrarVenta("Arroz", 1.25, cantidad)

		case 2:
			RegistrarVenta("Leche", 0.95, cantidad)

		case 3:
			RegistrarVenta("Pan", 0.50, cantidad)

		default:
			fmt.Println("Producto no válido.")
		}
	}

}
