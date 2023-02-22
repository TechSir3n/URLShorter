package database

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/subosito/gotenv"
	"time"
//	"os"
)

func init() {
	if err := gotenv.Load(".env"); err != nil {
		panic("Failed to load .env file")
	}
}

type RedisService struct {
	redisClient *redis.Client
}

var (
	redisStore = &RedisService{}
)

const (
	CacheDuration = 6 * time.Hour
)


func InitRedis() *RedisService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:   "redis:6379",
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Println("Redis Success connected:")
	redisStore.redisClient = redisClient

	return redisStore
}

func SaveUrlShort(shortUrl string, longUrl string, userID string) {
	err := redisStore.redisClient.Set(shortUrl, longUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url,Error: %v", err))
	}
}

func GetUrlShort(shortUrl string) string {
	result, err := redisStore.redisClient.Get(shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed get key url,Error: %v", err))
	}

	return result // return long_url 
}
