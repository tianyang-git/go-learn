package main

import (
	"day7-panic-recover/gee"
	"net/http"
)

// addRoute -> Engine wait
// http -> handle(getRoute -> func)
func main() {
	r := gee.Default()

	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
