package rejonson_test

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dxtrleague/rejonson"
	"github.com/go-redis/redis/v8"
)

func init() {
	goRedisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	goRedisClient.Del(context.Background(), "go-redis-cmd", "rejson-cmd", "rejson-cmd-pipeline", "go-redis-pipeline-command")
	_ = goRedisClient.Close()
}

func ExampleExtendClient() {
	ctx := context.Background()
	goRedisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	client := rejonson.ExtendClient(goRedisClient)
	defer client.Close()

	arr := []interface{}{"hello", "world", 1, map[string]interface{}{"key": 12}}
	js, err := json.Marshal(arr)
	if err != nil {
		// handle
	}
	// redis "native" command
	client.Set(ctx, "go-redis-cmd", "hello", time.Second)
	client.JsonSet(ctx, "rejson-cmd", ".", string(js))

	// int command
	arrLen, err := client.JsonArrLen(ctx, "rejson-cmd", ".").Result()
	if err != nil {
		// handle
	}

	fmt.Printf("Array length: %d", arrLen)
	// Output: Array length: 4
}

func ExampleExtendPipeline() {
	ctx := context.Background()
	goRedisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	client := rejonson.ExtendClient(goRedisClient)

	pipeline := client.Pipeline()
	pipeline.JsonSet(ctx, "rejson-cmd-pipeline", ".", "[10]")
	pipeline.JsonNumMultBy(ctx, "rejson-cmd-pipeline", "[0]", 10)
	pipeline.Set(ctx, "go-redis-pipeline-command", "hello from go-redis", time.Second)

	_, err := pipeline.Exec()
	if err != nil {
		// handle error
	}
	jsonString, err := client.JsonGet(ctx, "rejson-cmd-pipeline").Result()
	if err != nil {
		// handle error
	}

	fmt.Printf("Array %s", jsonString)

	// Output: Array [100]
}
