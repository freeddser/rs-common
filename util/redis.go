package util

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var Redis Redisconection

type Redisconection struct {
	Main *redis.Pool
}

//InitRedis initialize connection for redis and store it into redis pool
func InitRedis(redisAddr string) {
	clientRedisMain := ConnectRedis(redisAddr)

	Redis = Redisconection{
		Main: clientRedisMain,
	}
}

func ConnectRedis(connStr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 10 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", connStr) },
	}
}

func (r Redisconection) GetIntAndDelete(key string) (int, error) {
	mainRedis := r.Main.Get()
	defer mainRedis.Close()

	dataInt, err := redis.Int(mainRedis.Do("GET", key))
	if err != nil && err != redis.ErrNil {
		log.Println("[redis][getInt] error in retrieving redis data [key: %s] %s", key, err.Error())
		return dataInt, err
	}
	_, err = mainRedis.Do("DEL", key)
	if err != nil {
		log.Println("[redis][getInt] error in del redis data [key: %s] %s", key, err.Error())
		return dataInt, err
	}
	return dataInt, nil
}

func (r Redisconection) IncreaseAuto(key string) error {
	mainRedis := r.Main.Get()
	defer mainRedis.Close()

	//storing to redis
	_, err := mainRedis.Do("INCRBY", key, 1)
	if err != nil {
		log.Println("[redis][increaseAuto] error storing data to redis [key: %s]: %s", key, err.Error())
	}
	log.Println("updating redis", key)

	return nil
}

func (r Redisconection) GetBytes(key string) ([]byte, error) {
	mainRedis := r.Main.Get()
	defer mainRedis.Close()

	dataInBytes, err := redis.Bytes(mainRedis.Do("GET", key))
	if err != nil && err != redis.ErrNil {
		log.Println("[redis][getBytes] error in retrieving redis data [key: %s] %s", key, err.Error())
		return dataInBytes, err
	}

	return dataInBytes, nil
}

func (r Redisconection) SetBytes(key string, data []byte, ttl int) error {
	mainRedis := r.Main.Get()
	defer mainRedis.Close()

	//storing to redis
	_, err := redis.String(mainRedis.Do("SETEX", key, ttl, data))
	if err != nil {
		log.Println("[redis][setBytes] error storing data to redis [key: %s]: %s", key, err.Error())
	}
	log.Println("updating redis", key)

	return nil
}

func (r Redisconection) DelCache(key string) error {
	mainRedis := r.Main.Get()
	defer mainRedis.Close()

	_, err := mainRedis.Do("DEL", key)
	if err != nil {
		log.Println("[redis][getBytes] error in deleting keys : %s, %s", key, err.Error())
	}
	return nil
}
