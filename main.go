package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func sendEmail(email string) {
	fmt.Println("sent")
}

func worker(client *redis.Client, ctx context.Context, key string) {
	for {
		//result, err := client.BLPop(ctx, 4*time.Second, key, "ss").Result()
		//if err != nil {
		//	fmt.Println(err)
		//	continue
		//}

		result, err := client.LPop(ctx, key).Result()
		if err != nil {
			continue
		}

		fmt.Println("runing")
		marshal, _ := json.Marshal(result)

		fmt.Println(string(marshal))
	}
}

func main() {
	ctx := context.Background()
	key := "tes"
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()

	go worker(client, ctx, key)

	tasks := []string{"task1", "task2", "task3", "task4"}

	for i := 0; i < 10; i++ {
		client.RPush(ctx, key, tasks)
		time.Sleep(2000 * time.Millisecond)

	}

	select {}

}
