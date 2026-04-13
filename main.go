package main

import (
	"fmt"
	"context"
	"github.com/matybq/orders-api/app"
)

func main() {
	app := app.New()
	err := app.Start(context.TODO())

	if err != nil {
		fmt.Println("Error starting app:", err)
	}
}
