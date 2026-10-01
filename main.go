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

}
