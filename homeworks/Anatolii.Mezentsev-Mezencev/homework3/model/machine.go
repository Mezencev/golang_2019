package model

import (
	"fmt"
	"sync"
)

// VendingMachine struct
type VendingMachine struct {
	Backed [][]string
}

// SearchOrder
func (c *VendingMachine) SearchOrder(receivedOrder []string, wg *sync.WaitGroup) {

	defer wg.Done()
	complete := []string{}
	lengthOrder := len(receivedOrder)

	for i := 0; i < len(c.Backed); i++ {
		if len(complete) == lengthOrder {
			break
		}
		for k := 0; k < len(c.Backed[i]); k++ {
			if len(complete) == lengthOrder {
				break
			}
			for index := 0; index < len(receivedOrder); index++ {

				if receivedOrder[index] == c.Backed[i][k] {
					complete = append(complete, c.Backed[i][k])
					receivedOrder = append(receivedOrder[:index], receivedOrder[index+1:]...)
					c.Backed[i] = append(c.Backed[i][:k], c.Backed[i][k+1:]...)
				}
			}

		}
	}
	fmt.Println("status Vending Machine", c.Backed)
	fmt.Println("status order", complete)
}
