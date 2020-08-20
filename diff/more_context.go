package main

import (
	"context"
	"fmt"
)

type key string

func searchKey(ctx context.Context, k key) {
	v := ctx.Value(k)
	if v != nil {
		fmt.Println("found value:", v)
		return
	}
	fmt.Println("key not found:", k)
}

func main() {
	myKey := key("secretvalue")
	ctx := context.WithValue(context.Background(), myKey, "secretvalue")

	searchKey(ctx, myKey)

	searchKey(ctx, key("notthere"))
	emptyCtx := context.TODO()
	searchKey(emptyCtx, key("notthere"))
}
