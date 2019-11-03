package redisfunc
//package redisfunc helps with basic redis funtions like get and set.

import(
	"github.com/go-redis/redis"
	"time"
)

//RedisGet replicates the redis get function.
//It takes as parameters a key and the redis client.
//Returns the corresponding value for the key and error if any.
func RedisGet(key string, redisHandler *redis.Client)(string,error){
	productId, err:=redisHandler.Get(key).Result()
	if err!=nil{
		if err==redis.Nil{  //The redis get function returns a redis.Nil error if there is no entry for a key
				    //within redis
			return "No product info available", nil
		}else{
			return "",err
		}
	}else{
		return productId,nil
	}
}

//RedisPost repliacates the redis set function.
//It takes as parameters a key, value and expiration.
//It returns the status code for the set function and error if any.
func RedisPost(key string, value string, expiration time.Duration, redisHandler *redis.Client)(string,error){
	status, err := redisHandler.Set(key,value,expiration).Result()
	if err!=nil{
		return "",err
	}else{
		return status,nil
	}
}
