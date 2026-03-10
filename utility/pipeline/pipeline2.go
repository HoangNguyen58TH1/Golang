package pipeline

import (
	"fmt"
	"time"
)

type Order struct {
	ID       int
	Price    float64
	Quantity int
	Total    float64
}

func generateOrders(n int) <-chan Order {
	out := make(chan Order)

	go func() {
		for i := 1; i <= n; i++ { // generate order
			out <- Order{
				ID:       i,
				Price:    10.2,
				Quantity: i,
			}
		}
		close(out)
	}()

	return out
}

func validateOrders(in <-chan Order) <-chan Order {
	out := make(chan Order)

	go func() {
		for order := range in {
			if order.Quantity == 0 {
				fmt.Println("Quantity must be greater than 0")
			}

			if order.Price == 0 {
				fmt.Println("Price must be greater than 0")
			}

			if order.Quantity > 0 && order.Price > 0 { // validate
				out <- order
			}
		}
		close(out)
	}()

	return out
}

func calculateTotal(in <-chan Order) <-chan Order {
	out := make(chan Order)

	go func() {
		for order := range in {
			order.Total = order.Price * float64(order.Quantity)
			time.Sleep(200 * time.Millisecond) // simulate heavy processing
			out <- order
		}
		close(out)
	}()

	return out
}

func store(in <-chan Order) {
	for order := range in {
		fmt.Printf("Saving order %d total=%.2f\n", order.ID, order.Total)
	}
}

// - OrderPipeLine(main) --> generateOrders --(ch1)--> validateOrders --(ch2)--> calculateTotal --(ch3)--> store DB
// - STAGE: generateOrders, validateOrders, calculateTotal
// - Channel: ch1, ch2, ch3
func OrderPipeLine() {
	ch1 := generateOrders(3)
	ch2 := validateOrders(ch1)
	ch3 := calculateTotal(ch2)

	store(ch3)
}
