package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/dibyochakraborty/kitchen/services/common/genproto/orders"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	conn := NewGRPCClient(":9000")
	defer conn.Close()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		c := orders.NewOrderServiceClient(conn)

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()

		_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerID: 69,
			ProductID:  6969,
			Quantity:   6,
		})

		if err != nil {
			panic("Client error : " + err.Error())
		}

		o, err:=c.GetOrders(ctx, &orders.GetOrdersRequest{
			CustomerID: 69,
			
		})

		if err!=nil{
			panic("Client error : " + err.Error())
		}
		t:=template.Must(template.New("orders").Parse(ordersTemplate))

		nerr := t.Execute(w,o)

		if nerr!=nil{
			log.Fatalf("template error: %v", nerr)
		}


	})

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kitchen Orders</title>
</head>
<body>
    <h1>Orders List</h1>
    <table border="1">
        <tr>
            <th>Order ID</th>
            <th>Customer ID</th>
            <th>Quantity</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.OrderID}}</td>
            <td>{{.CustomerID}}</td>
            <td>{{.Quantity}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`
