package main

import (
	"github.com/kapustaprusta/elbrus/v2/internal/app/elbrus"
)

func main() {
	elbrus.Start(elbrus.NewConfig())
}
