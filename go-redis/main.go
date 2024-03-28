package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		//Addr: "ip:port"
	})

	s := time.Now()
	for i := 0; i < 500; i++ {
		rdb.Set(context.Background(), fmt.Sprintf("page:%d", i), i, 0)
	}
	fmt.Println(time.Since(s))

	//16.7173ms

	a := time.Now()
	var params []interface{}
	for i := 0; i < 500; i++ {
		params = append(params, fmt.Sprintf("page:%d", i), i)
	}
	rdb.MSet(context.Background(), params...)
	fmt.Println(time.Since(a))
	//528.56µs
}
