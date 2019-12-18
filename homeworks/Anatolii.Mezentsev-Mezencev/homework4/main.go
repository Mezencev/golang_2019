package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

// Const status order
const (
	Pending = "pending"
	Done    = "done"
)

// Order struct for map
type Order struct {
	ID        string `json:"id"`         // id order
	Status    string `json:"status"`     // status pending or done
	Products  []int  `json:"products"`   // Products for order
	ReadyList []int  `json:"ready list"` // complete products list
}

// Orders mop for products
var Orders = make(map[string]Order)

// Vendor slice
var Vendor = [][]int{
	{1, 2},
	{3, 8, 4},
	{5, 6},
}

// Vendors type that can be safely shared between goroutines
type Vendors struct {
	sync.RWMutex
	items [][]int
}

var S = &Vendors{}

// Complete Doing order in vendors and return list of complete products and products
func (cs *Vendors) Complete(order []int) ([]int, []int) {
	cs.Lock()
	defer cs.Unlock()
	complete := []int{}
	required := []int{}
	for k := 0; k < len(order); k++ {
		a := true
		for i, m := range cs.items {
			if len(m) != 0 && order[k] == m[0] {
				a = false
				complete = append(complete, m[0])
				cs.items[i] = append(cs.items[i][:0], cs.items[i][1:]...)
				break
			}
		}
		if a == true {
			required = append(required, order[k])
		}
	}
	fmt.Println("Vendors items", cs.items)
	return complete, required
}

// FillUp add items to the Vendors after delete order
func (cs *Vendors) FillUp(reversion []int) {
	cs.Lock()
	defer cs.Unlock()
	a := cs.items
	sort.Slice(a, func(i, j int) bool {
		return len(a[i]) < len(a[j])
	})

	for i := 0; i < len(reversion); i++ {
		a[0] = append(a[0], reversion[i])
	}
	cs.items = a
	fmt.Println("cs.items", cs.items)
}

// Replenish search and reComplete order
func Replenish() {
	fmt.Println("===========================", Orders)
	for key, value := range Orders {
		if value.Status == Pending && len(value.Products) > 0 {
			complete, required := S.Complete(value.Products)

			if len(complete) > 0 {
				var order Order
				order.ID = key
				order.ReadyList = append(value.ReadyList, complete...)
				order.Products = required
				order.Status = Pending
				if len(required) == 0 {
					order.Status = Done
				}
				Orders[key] = order
			}
		}
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("ok")
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	_ = json.NewDecoder(r.Body).Decode(&order)
	order.ID = strconv.Itoa(rand.Intn(1000000))
	order.Status = Done
	S.items = Vendor
	order.ReadyList, order.Products = S.Complete(order.Products)
	if len(order.Products) > 0 {
		order.Status = Pending

	}
	Orders[order.ID] = order
	json.NewEncoder(w).Encode(&order.ID)
	Replenish()
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	order, ok := Orders[params["id"]]
	if !ok {
		w.Write([]byte("order not found"))
	} else {
		json.NewEncoder(w).Encode(order)
	}
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	order, ok := Orders[params["id"]]
	if !ok {
		w.Write([]byte("order not found"))
	} else {
		delete(Orders, params["id"])
		json.NewEncoder(w).Encode(params["id"])
		S.FillUp(order.ReadyList)
		Replenish()
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", health).Methods("GET")
	router.HandleFunc("/order", createOrder).Methods("POST")
	router.HandleFunc("/order/{id}", getOrder).Methods("GET")
	router.HandleFunc("/order/{id}", deleteOrder).Methods("DELETE")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
