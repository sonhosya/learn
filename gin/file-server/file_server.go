package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.Uint("p", 7890, "listen prot")
	dir := flag.String("d", "./", "share dir")

	flag.Parse()

	g := gin.Default()
	g.StaticFS("/", http.Dir(*dir))
	g.Run(fmt.Sprintf(":%d", *port))
}
