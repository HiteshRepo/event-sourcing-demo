package main

import (
	"fmt"
	"github.com/hiteshpattanayak-tw/EventSourcing/repositories"
	"log"
)

func main() {
	wr := repositories.NewWarehouseRepository()

	for {
		fmt.Println()
		fmt.Println("Activities - Press key to perform")
		fmt.Println("---------------------------------")
		fmt.Println("R: Receive Inventory")
		fmt.Println("S: Ship Inventory")
		fmt.Println("A: Adjust Inventory")
		fmt.Println("Q: Check Quantity On Hand")
		fmt.Println("E: Get All Events Chronologically")
		fmt.Println("X: Close")
		fmt.Println("> ")

		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Enter Product Id")
		fmt.Println("> ")

		var id int
		_, err = fmt.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}

		product := wr.Get(id)

		if input == "X" || input == "x" {
			break
		}

		if input == "R" || input == "r" {
			fmt.Println("Enter Quantity")
			fmt.Println("> ")

			var qty int
			_, err = fmt.Scan(&qty)
			if err != nil {
				log.Println("invalid quantity...")
				continue
			}

			product.ReceiveProduct(qty)
			log.Println("received product...")
		}

		if input == "S" || input == "s" {
			fmt.Println("Enter Quantity")
			fmt.Println("> ")

			var qty int
			_, err = fmt.Scan(&qty)
			if err != nil {
				log.Println("invalid quantity...")
				continue
			}

			err = product.ShipProduct(qty)
			if err != nil {
				log.Println("error in shipping: ", err.Error())
				continue
			}

			log.Println("shipped product...")
		}

		if input == "A" || input == "a" {
			fmt.Println("Enter Quantity")
			fmt.Println("> ")

			var qty int
			_, err = fmt.Scan(&qty)
			if err != nil {
				log.Println("invalid quantity...")
				continue
			}

			var reason string
			_, err = fmt.Scan(&reason)
			if err != nil {
				log.Println("invalid reason...")
				continue
			}

			err = product.AdjustInventory(qty, reason)
			if err != nil {
				log.Println("error in adjusting inventory: ", err.Error())
				continue
			}

			log.Println("inventory adjusted...")
		}

		if input == "Q" || input == "q" {
			cs := product.GetCurrentState()
			log.Println("quantity on hand: ", cs.QuantityOnHand)
		}

		if input == "E" || input == "e" {
			events := product.GetEvents()
			for _, e := range events {
				e.Display()
			}
		}

		wr.Save(product)
	}
}
