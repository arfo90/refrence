package main

import (
	"context"
	"fmt"
)

type Request struct {
	Value string
}

func main() {
	r := Request{
		Value: "hi",
	}
	ctx := context.WithValue(context.Background(), "tset", r)
	scrub(ctx)

}

func scrub(ctx context.Context) {
	r := ctx.Value("test")
	fmt.Printf("%v", r)
}
