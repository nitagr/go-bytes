package initial

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

type Cache struct {
	*redis.Client
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func NewRedisClient() *Cache {
	address := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	fmt.Println("add", address)
	return &Cache{
		redis.NewClient(&redis.Options{
			Addr: address,
		}),
	}
}
