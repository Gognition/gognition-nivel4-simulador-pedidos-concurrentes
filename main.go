package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
	ID       int
	Item     string
	Quantity int
}

func generateOrders(orderChan chan<- Order, r *rand.Rand) {
	for i := 1; ; i++ {
		order := Order{
			ID:       i,
			Item:     fmt.Sprintf("Item-%d", r.Intn(100)),
			Quantity: r.Intn(10) + 1,
		}
		orderChan <- order
		time.Sleep(time.Millisecond * time.Duration(r.Intn(500)))
	}
}

func checkInventory(orderChan <-chan Order, validOrderChan chan<- Order, r *rand.Rand) {
	for order := range orderChan {
		if r.Float32() < 0.8 { // 80% de probabilidad de tener inventario
			validOrderChan <- order
			fmt.Printf("Inventario verificado para el pedido %d\n", order.ID)
		} else {
			fmt.Printf("Inventario insuficiente para el pedido %d\n", order.ID)
		}
		time.Sleep(time.Millisecond * time.Duration(r.Intn(200)))
	}
}

func processPayment(validOrderChan <-chan Order, paidOrderChan chan<- Order, r *rand.Rand) {
	for order := range validOrderChan {
		if r.Float32() < 0.9 { // 90% de probabilidad de pago exitoso
			paidOrderChan <- order
			fmt.Printf("Pago procesado para el pedido %d\n", order.ID)
		} else {
			fmt.Printf("Fallo en el pago para el pedido %d\n", order.ID)
		}
		time.Sleep(time.Millisecond * time.Duration(r.Intn(300)))
	}
}

func prepareShipment(paidOrderChan <-chan Order, r *rand.Rand) {
	for order := range paidOrderChan {
		fmt.Printf("Envío preparado para el pedido %d\n", order.ID)
		time.Sleep(time.Millisecond * time.Duration(r.Intn(400)))
	}
}

func main() {
	fmt.Println("Simulador de Procesamiento de Pedidos Concurrente")

	// Crear una nueva fuente de aleatoriedad
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	orderChan := make(chan Order)
	validOrderChan := make(chan Order)
	paidOrderChan := make(chan Order)

	go generateOrders(orderChan, r)
	go checkInventory(orderChan, validOrderChan, r)
	go processPayment(validOrderChan, paidOrderChan, r)
	go prepareShipment(paidOrderChan, r)

	// Dejamos que el programa corra por un tiempo
	time.Sleep(time.Second * 30)
}
