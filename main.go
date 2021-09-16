package main

import (
	"context"
	"fmt"

	"github.com/lnikon/go-google-search/pkg/userip"
)

func main() {
	fmt.Printf("%s", "Hello!")
	userip.FromContext(context.TODO())
}
