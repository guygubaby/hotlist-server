package main

import (
	"github.com/guygubaby/hotlist-server/router"
)

func main() {
	r := router.InitRouter()
	_ = r.Run(":5000")
}
