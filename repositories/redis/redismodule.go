package redismodule

import (
	"fmt"
	"github.com/go-redis/redis"
)

// RedisClient структура для представления клиента Redis
type RedisClient struct {
	client *redis.Client
}

// NewRedisClient создает новый экземпляр RedisClient
func NewRedisClient(addr, password string, db int) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Проверяем подключение к Redis
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Ошибка при подключении к Redis:", err)
		return nil
	}

	return &RedisClient{client}
}

// Set устанавливает значение в Redis для указанного ключа
func (rc *RedisClient) Set(key, value string) error {
	err := rc.client.Set(key, value, 0).Err()
	return err
}

// Get получает значение из Redis для указанного ключа
func (rc *RedisClient) Get(key string) (string, error) {
	val, err := rc.client.Get(key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("Ключ не найден: %s", key)
	} else if err != nil {
		return "", err
	}
	return val, nil
}

// Delete удаляет значение из Redis для указанного ключа
func (rc *RedisClient) Delete(key string) error {
	err := rc.client.Del(key).Err()
	return err
}

// Close закрывает соединение с Redis
func (rc *RedisClient) Close() {
	rc.client.Close()
}