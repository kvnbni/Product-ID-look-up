package main

import (
	"encoding/json"
	"log"
	"github.com/go-redis/redis"
	"io/ioutil"
)

type productRaw struct{
	Status string `json:"status"`
	Prod []productInfo `json:"prod"`
}

type productInfo struct{
	Product_id int `json:"productID"`
	Product_name string `json:"product_name"`
}

func main(){
	//Reading file
	fileContent, err:= ioutil.ReadFile("../productdetails/product_list.txt")
	if err!=nil{
		log.Fatal(err)
	}
	
	//Unmarshaling file content into struct
	var t productRaw
	err=json.Unmarshal(fileContent, &t)
	if err!=nil{
		log.Fatal(err)
	}
	
	//Opening redis client
	redisHandler:= redis.NewClient(&redis.Options{
		Addr: "localhost:6379", //default localhost:redisPort
	})
	
	//Checking redis connectivity
	_, err=redisHandler.Ping().Result() //This returns a *StatusCmd response which has a string, err
	if err!=nil{
		log.Fatal(err)
	}
	
	//Adding product details to redis
	for _, productList := range t.Prod{
		_, err=redisHandler.Set(productList.Product_name, productList.Product_id,0).Result() //Setting key as product name and value as product id and expiration as never (0)
		if err!=nil{
			log.Fatal(err)
		}
	}
}
